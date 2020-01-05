package worldRouter

import (
	"context"
	"net/url"
	"regexp"
	"fmt"
	"google.golang.org/grpc"
	worldpb "github.com/mdetweil/demo-reverse-proxy/proxy/pb/world"

)

var (
	worldMap map[string]*regexp.Regexp
)

type WorldDialer interface {
	RouteWorldRequest(path string, queries url.Values) (res interface{}, err error)
}

type worldDialer struct {
	Client worldpb.WorldClient
}

func NewWorldDialer(conn *grpc.ClientConn) WorldDialer {
	client := worldpb.NewWorldClient(conn)
	initalizeWorldMap()
	return &worldDialer{Client: client}
}

func initalizeWorldMap() {
	worldMap = make(map[string]*regexp.Regexp)
	worldMap["/"] = regexp.MustCompile(`^/world`)
}

func(w *worldDialer) RouteWorldRequest(path string, q url.Values) (interface{}, error) {
	switch {
	case worldMap["/"].MatchString(path):
		return w.processIndex(q)
	}
	return nil, fmt.Errorf("Failed to find matching path: %v", path)
}

func (w *worldDialer) processIndex(q url.Values) (interface{}, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	res, err := w.Client.Hello(ctx, &worldpb.HelloRequest{})
	if err != nil {
		return nil, err
	}
	return res.GetField(), nil
}

/*
func (c *calcDialer) processMultiplyRequest(q url.Values) (interface{}, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//GET A and B from Request
	a, err := strconv.ParseInt(q["a"][0], 10, 32)
	if err != nil {
		return nil, err
	}
	b, err := strconv.ParseInt(q["b"][0], 10, 32)
	if err != nil {
		return nil, err
	}
	res, err := c.Client.Multiply(ctx, &calcpb.MultiplyRequest{A: int32(a), B: int32(b)})
	if err != nil {
		return nil, err
	}
	return res.GetField(), nil
}
*/