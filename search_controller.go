package main

import (
	"fmt"
	"net/http"
)

type SearchController struct {
	cache         interface{} //TODO:
	elasticClient *ElasticClient
}

func NewSearchController(elasticClient *ElasticClient) *SearchController {
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

	sc.elasticClient.QueryByField(typeParam, valueParam)

}

func (sc *SearchController) SearchByString(responseWriter http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	param := r.URL.Query().Get("s")
	if !isEmptyParam(param) {
		HandleMissingHttpRequstParam(responseWriter, "s")
		return
	}
	sc.elasticClient.QueryAll(param)
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
	for _, allowedValue := range TypeFieldsValues {
		if allowedValue == value {
			return true
		}
	}
	return false
}
