package calcrouter

import (
	"context"
	"encoding/json"
	calcpb "github.com/mdetweil/demo-reverse-proxy/http_proxy/pb/calc"
	"google.golang.org/grpc"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type calcDialer struct {
	Client calcpb.CalcClient
}

func NewCalcRouter(conn *grpc.ClientConn, r *mux.Router) {
	client := calcpb.NewCalcClient(conn)
	cr := calcDialer{Client: client}
	sr := r.PathPrefix("/calc").Subrouter()
	initalizeRoutes(cr, sr)
}

func initalizeRoutes(cD calcDialer, r *mux.Router) {
	r.HandleFunc("/multiply", cD.processMultiplyRequest).Methods("GET").Name("Multiply 2 numbers, a & b")
	r.HandleFunc("/add", cD.processAddRequest).Methods("GET").Name("Add 2 numbers, a & b")
}

func (c *calcDialer) processAddRequest(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//GET A and B from Request
	q := r.URL.Query()
	a, err := strconv.ParseInt(q.Get("a"), 10, 32)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	b, err := strconv.ParseInt(q.Get("b"), 10, 32)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	res, err := c.Client.Add(ctx, &calcpb.AddRequest{A: int32(a), B: int32(b)})
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	js, _ := json.Marshal(res.GetField())
	w.Write(js)
}

func (c *calcDialer) processMultiplyRequest(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//GET A and B from Request
	q := r.URL.Query()
	a, err := strconv.ParseInt(q.Get("a"), 10, 32)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	b, err := strconv.ParseInt(q.Get("b"), 10, 32)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	res, err := c.Client.Multiply(ctx, &calcpb.MultiplyRequest{A: int32(a), B: int32(b)})
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	js, _ := json.Marshal(res.GetField())
	w.Write(js)
}
