// Code generated by Godin vv0.4.0; DO NOT EDIT.
package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-godin/grpc-interceptor"
	"github.com/go-godin/rabbitmq"
	"github.com/oklog/oklog/pkg/group"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	googleGrpc "google.golang.org/grpc"

	pb "github.com/lukasjarosch/godin-examples/greeter/api"
	"github.com/lukasjarosch/godin-examples/stringer/internal/amqp"
	svcGrpc "github.com/lukasjarosch/godin-examples/stringer/internal/grpc"
	"github.com/lukasjarosch/godin-examples/stringer/internal/service"
	"github.com/lukasjarosch/godin-examples/stringer/internal/service/endpoint"
	"github.com/lukasjarosch/godin-examples/stringer/internal/service/middleware"
	"github.com/lukasjarosch/godin-examples/stringer/internal/service/usecase"

	"github.com/go-godin/log"
)

var DebugAddr = getEnv("DEBUG_ADDRESS", "0.0.0.0:3000")
var GrpcAddr = getEnv("GRPC_ADDRESS", "0.0.0.0:50051")

// group to manage the lifecycle for goroutines
var g group.Group

func main() {
	logger := log.NewLoggerFromEnv()
	rabbitmqSubConn := initRabbitMQ(logger)
	defer rabbitmqSubConn.Close()
	rabbitmqPubConn := initRabbitMQ(logger)
	defer rabbitmqPubConn.Close()
	// init publishers
	publishers := amqp.Publishers(rabbitmqPubConn, logger)

	// initialize service including middleware
	var svc service.Stringer

	svc = usecase.NewServiceImplementation(logger, publishers)
	svc = middleware.LoggingMiddleware(logger)(svc)

	// initialize endpoint and transport layers
	var (
		endpoints   = endpoint.Endpoints(svc, logger)
		grpcHandler = svcGrpc.NewServer(endpoints, logger)
	)
	// setup AMQP subscriptions
	subscriptions := amqp.Subscriptions(rabbitmqSubConn)
	if err := subscriptions.UserCreatedSubscriber(logger, svc); err != nil {
		logger.Error("failed to create subscription", "err", err)
		os.Exit(-1)
	}
	if err := subscriptions.UserDeletedSubscriber(logger, svc); err != nil {
		logger.Error("failed to create subscription", "err", err)
		os.Exit(-1)
	}

	// serve gRPC server
	grpcServer := googleGrpc.NewServer(
		googleGrpc.UnaryInterceptor(grpc_interceptor.UnaryInterceptor),
	)
	g.Add(initGrpc(grpcServer, grpcHandler, logger), func(error) {
		grpcServer.GracefulStop()
	})

	// serve debug http server (prometheus)
	http.DefaultServeMux.Handle("/metrics", promhttp.Handler())
	debugListener := initDebugHttp(logger)
	g.Add(func() error {
		logger.Log("transport", "debug/HTTP", "addr", DebugAddr)
		return http.Serve(debugListener, http.DefaultServeMux)
	}, func(error) {
		debugListener.Close()
	})

	// Wait for SIGINT or SIGTERM and stop gracefully
	cancelInterrupt := make(chan struct{})
	g.Add(shutdownHandler(cancelInterrupt), func(e error) {
		close(cancelInterrupt)
	})

	// run
	if err := g.Run(); err != nil {
		logger.Log("fatal", err)
		os.Exit(1)
	}
}

// getEnv get key environment variable if exist otherwise return defalutValue
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

// shutdownHandler to handle graceful shutdowns
func shutdownHandler(interruptChannel chan struct{}) func() error {
	return func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-interruptChannel:
			return nil
		}
	}
}

// initGrpc serve up GRPC
func initGrpc(grpcServer *googleGrpc.Server, handler pb.GreeterServiceServer, logger log.Logger) func() error {
	grpcListener, err := net.Listen("tcp", GrpcAddr)
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}

	return func() error {
		logger.Log("transport", "gRPC", "addr", GrpcAddr)
		pb.RegisterGreeterServiceServer(grpcServer, handler)
		return grpcServer.Serve(grpcListener)
	}
}

// initRabbitMQ will initialize the amqp connection and create a new channel
func initRabbitMQ(logger log.Logger) *rabbitmq.RabbitMQ {
	rabbitmqConn, err := rabbitmq.NewRabbitMQFromEnv()
	if err != nil {
		logger.Error("failed to initialize RabbitMQ connection", "err", err)
		os.Exit(-1)
	}
	if err := rabbitmqConn.Connect(); err != nil {
		logger.Error("failed to connect to RabbitMQ", "err", err)
		os.Exit(-1)
	}
	if rabbitmqConn.Channel, err = rabbitmqConn.NewChannel(); err != nil {
		logger.Error("failed to create AMQP channel", "err", err)
		os.Exit(-1)
	}

	return rabbitmqConn
}

func initDebugHttp(logger log.Logger) net.Listener {
	debugListener, err := net.Listen("tcp", DebugAddr)
	if err != nil {
		logger.Log("transport", "debug/HTTP", "during", "Listen", "err", err)
		os.Exit(1)
	}
	return debugListener
}
