package main

import (
	"time"

	"github.com/bzelaznicki/pokedexcli/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: client,
	}
	startRepl(cfg)

}
