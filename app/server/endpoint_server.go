package server

import (
	"fmt"
	"log"
	"net/http"
)

type EndpointServer struct {
	router *http.ServeMux
	server *http.Server
	host   string
}

func (server EndpointServer) Router() *http.ServeMux {
	return server.router
}

func NewEndpointServer(host string) *EndpointServer {
	router := http.NewServeMux()
	endpointServer := &EndpointServer{
		router: router,
		server: &http.Server{
			Addr:    host,
			Handler: router,
		},
		host: host}
	return endpointServer
}

func (server EndpointServer) Start() {
	fmt.Printf("Starting server on - %v\n", server.host)
	if err := server.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}
