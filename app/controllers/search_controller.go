package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/YafimK/go-elastic-server-endpoint-server/elastic_service"
	"github.com/YafimK/go-elastic-server-endpoint-server/model"
	"net/http"
)

type SearchController struct {
	cache         interface{} //TODO:
	elasticClient *elastic_service.ElasticClient
}

func NewSearchController(elasticClient *elastic_service.ElasticClient) *SearchController {
	return &SearchController{elasticClient: elasticClient}
}

func (sc *SearchController) GetByField(responseWriter http.ResponseWriter, r *http.Request) {
	typeParam := r.URL.Query().Get("type")
	if !isEmptyParam(typeParam) {
		HandleMissingHttpRequstParam(responseWriter, "type")
		return
	}
	if !validateTypeParamFieldValue(typeParam) {
		HandleBadRequestTypeParam(responseWriter, typeParam)
		return
	}
	valueParam := r.URL.Query().Get("value")
	if len(valueParam) == 0 {
		HandleMissingHttpRequstParam(responseWriter, "value")
		return
	}

	response, err := sc.elasticClient.QueryByField(typeParam, valueParam)
	if err != nil {
		http.Error(responseWriter, fmt.Sprintf("Failed getting search results from server %v", err), http.StatusInternalServerError)
	}
	err = respondAsJson(responseWriter, response)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}

func (sc *SearchController) SearchByString(responseWriter http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("s")
	if !isEmptyParam(param) {
		HandleMissingHttpRequstParam(responseWriter, "s")
		return
	}
	response, err := sc.elasticClient.QueryAll(param)
	if err != nil {
		http.Error(responseWriter, fmt.Sprintf("Failed getting search results from server %v", err), http.StatusInternalServerError)
	}

	err = respondAsJson(responseWriter, response)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}

func respondAsJson(responseWriter http.ResponseWriter, response model.Documents) error {
	marshaledResponse, err := json.Marshal(response)
	if err != nil {
		return err
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	_, err = responseWriter.Write(marshaledResponse)
	if err != nil {
		return err
	}
}

func HandleMissingHttpRequstParam(w http.ResponseWriter, paramName string) {
	http.Error(w, fmt.Sprintf("Missing mandatory request parameter in request - %v", paramName), http.StatusBadRequest)
}

func HandleBadRequestTypeParam(w http.ResponseWriter, paramValue string) {
	http.Error(w, fmt.Sprintf("Bad type param value in request, recieved - %v", paramValue), http.StatusBadRequest)
}

func isEmptyParam(value string) bool {
	if len(value) == 0 {
		return false
	}

	return true
}

func validateTypeParamFieldValue(value string) bool {
	for _, allowedValue := range model.TypeFieldsValues {
		if allowedValue == value {
			return true
		}
	}
	return false
}
