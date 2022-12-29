package modules

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(logger *log.Logger) *mux.Router {
	router := mux.NewRouter()
	topRouterHandler := NewTopRouterModule()
	router.Handle("/pokemon/top-five", http.HandlerFunc(topRouterHandler.TopFivePokemonInfoHandler)).Methods(http.MethodGet)

	return router
}
