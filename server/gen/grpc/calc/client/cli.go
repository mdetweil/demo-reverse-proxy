// Code generated by goa v3.0.9, DO NOT EDIT.
//
// calc gRPC client CLI support package
//
// Command:
// $ goa gen github.com/mdetweil/demo-reverse-proxy/server/design

package client

import (
	"encoding/json"
	"fmt"

	calc "github.com/mdetweil/demo-reverse-proxy/server/gen/calc"
	calcpb "github.com/mdetweil/demo-reverse-proxy/server/gen/grpc/calc/pb"
)

// BuildAddPayload builds the payload for the calc add endpoint from CLI flags.
func BuildAddPayload(calcAddMessage string) (*calc.AddPayload, error) {
	var err error
	var message calcpb.AddRequest
	{
		if calcAddMessage != "" {
			err = json.Unmarshal([]byte(calcAddMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"a\": 360622074634248926,\n      \"b\": 8133055152903002499\n   }'")
			}
		}
	}
	v := &calc.AddPayload{
		A: int(message.A),
		B: int(message.B),
	}
	return v, nil
}

// BuildMultiplyPayload builds the payload for the calc multiply endpoint from
// CLI flags.
func BuildMultiplyPayload(calcMultiplyMessage string) (*calc.MultiplyPayload, error) {
	var err error
	var message calcpb.MultiplyRequest
	{
		if calcMultiplyMessage != "" {
			err = json.Unmarshal([]byte(calcMultiplyMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"a\": 5401762099778430809,\n      \"b\": 1918630006328122782\n   }'")
			}
		}
	}
	v := &calc.MultiplyPayload{
		A: int(message.A),
		B: int(message.B),
	}
	return v, nil
}