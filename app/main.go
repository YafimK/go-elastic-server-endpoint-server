package main

import (
	"github.com/YafimK/go-elastic-server-endpoint-server/api"
	"github.com/YafimK/go-elastic-server-endpoint-server/elastic_service"
	"log"
)

func main() {
	_, err := elastic_service.NewElasticClient(RuntimeSettings().ElasticServerAddress.String(), RuntimeSettings().ElasticServerIndex)
	if err != nil {
		log.Fatalf("Failed startiing elastic server client: %v\n", err)
	}
	endpointServer := api.NewEndpointServer(RuntimeSettings().EndpointServerHostAddress.Host)
	endpointServer.Start()
}
