package modules

import (
	data "github.com/jaimeag/skims-takehome/data"
)

type TopModule struct{}

var (
	favoritePokemonIds = [5]int{1, 2, 3, 4, 5}
)

func NewTopModule() *TopModule {
	return &TopModule{}
}

func (m *TopModule) GetTopFivePokemonInfo() {
	// get pokemon data for five pokemon
	info, err := data.GetPokemonInfo(479)
	if err != nil {
	}
	// for each pokemon, call a function parse-data that parses the response according to the
	// challenge

	// once done, write response
}
