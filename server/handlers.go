package server

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}

	fmt.Fprintf(w, "Hello visitor, \nUse /loadPokemon to load the pokemons to the CSV \nUse /getPokemon to fetch the pokemon list from the CSV")
}
