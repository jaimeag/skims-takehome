package modules

import (
	"encoding/json"
	"net/http"

	top "github.com/jaimeag/skims-takehome/top"
)

type TopRouterModule struct {
	topModule *top.TopModule
}

func NewTopRouterModule() *TopRouterModule {
	return &TopRouterModule{
		topModule: top.NewTopModule(),
	}
}

func (m *TopRouterModule) TopFivePokemonInfoHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := m.topModule.GetTopFivePokemonInfo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		// TODO - handle error logging
		http.Error(w, "an issue occured while encoding response", http.StatusInternalServerError)
	}
}
