package data

type PokemonApiResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
}

type PokemonApiSpeciesResponse struct {
	Id int `json:"id"`
}
