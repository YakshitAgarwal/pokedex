package main

import (
	"time"

	"github.com/YakshitAgarwal/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
	caughtPokemon           map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}
	startRepl(&cfg)
}
