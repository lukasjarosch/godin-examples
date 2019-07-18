module github.com/lukasjarosch/godin-examples/hello

go 1.12

require (
	github.com/go-godin/grpc-metadata v0.0.0-20190717081641-be8cff64989a
	github.com/go-godin/log v0.0.0-20190716173926-b62a2fca0801
	github.com/go-godin/middleware v0.0.0-20190716125117-8d9e256a3b95
	github.com/go-godin/rabbitmq v0.0.0-20190717074815-2a37ca6d6428
	github.com/go-kit/kit v0.9.0
	github.com/golang/protobuf v1.3.1
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/oklog/oklog v0.3.2
	github.com/oklog/run v1.0.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.0.0
	github.com/streadway/amqp v0.0.0-20190404075320-75d898a42a94
	google.golang.org/grpc v1.22.0
)

replace github.com/lukasjarosch/godin-examples/hello => ./

replace github.com/go-godin/rabbitmq => ../../go-godin/rabbitmq

replace github.com/go-godin/middleware => ../../go-godin/middleware
