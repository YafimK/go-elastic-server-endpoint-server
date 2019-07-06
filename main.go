package main

func main() {
	elasticClient := NewElasticClient()
	endpointServer := NewEndpointServer(RuntimeSettings().EndpointServerHostAddress.Host)
	endpointServer.Start()
}
