package main

import (
	"github.com/bzelaznicki/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	previousUrl   string
	nextUrl       string
}
