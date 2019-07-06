package main

import "log"

func main() {
	_, err := NewElasticClient(RuntimeSettings().ElasticServerAddress.String(), RuntimeSettings().ElasticServerIndex)
	if err != nil {
		log.Fatalf("Failed startiing elastic server client: %v\n", err)
	}
	endpointServer := NewEndpointServer(RuntimeSettings().EndpointServerHostAddress.Host)
	endpointServer.Start()
}
