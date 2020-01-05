// Code generated by goa v3.0.9, DO NOT EDIT.
//
// calc client
//
// Command:
// $ goa gen github.com/mdetweil/demo-reverse-proxy/server/design

package calc

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "calc" service client.
type Client struct {
	AddEndpoint      goa.Endpoint
	MultiplyEndpoint goa.Endpoint
}

// NewClient initializes a "calc" service client given the endpoints.
func NewClient(add, multiply goa.Endpoint) *Client {
	return &Client{
		AddEndpoint:      add,
		MultiplyEndpoint: multiply,
	}
}

// Add calls the "add" endpoint of the "calc" service.
func (c *Client) Add(ctx context.Context, p *AddPayload) (res int, err error) {
	var ires interface{}
	ires, err = c.AddEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}

// Multiply calls the "multiply" endpoint of the "calc" service.
func (c *Client) Multiply(ctx context.Context, p *MultiplyPayload) (res int, err error) {
	var ires interface{}
	ires, err = c.MultiplyEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}