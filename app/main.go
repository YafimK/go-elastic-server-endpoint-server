package main

import (
	"github.com/YafimK/go-elastic-server-endpoint-server/controllers"
	"github.com/YafimK/go-elastic-server-endpoint-server/elastic_service"
	"github.com/YafimK/go-elastic-server-endpoint-server/routes"
	"github.com/YafimK/go-elastic-server-endpoint-server/server"
	"log"
)

func main() {
	elasticClient, err := elastic_service.NewElasticClient(RuntimeSettings().ElasticServerAddress.String(), RuntimeSettings().ElasticServerIndex)
	if err != nil {
		log.Fatalf("Failed startiing elastic server client: %v\n", err)
	}
	endpointServer := server.NewEndpointServer(RuntimeSettings().EndpointServerHostAddress.Host)
	routes.NewRoutingMap(controllers.NewSearchController(elasticClient)).RegisterRoutes(endpointServer.Router())

	endpointServer.Start()
}
