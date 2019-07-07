package controllers

import (
	"fmt"
	"github.com/YafimK/go-elastic-server-endpoint-server/common"
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

type ElasticClientQueryByField func(typeParam string, valueParam string) (model.Documents, error)

func (sc *SearchController) GetByField(elasticClientQueryByField ElasticClientQueryByField) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, r *http.Request) {
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

		response, err := elasticClientQueryByField(typeParam, valueParam)
		if err != nil {
			http.Error(responseWriter, fmt.Sprintf("Failed getting search results from server %v", err), http.StatusInternalServerError)
		}
		err = common.RespondAsJson(responseWriter, response)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		}
	}
}

type ElasticClientQueryAllCallback func(value string) (model.Documents, error)

func (sc *SearchController) SearchByString(elasticClientQueryAllCallback ElasticClientQueryAllCallback) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("s")
		if !isEmptyParam(param) {
			HandleMissingHttpRequstParam(responseWriter, "s")
			return
		}
		response, err := elasticClientQueryAllCallback(param)
		if err != nil {
			http.Error(responseWriter, fmt.Sprintf("Failed getting search results from server %v", err), http.StatusInternalServerError)
		}

		err = common.RespondAsJson(responseWriter, response)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		}
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
