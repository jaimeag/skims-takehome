package modules

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	topRouterHandler := NewTopRouterModule()
	router.Handle("/pokemon/top-five", http.HandlerFunc(topRouterHandler.TopFivePokemonInfoHandler)).Methods(http.MethodGet)

	return router
}
