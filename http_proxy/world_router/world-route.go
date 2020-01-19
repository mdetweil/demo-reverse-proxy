package worldrouter

import (
	"context"
	"github.com/gorilla/mux"
	worldpb "github.com/mdetweil/demo-reverse-proxy/http_proxy/pb/world"
	"google.golang.org/grpc"
	"net/http"
)

type worldDialer struct {
	Client worldpb.WorldClient
}

func NewWorldRouter(conn *grpc.ClientConn, r *mux.Router) {
	client := worldpb.NewWorldClient(conn)
	wd := worldDialer{Client: client}
	sr := r.PathPrefix("/world").Subrouter()
	initalizeRoutes(wd, sr)
}

func initalizeRoutes(wD worldDialer, r *mux.Router) {
	r.HandleFunc("", wD.processIndex).Methods("GET").Name("Hello World")
}

func (wd *worldDialer) processIndex(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	res, err := wd.Client.Hello(ctx, &worldpb.HelloRequest{})
	if err != nil {
		w.Write([]byte(string(err.Error())))
	}
	w.Write([]byte(res.GetField()))
}
