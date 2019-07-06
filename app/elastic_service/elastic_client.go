package elastic_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/YafimK/go-elastic-server-endpoint-server/model"
	elasticsearch "github.com/elastic/go-elasticsearch/v5"
	"log"
)

type ElasticClient struct {
	elasticClient *elasticsearch.Client
	index         string
}

func (client ElasticClient) QueryAll(searchValue string) (model.Documents, error) {
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
			return nil, fmt.Errorf("Error parsing the response body: %s", err)
		} else {
			return nil, fmt.Errorf("[%s] %s: %s",
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
	var results model.Documents
	for _, hit := range parsedResponse["hits"].(map[string]interface{})["hits"].([]interface{}) {
		result := hit.(map[string]interface{})["_source"].(map[string]interface{})
		newDoc := model.Document{
			Ip:            result["ip"].(string),
			Timestamp:     result["timestamp"].(string),
			Domain:        result["domain"].(string),
			IsBlacklisted: result["blacklisted"].(bool),
			EventType:     result["event_type"].(string),
		}
		results = append(results, newDoc)
	}

	return results, nil
}

func (client ElasticClient) QueryByField(fieldType string, fieldValue string) (model.Documents, error) {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"fields": []string{fieldType},
				"query":  fieldValue,
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
			return nil, fmt.Errorf("Error parsing the response body: %s", err)
		} else {
			return nil, fmt.Errorf("[%s] %s: %s",
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
	var results model.Documents
	for _, hit := range parsedResponse["hits"].(map[string]interface{})["hits"].([]interface{}) {
		result := hit.(map[string]interface{})["_source"].(map[string]interface{})
		newDoc := model.Document{
			Ip:            result["ip"].(string),
			Timestamp:     result["timestamp"].(string),
			Domain:        result["domain"].(string),
			IsBlacklisted: result["blacklisted"].(bool),
			EventType:     result["event_type"].(string),
		}
		results = append(results, newDoc)
	}

	return results, nil
}

func NewElasticClient(serverUrl string, index string) (*ElasticClient, error) {
	elasticsearchConfig := elasticsearch.Config{
		Addresses: []string{serverUrl},
	}
	escServer, err := elasticsearch.NewClient(elasticsearchConfig)
	if err != nil {
		return nil, err
	}
	log.Printf("Starting elastic server; version - %v", elasticsearch.Version)
	res, err := escServer.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	// Check response status
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	log.Printf("Elastic server info: - %v", res)

	return &ElasticClient{
		escServer, index,
	}, nil
}
