package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	HTTP_HOST = "127.0.0.1"
	HTTP_PORT = 4444
)

// HTTPServer ...
type HTTPServer struct {
	router *mux.Router
}

// Init ...
func (h *HTTPServer) Init() {
	h.router = mux.NewRouter()
}

// RegisterRouteHandler ...
func (h *HTTPServer) RegisterRouteHandler(path string, handle http.HandlerFunc, method string) {
	h.router.HandleFunc(path, handle).Methods(method)
}

// StartHTPServer ...
func (h *HTTPServer) StartHTPServer() {

	server := &http.Server{
		Handler:      handlers.CompressHandler(h.router),
		Addr:         fmt.Sprintf("%s:%d", HTTP_HOST, HTTP_PORT),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server Running on: ", HTTP_HOST, ":", HTTP_PORT)
	log.Fatalln(server.ListenAndServe())
}
