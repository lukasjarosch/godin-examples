// Code generated by Godin vv0.4.0; DO NOT EDIT.

package endpoint

import (
	"github.com/lukasjarosch/godin-examples/user/internal/service/usecase"
)

type (
	CreateRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	CreateResponse struct {
		User *usecase.User `json:"user"`
		Err  error         `json:"-"`
	}

	GetRequest struct {
		Id string `json:"id"`
	}

	GetResponse struct {
		User *usecase.User `json:"user"`
		Err  error         `json:"-"`
	}

	ListRequest struct {
	}

	ListResponse struct {
		Users []*usecase.User `json:"users"`
		Err   error           `json:"-"`
	}

	DeleteRequest struct {
		Id string `json:"id"`
	}

	DeleteResponse struct {
		Err error `json:"-"`
	}

	UserCreatedEvent struct {
		User *usecase.User
	}
)

// Implement the Failer interface for all responses
func (resp CreateResponse) Failed() error { return resp.Err }
func (resp GetResponse) Failed() error    { return resp.Err }
func (resp ListResponse) Failed() error   { return resp.Err }
func (resp DeleteResponse) Failed() error { return resp.Err }