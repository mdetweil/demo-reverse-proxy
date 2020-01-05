// Code generated by goa v3.0.9, DO NOT EDIT.
//
// world endpoints
//
// Command:
// $ goa gen github.com/mdetweil/demo-reverse-proxy/server/design

package world

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "world" service endpoints.
type Endpoints struct {
	Hello goa.Endpoint
}

// NewEndpoints wraps the methods of the "world" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Hello: NewHelloEndpoint(s),
	}
}

// Use applies the given middleware to all the "world" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Hello = m(e.Hello)
}

// NewHelloEndpoint returns an endpoint function that calls the method "hello"
// of service "world".
func NewHelloEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Hello(ctx)
	}
}
