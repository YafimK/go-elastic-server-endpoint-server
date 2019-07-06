package api

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

func (es EndpointServer) Start() {
	fmt.Printf("Starting server on - %v\n", es.host)
	if err := es.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}
