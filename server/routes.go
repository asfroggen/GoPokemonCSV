package server

import (
	"encoding/json"
	"fmt"
	"main/service"
	"net/http"
)

func initRoutes() {

	pokemonService := service.NewPokemonService()

	http.HandleFunc("/", index)

	http.HandleFunc("/loadPokemon", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			err := pokemonService.LoadPokemon()
			if err != nil {
				w.WriteHeader(http.StatusConflict)
				fmt.Fprintf(w, "Something went wrong")
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Pokemon list loaded succesfully")

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}

	})

	http.HandleFunc("/getPokemon", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			pokemonList, err := pokemonService.GetPokemon()
			if err != nil {
				w.WriteHeader(http.StatusConflict)
				fmt.Fprintf(w, "Something went wrong")
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(pokemonList)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}

	})
}
