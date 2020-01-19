package main

import (
	"context"
	"flag"
	"fmt"
	calc "github.com/mdetweil/demo-reverse-proxy/http_proxy/calc_router"
	world "github.com/mdetweil/demo-reverse-proxy/http_proxy/world_router"

	"github.com/gorilla/mux"

	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"time"
)

var (
	listenAddr string
	healthy    int32
)

func main() {
	flag.StringVar(&listenAddr, "listen-addr", "8088", "server listen address")
	flag.Parse()

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Server is starting...")

	conn, _ := grpc.Dial("localhost:8080", grpc.WithInsecure())
	defer conn.Close()

	r := mux.NewRouter()
	r.HandleFunc("/healthz", healthz)

	calc.NewCalcRouter(conn, r)
	world.NewWorldRouter(conn, r)

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%v", listenAddr),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
		Handler:      r,
	}
	log.Println("The following routes have been established:")
	_ = r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		log.Printf("\t%q", t)
		return nil
	})

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		logger.Println("Server is shutting down...")
		atomic.StoreInt32(&healthy, 0)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	logger.Println("Server is ready to handle requests at", listenAddr)
	atomic.StoreInt32(&healthy, 1)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}

	<-done
	logger.Println("Server stopped")
}

type Response struct {
	Data        interface{}
	Error       interface{}
	RequestID   string
	PagingToken interface{}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	if atomic.LoadInt32(&healthy) == 1 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusServiceUnavailable)
}
