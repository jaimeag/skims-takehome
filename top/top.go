package top

import (
	"math"
	"sort"

	data "github.com/jaimeag/skims-takehome/data"
)

type TopModule struct {
	topProvider *data.TopProvider
}

type PokemonAndSpeciesInfo struct {
	data.PokemonApiSpeciesResponse
	data.PokemonApiResponse
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type FavoritePokemonInfo struct {
	FavoritePokemon      []PokemonAndSpeciesInfo `json:"favorite_pokemon"`
	BaseHappinessAverage float64                 `json:"base_happiness_average"`
	BaseHappinessMedian  float64                 `json:"base_happiness_median"`
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
func (m *TopModule) GetTopFivePokemonInfo() (*FavoritePokemonInfo, error) {
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
	// sort in descending order by base_happiness
	sort.Slice(pokemon, func(i, j int) bool {
		return pokemon[i].BaseHappiness < pokemon[j].BaseHappiness
	})

	average := m.getAverageFromPokemonBaseHappiness(pokemon)

	median := m.getMedianFromPokemonBaseHappiness(pokemon)

	return &FavoritePokemonInfo{
		FavoritePokemon:      pokemon,
		BaseHappinessAverage: average,
		BaseHappinessMedian:  median,
	}, nil
}

func (m *TopModule) getPokemonAndSpeciesInfo(pokemonId int) (*PokemonAndSpeciesInfo, error) {

	pokemonInfo, err := m.topProvider.GetPokemonInfo(pokemonId)
	if err != nil {
		return nil, err
	}
	randomizedMoves := m.topProvider.GetRandomizedPokemonMovesFromPokemonInfo(pokemonInfo)

	// get first 2 moves from randomized list
	pokemonInfo.RandomMoves = randomizedMoves[0:2]

	speciesInfo, err := m.topProvider.GetPokemonSpeciesInfo(pokemonId)
	if err != nil {
		return nil, err
	}

	return &PokemonAndSpeciesInfo{
		*speciesInfo,
		*pokemonInfo,
		pokemonInfo.Id,
		pokemonInfo.Name,
	}, nil

}

// THESE FUNCTIONS ASSUME THAT THE SLICE IS SORTED

func (m *TopModule) getMedianFromPokemonBaseHappiness(pokemon []PokemonAndSpeciesInfo) float64 {
	idxA := math.Floor(float64(len(pokemon) / 2))
	idxAInt := int(idxA)
	if len(pokemon)%2 == 0 {
		idxB := idxAInt - 1
		return (float64(pokemon[idxAInt].BaseHappiness) + float64(pokemon[idxB].BaseHappiness)) / 2
	} else {
		return float64(pokemon[idxAInt].BaseHappiness)
	}
}

func (m *TopModule) getAverageFromPokemonBaseHappiness(pokemon []PokemonAndSpeciesInfo) float64 {
	numerator := 0
	denominator := len(pokemon)

	for _, pokemonInfo := range pokemon {
		numerator += pokemonInfo.BaseHappiness
	}

	return float64(numerator) / float64(denominator)
}

// ? - mean and average are interchangeable
// func getMeanFromPokemonBaseHappiness(pokemon []PokemonAndSpeciesInfo) float32 {

// }
