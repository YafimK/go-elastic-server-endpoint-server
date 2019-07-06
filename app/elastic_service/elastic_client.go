package elastic_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	elasticsearch "github.com/elastic/go-elasticsearch/v5"
	"log"
)

type ElasticClient struct {
	elasticClient *elasticsearch.Client
	index         string
}

func (client ElasticClient) QueryAll(searchValue string) ([]interface{}, error) {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"query": fmt.Sprintf("*%v*", searchValue),
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	res, err := client.elasticClient.Search(
		client.elasticClient.Search.WithPretty(),
		client.elasticClient.Search.WithBody(&buf),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	var parsedResponse map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&parsedResponse); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	fmt.Println(parsedResponse)
	// Print the ID and document source for each hit.
	var results []interface{}
	for _, hit := range parsedResponse["hits"].(map[string]interface{})["hits"].([]interface{}) {
		results = append(results, hit.(map[string]interface{})["_source"])
	}

	return results, nil
}

func (client ElasticClient) QueryByField(fieldType string, fieldValue string) (string, error) {
	return "", nil
}

func NewElasticClient(serverUrl string, index string) (*ElasticClient, error) {
	elasticsearchConfig := elasticsearch.Config{
		Addresses: []string{serverUrl},
	}
	escServer, err := elasticsearch.NewClient(elasticsearchConfig)
	if err != nil {
		return nil, err
	}
	log.Printf("Starting elastic server elasticClient with elastic server version - %v", elasticsearch.Version)
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
