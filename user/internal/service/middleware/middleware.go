package middleware

import (
	"github.com/lukasjarosch/godin-examples/user/internal/service/usecase"
)

type Middleware func(service usecase.Service) usecase.Service
