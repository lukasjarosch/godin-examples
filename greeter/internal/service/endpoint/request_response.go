// Code generated by Godin vv0.4.0; DO NOT EDIT.

package endpoint

import (
	_ "github.com/lukasjarosch/godin-examples/greeter/internal/service"
)

type (
	HelloRequest struct {
		Name string `json:"name"`
	}

	HelloResponse struct {
		Greeting string `json:"greeting"`
		Err      error  `json:"-"`
	}
)

// Implement the Failer interface for all responses
func (resp HelloResponse) Failed() error { return resp.Err }
