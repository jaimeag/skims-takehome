package modules

import (
	data "github.com/jaimeag/skims-takehome/data"
)

type TopModule struct {
	topProvider *data.TopProvider
}

type PokemonAndSpeciesInfo struct {
	data.PokemonApiResponse
	data.PokemonApiSpeciesResponse
}

var (
	favoritePokemonIds = [5]int{1, 2, 3, 4, 5}
)

func NewTopModule() *TopModule {
	return &TopModule{
		topProvider: data.NewTopProvider(),
	}
}

func (m *TopModule) GetTopFivePokemonInfo() {
	// get pokemon data for five pokemon
	// info, err := data.GetPokemonInfo(479)
	// if err != nil {
	// }
	// for each pokemon, call a function parse-data that parses the response according to the
	// challenge

	// once done, write response
}

func (m *TopModule) GetPokemonAndSpeciesInfo(pokemonId int) (*PokemonAndSpeciesInfo, error) {

	pokemonInfo, err := m.topProvider.GetPokemonInfo(pokemonId)
	if err != nil {
		return nil, err
	}

	speciesInfo, err := m.topProvider.GetPokemonSpeciesInfo(pokemonId)
	if err != nil {
		return nil, err
	}

	return &PokemonAndSpeciesInfo{
		*pokemonInfo,
		*speciesInfo,
	}, nil

}
