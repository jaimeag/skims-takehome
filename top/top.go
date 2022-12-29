package top

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
	favoritePokemonIds = [5]int{479, 350, 491, 488, 385}
)

func NewTopModule() *TopModule {
	return &TopModule{
		topProvider: data.NewTopProvider(),
	}
}

// TODO - use routines to fetch data async
func (m *TopModule) GetTopFivePokemonInfo() ([]PokemonAndSpeciesInfo, error) {
	pokemon := []PokemonAndSpeciesInfo{}
	// get pokemon data for five pokemon
	for _, id := range favoritePokemonIds {
		pokemonInfo, err := m.getPokemonAndSpeciesInfo(id)
		// remove non randomized moves
		pokemonInfo.Moves = nil
		if err != nil {
			return nil, err
		}
		pokemon = append(pokemon, *pokemonInfo)
	}
	return pokemon, nil
	// info, err := data.GetPokemonInfo(479)
	// if err != nil {
	// }
	// for each pokemon, call a function parse-data that parses the response according to the
	// challenge

	// once done, write response
}

func (m *TopModule) getPokemonAndSpeciesInfo(pokemonId int) (*PokemonAndSpeciesInfo, error) {

	pokemonInfo, err := m.topProvider.GetPokemonInfo(pokemonId)
	if err != nil {
		return nil, err
	}
	randomziedMoves := m.topProvider.GetRandomizedPokemonMovesFromPokemonInfo(pokemonInfo)

	pokemonInfo.RandomMoves = randomziedMoves[0:2]

	speciesInfo, err := m.topProvider.GetPokemonSpeciesInfo(pokemonId)
	if err != nil {
		return nil, err
	}

	return &PokemonAndSpeciesInfo{
		*pokemonInfo,
		*speciesInfo,
	}, nil

}
