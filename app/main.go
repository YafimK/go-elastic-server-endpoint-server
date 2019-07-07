package main

import (
	"github.com/YafimK/go-elastic-server-endpoint-server/elastic_service"
	"github.com/YafimK/go-elastic-server-endpoint-server/routes"
	"github.com/YafimK/go-elastic-server-endpoint-server/server"
	"log"
)

func main() {
	endpointServer := server.NewEndpointServer(RuntimeSettings().EndpointServerHostAddress.Host)
	routes.NewRoutingMap().RegisterRoutes(endpointServer.Router())

	endpointServer.Start()
}
