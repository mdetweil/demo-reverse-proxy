package calcapi

import (
	"context"
	"log"

	world "github.com/mdetweil/demo-reverse-proxy/server/gen/world"
)

// world service example implementation.
// The example methods log the requests and return zero values.
type worldsrvc struct {
	logger *log.Logger
}

// NewWorld returns the world service implementation.
func NewWorld(logger *log.Logger) world.Service {
	return &worldsrvc{logger}
}

// Hello implements hello.
func (s *worldsrvc) Hello(ctx context.Context) (res string, err error) {
	s.logger.Print("hello world function called")
	return "Hello World", nil
}
