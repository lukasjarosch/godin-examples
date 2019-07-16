module github.com/lukasjarosch/godin-examples/user

go 1.12

require (
	github.com/go-godin/log v0.0.0-20190715125052-26f1fab6b64a
	github.com/go-godin/middleware v0.0.0-20190715143930-be7d7bc7f5dd
	github.com/go-kit/kit v0.9.0
	github.com/golang/protobuf v1.3.1
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/lukasjarosch/godin v0.0.0-20190623111143-62716b922f9c
	github.com/oklog/oklog v0.3.2
	github.com/oklog/run v1.0.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.0.0
	github.com/rs/xid v1.2.1
	github.com/sirupsen/logrus v1.4.2
	github.com/streadway/amqp v0.0.0-20190404075320-75d898a42a94
	google.golang.org/grpc v1.22.0
)

replace github.com/lukasjarosch/godin => ../../godin
