module github.com/lukasjarosch/godin-examples/greeter2

go 1.12

require (
	github.com/go-godin/log v0.0.0-20190715125052-26f1fab6b64a
	github.com/go-godin/middleware v0.0.0-20190715143930-be7d7bc7f5dd
	github.com/go-kit/kit v0.9.0
	github.com/lukasjarosch/godin-examples/greeter v0.0.0-00010101000000-000000000000 // indirect
	google.golang.org/grpc v1.22.0
)

replace github.com/lukasjarosch/godin-examples/greeter => ../greeter
