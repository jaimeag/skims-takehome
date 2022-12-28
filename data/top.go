package data

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	pokemonApiEndpoint        = "https://pokeapi.co/api/v2/pokemon/"
	pokemonSpeciesApiEndpoint = "https://pokeapi.co/api/v2/pokemon-species/"
)

type TopProvider interface {
	GetPokemonAndSpeciesInfo(pokemonId string) (*PokemonApiResponse, error)
}

// type PokemonSpeciesApiResponse {

// }

func GetPokemonInfo(pokemonId int) (*PokemonApiResponse, error) {
	pokemonEndpoint := fmt.Sprintf("%s%d", pokemonApiEndpoint, pokemonId)
	fmt.Print(pokemonEndpoint)
	// speciesEndpoint := fmt.Sprintf("%s%s", pokemonSpeciesApiEndpoint, pokemonId)
	// call api with pokemonId for pokmeon info
	r, err := http.Get(pokemonEndpoint)
	fmt.Printf("%+v\n", r.Body)
	if err != nil {
		return nil, err
	}
	payload := &PokemonApiResponse{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, err
	}
	return payload, nil
	// call api with id for species info
}

func GetPokemonSpeciesInfo(pokemonId int) error {

}
