package main

import (
	"PeX/controllers"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Routes struct {
	searchController *controllers.SearchController
	handlerMap       map[string]http.HandlerFunc
}

func NewRoutes(searchController *SearchController) *Routes {
	return &Routes{
		searchController: searchController,
		handlerMap:       make(map[string]http.HandlerFunc),
	}
}

func (routes *Routes) AddHandler(pattern string, handlerFunc http.HandlerFunc) error {
	if _, isFound := routes.handlerMap[pattern]; isFound {
		return fmt.Errorf("pattern [%v] already has a registered handler", handlerFunc)
	}
	routes.handlerMap[pattern] = handlerFunc
	return nil
}

func (routes Routes) RegisterRoutes(mux *http.ServeMux) {
	err := routes.AddHandler("/search", controllers.Get(routes.searchController.SearchAllByString))
	if err != nil {
		log.Fatalf("error during registering routes: %v\n", err)
	}
	err = routes.AddHandler("/field", controllers.Get(routes.searchController.GetByField))
	if err != nil {
		log.Fatalf("error during registering routes: %v\n", err)
	}
	RegisterHandlers(routes.handlerMap, mux)
}

func (routes *Routes) String() string {
	message := strings.Builder{}
	message.WriteString("Current registered patterns:\n")
	message.WriteString("------------------------------\n")
	for pattern := range routes.handlerMap {
		message.WriteString(fmt.Sprintf("%s\n", pattern))
	}
	return message.String()
}

func RegisterHandlers(handlerMap map[string]http.HandlerFunc, mux *http.ServeMux) {
	for pattern, handler := range handlerMap {
		mux.HandleFunc(pattern, handler)
	}
}
