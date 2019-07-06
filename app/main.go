package main

import (
	"PeX/api"
	"PeX/elastic_service"
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
