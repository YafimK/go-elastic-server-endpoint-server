package config

import (
	"flag"
	"fmt"
	"github.com/YafimK/go-elastic-server-endpoint-server/common"
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
	ElasticServerIndex        string
}

func badArgumentError(argumentName string, err error) error {
	return fmt.Errorf("recieved argument [%v] is in wrong format: %v", argumentName, err)
}

func newRuntimeSettings() *Settings {
	settings := Settings{}
	host := flag.String("host", "http://localhost:8080", "Gateway server address <host:port>")
	elasticServerAddress := flag.String("es_host", "http://localhost:9200", "Elastic Server Address <protocol://host:port>")
	settings.ElasticServerIndex = *flag.String("es_index", "page-views", "Elastic Server index")
	flag.Parse()

	var err error
	settings.EndpointServerHostAddress, err = common.ParseUrl(host, false, false, true)
	if err != nil {
		log.Fatal(badArgumentError("host", err))
	}
	settings.ElasticServerAddress, err = common.ParseUrl(elasticServerAddress, false, false, true)
	if err != nil {
		log.Fatal(badArgumentError("es_host", err))
	}
	return &settings
}

func RuntimeSettings() *Settings {
	settingsInit.Do(func() {
		runtimeSettings = newRuntimeSettings()
	})
	return runtimeSettings
}
