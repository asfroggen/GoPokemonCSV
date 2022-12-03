package service

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"main/entity"
	"net/http"
	"os"
)

type PokemonService interface {
	GetPokemon() ([]entity.Pokemon, error)
	LoadPokemon() error
}

type service struct{}

func NewPokemonService() PokemonService {
	return &service{}
}

func (*service) GetPokemon() ([]entity.Pokemon, error) {
	file, err := os.Open("pokemon.csv")

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	var pokemonList []entity.Pokemon
	for _, record := range records {
		data := entity.Pokemon{
			Name: record[0],
		}
		pokemonList = append(pokemonList, data)
	}

	return pokemonList, err
}
func (*service) LoadPokemon() error {
	url := "https://pokeapi.co/api/v2/pokemon/"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	pokemonResponse := &entity.PokemonResponse{}

	err = json.NewDecoder(resp.Body).Decode(pokemonResponse)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(pokemonResponse)

	savePokemon(pokemonResponse.Results)
	return nil
}

func savePokemon(column []*entity.Pokemon) {

	file, err := os.Create("pokemon.csv")
	if err != nil {
		log.Fatalln("Failed to open file", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	// Write all the pokemon in CSV file
	var data [][]string
	for _, pokemon := range column {
		row := []string{pokemon.Name}
		data = append(data, row)
	}
	w.WriteAll(data)
}
