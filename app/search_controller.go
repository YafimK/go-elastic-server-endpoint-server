package main

import "net/http"

type SearchController struct {
	cache interface{} //TODO:
}

func (sc *SearchController) GetByField(responseWriter http.ResponseWriter, r *http.Request) {

}

func (sc *SearchController) SearchByString(responseWriter http.ResponseWriter, r *http.Request) {

}
