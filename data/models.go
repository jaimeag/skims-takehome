package data

type Move struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type PokemonMove struct {
	// MVP only requires move key from api resp but nesting
	// is used to emulate resp structure
	Move Move `json:"move"`
}

type PokemonColor struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type PokemonApiResponse struct {
	Id          int           `json:"id"`
	Name        string        `json:"name"`
	Height      int           `json:"height"`
	Weight      int           `json:"weight"`
	Moves       []PokemonMove `json:"moves,omitempty"`
	RandomMoves []Move        `json:"random_moves"`
}

type PokemonApiSpeciesResponse struct {
	Id            int          `json:"id"`
	Name          string       `json:"name"`
	Color         PokemonColor `json:"color"`
	BaseHappiness int          `json:"base_happiness"`
}
