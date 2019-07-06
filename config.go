package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"sync"
)

var (
	settingsInit    sync.Once
	runtimeSettings *Settings
)

type Settings struct {
	EndpointServerHostAddress *url.URL
	ElasticServerAddress      *url.URL
}

func BadArgumentError(argumentName string, err error) error {
	return fmt.Errorf("recieved argument [%v] is in wrong format: %v", argumentName, err)
}

func NewRuntimeSettings() *Settings {
	settings := Settings{}
	host := flag.String("host", "http://localhost:8080", "Gateway server address <host:port>")
	elasticServerAddress := flag.String("es_host", "http://localhost:9200", "Elastic Server Address <protocol://host:port>")
	flag.Parse()
	var err error
	settings.EndpointServerHostAddress, err = parseUrl(host, false, false, true)
	if err != nil {
		log.Fatal(BadArgumentError("host", err))
	}
	settings.ElasticServerAddress, err = parseUrl(elasticServerAddress, false, false, true)
	if err != nil {
		log.Fatal(BadArgumentError("es_host", err))
	}
	return &settings
}

func RuntimeSettings() *Settings {
	settingsInit.Do(func() {
		runtimeSettings = NewRuntimeSettings()
	})
	return runtimeSettings
}
