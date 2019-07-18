package middleware

import (
	"github.com/lukasjarosch/godin-examples/greeter/internal/service"
)

type Middleware func(service service.Greeter) service.Greeter
