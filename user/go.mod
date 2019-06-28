module github.com/lukasjarosch/godin-examples/user

go 1.12

require (
	github.com/go-kit/kit v0.8.0
	github.com/golang/protobuf v1.3.1
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/lukasjarosch/godin v0.0.0-20190623111143-62716b922f9c
	github.com/oklog/oklog v0.3.2
	github.com/oklog/run v1.0.0 // indirect
	github.com/prometheus/client_golang v1.0.0
	github.com/rs/xid v1.2.1
	github.com/streadway/amqp v0.0.0-20190404075320-75d898a42a94
	google.golang.org/grpc v1.21.1
)

replace github.com/lukasjarosch/godin => ../../godin
