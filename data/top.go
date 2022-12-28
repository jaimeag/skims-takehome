package data

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const (
	pokemonApiEndpoint        = "https://pokeapi.co/api/v2/pokemon/"
	pokemonSpeciesApiEndpoint = "https://pokeapi.co/api/v2/pokemon-species/"
)

type TopProvider struct {
}

func NewTopProvider() *TopProvider {
	return &TopProvider{}
}

func (p *TopProvider) GetPokemonInfo(pokemonId int) (*PokemonApiResponse, error) {
	pokemonEndpoint := fmt.Sprintf("%s%d", pokemonApiEndpoint, pokemonId)

	// call api with pokemonId for pokmeon info
	r, err := http.Get(pokemonEndpoint)
	if err != nil {
		return nil, err
	}
	payload := &PokemonApiResponse{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, err
	}
	return payload, nil
}

func (p *TopProvider) GetPokemonSpeciesInfo(pokemonId int) (*PokemonApiSpeciesResponse, error) {
	// call api with id for species info
	speciesEndpoint := fmt.Sprintf("%s%d", pokemonSpeciesApiEndpoint, pokemonId)
	r, err := http.Get(speciesEndpoint)
	if err != nil {
		return nil, err
	}
	payload := &PokemonApiSpeciesResponse{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, err
	}
	return payload, nil
}

func (p *TopProvider) GetRandomizedPokemonMoves(pokemonInfo *PokemonApiResponse) []Move {
	rand.Seed(time.Now().UnixNano())
	movesToRandomize := []Move{}

	// flatten moves from nested structure
	for _, pokemonMove := range pokemonInfo.Moves {
		movesToRandomize = append(movesToRandomize, pokemonMove.Move)
	}

	// shuffle movesToRandomize
	rand.Shuffle(len(movesToRandomize), func(i, j int) {
		movesToRandomize[i], movesToRandomize[j] = movesToRandomize[j], movesToRandomize[i]
	})

	return movesToRandomize
}
