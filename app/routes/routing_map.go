package routes

import (
	"fmt"
	"github.com/YafimK/go-elastic-server-endpoint-server/controllers"
	"log"
	"net/http"
	"strings"
)

type RoutingMap struct {
	searchController *controllers.SearchController
	handlerMap       map[string]http.HandlerFunc
}

func NewRoutingMap(searchController *controllers.SearchController) *RoutingMap {
	return &RoutingMap{
		searchController: searchController,
		handlerMap:       make(map[string]http.HandlerFunc),
	}
}

func (routes *RoutingMap) AddHandler(pattern string, handlerFunc http.HandlerFunc) error {
	if _, isFound := routes.handlerMap[pattern]; isFound {
		return fmt.Errorf("pattern [%v] already has a registered handler", handlerFunc)
	}
	routes.handlerMap[pattern] = handlerFunc
	return nil
}

func (routes RoutingMap) RegisterRoutes(mux *http.ServeMux) {
	err := routes.AddHandler("/search", controllers.Get(routes.searchController.SearchByString))
	if err != nil {
		log.Fatalf("error during registering routes: %v\n", err)
	}
	err = routes.AddHandler("/field", controllers.Get(routes.searchController.GetByField))
	if err != nil {
		log.Fatalf("error during registering routes: %v\n", err)
	}
	RegisterHandlers(routes.handlerMap, mux)
}

func (routes *RoutingMap) String() string {
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
