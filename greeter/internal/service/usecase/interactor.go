package usecase

import (
	_ "github.com/lukasjarosch/godin-examples/greeter/internal/service"
)

type Logger interface {
	Debug(message string, keyvals ...interface{})
	Info(message string, keyvals ...interface{})
	Warning(message string, keyvals ...interface{})
	Error(message string, keyvals ...interface{})
}

type Greeting struct {
	Name string
}

// Greeter documentation is automatically
// added to the README.
type Greeter interface {
	// SayHello greets you. This comment is also automatically added to the README.
	// Also make sure that all parameters are named, Godin requires this information in order to work.
	SayHello(name string) (greeting Greeting, err error)
}

type greeterInteractor struct {
	logger Logger
}

func NewGreeterInteractor(logger Logger) *greeterInteractor {
	return &greeterInteractor{
		logger: logger,
	}
}
