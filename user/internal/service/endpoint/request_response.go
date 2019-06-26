// Code generated by Godin vv0.4.0; DO NOT EDIT.

package endpoint

import (
	"github.com/lukasjarosch/godin-examples/user/internal/service"
)

type (
	CreateRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	CreateResponse struct {
		User *service.UserEntity `json:"user"`
		Err  error               `json:"-"`
	}

	GetRequest struct {
		Id string `json:"id"`
	}

	GetResponse struct {
		User *service.UserEntity `json:"user"`
		Err  error               `json:"-"`
	}

	ListRequest struct {
	}

	ListResponse struct {
		Users []*service.UserEntity `json:"users"`
		Err   error                 `json:"-"`
	}

	DeleteRequest struct {
		Id string `json:"id"`
	}

	DeleteResponse struct {
		Err error `json:"-"`
	}
)

// Implement the Failer interface for all responses
func (resp CreateResponse) Failed() error { return resp.Err }
func (resp GetResponse) Failed() error    { return resp.Err }
func (resp ListResponse) Failed() error   { return resp.Err }
func (resp DeleteResponse) Failed() error { return resp.Err }
