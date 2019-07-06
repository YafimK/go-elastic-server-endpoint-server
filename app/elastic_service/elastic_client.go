package elastic_service

import (
	elasticsearch "github.com/elastic/go-elasticsearch/v5"
	"log"
)

type ElasticClient struct {
	client *elasticsearch.Client
	index  string
}

func (client ElasticClient) QueryAll(s string) string {
	return ""
}

func (client ElasticClient) QueryByField(fieldType string, fieldValue string) string {
	return ""
}

func NewElasticClient(serverUrl string, index string) (*ElasticClient, error) {
	elasticsearchConfig := elasticsearch.Config{
		Addresses: []string{serverUrl},
	}
	escServer, err := elasticsearch.NewClient(elasticsearchConfig)
	if err != nil {
		return nil, err
	}
	log.Printf("Starting elastic server client with elastic server version - %v", elasticsearch.Version)
	res, err := escServer.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	// Check response status
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	log.Printf("Elastic server: - %v", res)

	return &ElasticClient{
		escServer, index,
	}, nil
}
