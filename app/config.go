package main

import (
	"flag"
	"sync"
)

var (
	settingsInit    sync.Once
	runtimeSettings *Settings
)

type Settings struct {
	Host                 string
	ElasticServerAddress string
}

func NewRuntimeSettings() *Settings {
	settings := Settings{}
	flag.StringVar(&settings.Host, "host", "localhost:8080", "Gateway server address (host:port)")
	flag.StringVar(&settings.ElasticServerAddress, "es_host", "http://localhost:9200", "Elastic Server Address host:port")
	flag.Parse()
	return &settings
}

func RuntimeSettings() *Settings {
	settingsInit.Do(func() {
		runtimeSettings = NewRuntimeSettings()
	})
	return runtimeSettings
}
