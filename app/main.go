package main

import (
	"github.com/YafimK/go-elastic-server-endpoint-server/config"
	"github.com/YafimK/go-elastic-server-endpoint-server/routes"
	"github.com/YafimK/go-elastic-server-endpoint-server/server"
)

func main() {
	endpointServer := server.NewEndpointServer(config.RuntimeSettings().EndpointServerHostAddress.Host)
	routes.NewSeachRoutingMap().RegisterRoutes(endpointServer.Router())
	endpointServer.Start()
}
