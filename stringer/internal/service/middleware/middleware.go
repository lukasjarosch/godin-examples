package middleware

import (
	"github.com/lukasjarosch/godin-examples/stringer/internal/service"
)

type Middleware func(service service.Stringer) service.Stringer
