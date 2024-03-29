// Code generated by Godin vv0.4.0; DO NOT EDIT.

package endpoint

import (
	"yyy/internal/service"
)

type (
	HelloRequest struct {
		Name string `json:"name"`
	}

	HelloResponse struct {
		Greeting service.string `json:"greeting"`
		Err      error          `json:"-"`
	}
)

// Implement the Failer interface for all responses
func (resp HelloResponse) Failed() error { return resp.Err }
