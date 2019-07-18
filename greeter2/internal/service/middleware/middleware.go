package middleware

import (
	"github.com/lukasjarosch/godin-examples/greeter2/internal/service"
)

type Middleware func(service service.Greeter) service.Greeter
