package main

import (
	"github.com/bzelaznicki/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	previousUrl   string
	nextUrl       string
	pokedex       map[string]pokeapi.Pokemon
}

const (
	minCatchChance = 10
	maxCatchChance = 90
)
