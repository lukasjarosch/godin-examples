module github.com/lukasjarosch/godin-examples/stringer

go 1.12

require (
	github.com/go-godin/grpc-interceptor v0.0.0-00010101000000-000000000000
	github.com/go-godin/grpc-metadata v0.0.0-20190717081641-be8cff64989a
	github.com/go-godin/log v0.0.0-20190716173926-b62a2fca0801
	github.com/go-godin/middleware v0.0.0-20190717080225-2a88e633669f
	github.com/go-godin/rabbitmq v0.0.0-20190719103636-6b979908c7b1
	github.com/go-kit/kit v0.9.0
	github.com/golang/protobuf v1.3.1
	github.com/lukasjarosch/godin-examples/greeter v0.0.0-00010101000000-000000000000
	github.com/lukasjarosch/godin-examples/user v0.0.0-20190717135725-2e6d11a47b78
	github.com/oklog/oklog v0.3.2
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.0.0
	google.golang.org/grpc v1.22.0
)

replace github.com/lukasjarosch/godin-examples/greeter => ../greeter

replace github.com/go-godin/grpc-interceptor => ../../go-godin/grpc-interceptor

replace github.com/go-godin/rabbitmq => ../../go-godin/rabbitmq

replace github.com/go-godin/middleware => ../../go-godin/middleware
