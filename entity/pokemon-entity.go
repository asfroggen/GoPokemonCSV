package entity

type PokemonResponse struct {
	Count   int        `json:"count"`
	Results []*Pokemon `json:"results"`
}

type Pokemon struct {
	Name string `json:"name"`
}
