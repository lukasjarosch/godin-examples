// Code generated by Godin v0.4.0; DO NOT EDIT.
package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/oklog/oklog/pkg/group"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	googleGrpc "google.golang.org/grpc"

	pb "github.com/lukasjarosch/godin-examples/user/api"
	svcGrpc "github.com/lukasjarosch/godin-examples/user/internal/grpc"
	"github.com/lukasjarosch/godin-examples/user/internal/service"
	"github.com/lukasjarosch/godin-examples/user/internal/service/endpoint"
	"github.com/lukasjarosch/godin-examples/user/internal/service/middleware"
	"github.com/lukasjarosch/godin-examples/user/internal/service/user"

	"github.com/lukasjarosch/godin/pkg/log"
)

var DebugAddr = getEnv("DEBUG_ADDRESS", "0.0.0.0:3000")
var GrpcAddr = getEnv("GRPC_ADDRESS", "0.0.0.0:50051")

// group to manage the lifecycle for goroutines
var g group.Group

func main() {
	logger := log.New()

	// initialize service layer
	var svc service.User
	svc = user.NewServiceImplementation(logger)
	svc = middleware.LoggingMiddleware(logger)(svc)
	//TODO: svc = middleware.AuthorizationMiddleware(logger)(svc)
	//TODO: svc = middleware.RecoveringMiddleware(logger)(svc)

	// initialize endpoint and transport layers
	var (
		endpoints   = endpoint.Endpoints(svc, logger)
		grpcHandler = svcGrpc.NewServer(endpoints, logger)
	)

	// serve gRPC server
	grpcServer := googleGrpc.NewServer(
		googleGrpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
	g.Add(initGrpc(grpcServer, grpcHandler, logger), func(error) {
		grpcServer.GracefulStop()
	})
	grpc_prometheus.Register(grpcServer)

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
func initGrpc(grpcServer *googleGrpc.Server, handler pb.UserServiceServer, logger log.Logger) func() error {
	grpcListener, err := net.Listen("tcp", GrpcAddr)
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}

	return func() error {
		logger.Log("transport", "gRPC", "addr", GrpcAddr)
		pb.RegisterUserServiceServer(grpcServer, handler)
		return grpcServer.Serve(grpcListener)
	}
}

func initDebugHttp(logger log.Logger) net.Listener {
	debugListener, err := net.Listen("tcp", DebugAddr)
	if err != nil {
		logger.Log("transport", "debug/HTTP", "during", "Listen", "err", err)
		os.Exit(1)
	}
	return debugListener
}
