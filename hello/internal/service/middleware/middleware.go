package middleware

import (
	"github.com/lukasjarosch/godin-examples/hello/internal/service"
)

type Middleware func(service service.Hello) service.Hello
