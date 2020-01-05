package calcRouter

import (
	"context"
	"net/url"
	"strconv"
	"regexp"
	"fmt"
	"google.golang.org/grpc"
	calcpb "github.com/mdetweil/demo-reverse-proxy/proxy/pb/calc"

)

var (
	calcMap map[string]*regexp.Regexp
)

type CalcDialer interface {
	RouteCalcRequest(path string, queries url.Values) (res interface{}, err error)
}

type calcDialer struct {
	Client calcpb.CalcClient
}

func NewCalcDialer(conn *grpc.ClientConn) CalcDialer {
	client := calcpb.NewCalcClient(conn)
	initalizeCalcMap()
	return &calcDialer{Client: client}
}

func initalizeCalcMap() {
	calcMap = make(map[string]*regexp.Regexp)
	calcMap["add"] = regexp.MustCompile(`^/calc/add`)
	calcMap["multiply"] = regexp.MustCompile(`^/calc/multiply`)
}

func(c *calcDialer) RouteCalcRequest(path string, q url.Values) (interface{}, error) {
	switch {
	case calcMap["add"].MatchString(path):
		return c.processAddRequest(q)
	case calcMap["multiply"].MatchString(path):
		return c.processMultiplyRequest(q)
	}
	return nil, fmt.Errorf("Failed to find matching path: %v", path)
}

func (c *calcDialer) processAddRequest(q url.Values) (interface{}, error) {
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
	res, err := c.Client.Add(ctx, &calcpb.AddRequest{A: int32(a), B: int32(b)})
	if err != nil {
		return nil, err
	}
	return res.GetField(), nil
}

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