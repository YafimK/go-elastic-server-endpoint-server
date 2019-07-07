package routes

import (
	"fmt"
	"github.com/YafimK/go-elastic-server-endpoint-server/config"
	"github.com/YafimK/go-elastic-server-endpoint-server/controllers"
	"github.com/YafimK/go-elastic-server-endpoint-server/elastic_service"
	"log"
	"net/http"
	"strings"
)

type SearchRoutingMap struct {
	elasticClient    *elastic_service.ElasticClient
	searchController *controllers.SearchController
	handlerMap       map[string]http.HandlerFunc
}

func NewSeachRoutingMap() *SearchRoutingMap {
	elasticClient, err := elastic_service.NewElasticClient(config.RuntimeSettings().ElasticServerAddress.String(), config.RuntimeSettings().ElasticServerIndex)
	if err != nil {
		log.Fatalf("Failed startiing elastic server client: %v\n", err)
	}
	searchController := controllers.NewSearchController(elasticClient)
	return &SearchRoutingMap{
		elasticClient:    elasticClient,
		searchController: searchController,
		handlerMap:       make(map[string]http.HandlerFunc),
	}
}

func (routes *SearchRoutingMap) AddHandler(pattern string, handlerFunc http.HandlerFunc) error {
	if _, isFound := routes.handlerMap[pattern]; isFound {
		return fmt.Errorf("pattern [%v] already has a registered handler", handlerFunc)
	}
	routes.handlerMap[pattern] = handlerFunc
	return nil
}

func (routes SearchRoutingMap) RegisterRoutes(mux *http.ServeMux) {
	err := routes.AddHandler("/search", controllers.Get(routes.searchController.SearchByString(routes.elasticClient.QueryAll)))
	if err != nil {
		log.Fatalf("error during registering routes: %v\n", err)
	}
	err = routes.AddHandler("/field", controllers.Get(routes.searchController.GetByField(routes.elasticClient.QueryByField)))
	if err != nil {
		log.Fatalf("error during registering routes: %v\n", err)
	}
	RegisterHandlers(routes.handlerMap, mux)
}

func (routes *SearchRoutingMap) String() string {
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
