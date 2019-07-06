package main

import (
	elasticsearch "github.com/elastic/go-elasticsearch/v5"
	"log"
)

type ElasticClient struct {
	client *elasticsearch.Client
	index  string
}

func NewElasticClient(serverUrl string, index string) (*ElasticClient, error) {
	elasticsearchConfig := elasticsearch.Config{
		Addresses: []string{serverUrl},
	}
	escServer, err := elasticsearch.NewClient(elasticsearchConfig)
	if err != nil {
		return nil, err
	}
	log.Println(elasticsearch.Version)
	res, err := escServer.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	// Check response status
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}

	return &ElasticClient{
		escServer, index,
	}, nil
}
