package main

func main() {
	endpointServer := NewEndpointServer(RuntimeSettings().EndpointServerHostAddress.Host)
	endpointServer.Start()
}
