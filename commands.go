package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/bzelaznicki/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Lists the next 20 map locations",
			callback:    commandMapn,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the previous 20 map locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Shows you more details about a specific location. Usage: explore location_name",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "attempts to catch a Pokemon. Usage: catch pokemon_name",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspects a pokemon in your Pokedex. Usage inspect pokemon_name",
			callback:    commandInspect,
		},
	}
}

func getSingleParam(params []string) (string, error) {
	if len(params) != 1 {
		return "", fmt.Errorf("you must provide a single argument")
	}
	return params[0], nil
}
func commandHelp(cfg *config, params []string) error {
	fmt.Println("Usage:")

	for commandName, command := range getCommands() {
		fmt.Printf("%s: %s\n", commandName, command.description)
	}
	return nil

}

func commandExit(cfg *config, params []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMapn(cfg *config, params []string) error {

	var urlPtr *string
	if cfg.nextUrl != "" {
		urlPtr = &cfg.nextUrl
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(urlPtr)
	if err != nil {
		return err
	}
	updateConfigUrls(cfg, &locationsResp)

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config, params []string) error {

	if cfg.previousUrl == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	locationsResp, err := cfg.pokeapiClient.ListLocations(&cfg.previousUrl)
	if err != nil {
		return err
	}

	updateConfigUrls(cfg, &locationsResp)

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandExplore(cfg *config, params []string) error {
	locationName, err := getSingleParam(params)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", locationName)

	locData, err := cfg.pokeapiClient.GetLocationArea(locationName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range locData.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config, params []string) error {
	pokemonName, err := getSingleParam(params)
	if err != nil {
		return err
	}
	if _, exists := cfg.pokedex[pokemonName]; exists {
		fmt.Printf("You already caught %s!\n", pokemonName)
		return nil
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemonInfo, err := cfg.pokeapiClient.GetPokemonInfo(pokemonName)
	if err != nil {
		return err
	}

	rand.Seed(time.Now().UnixNano())

	// Calculate the chance and cap it between 10% and 90%
	chance := 100 - pokemonInfo.BaseExp/10

	// Clamp chance within the min/max range
	if chance < minCatchChance {
		chance = minCatchChance
	} else if chance > maxCatchChance {
		chance = maxCatchChance
	}

	// Generate a random number and determine the result
	r := rand.Intn(100)
	if r < chance {
		fmt.Printf("Caught %s!\n", pokemonName)
		// Add the caught Pokemon to the Pokedex here
		cfg.pokedex[pokemonName] = pokemonInfo
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}

func commandInspect(cfg *config, params []string) error {
	pokemonName, err := getSingleParam(params)
	if err != nil {
		return err
	}

	poke, exists := cfg.pokedex[pokemonName]
	if !exists {
		if _, err := cfg.pokeapiClient.GetPokemonInfo(pokemonName); err != nil {
			return err
		}
		fmt.Printf("you have not caught %s yet!\n", pokemonName)
		return nil
	}

	fmt.Printf("Name: %s\n", poke.Name)
	fmt.Printf("Height: %d\n", poke.Height)
	fmt.Printf("Weight: %d\n", poke.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range poke.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, pkType := range poke.Types {
		fmt.Printf("  - %s\n", pkType.Type.Name)
	}

	return nil
}

func updateConfigUrls(cfg *config, locationsResp *pokeapi.LocationAreaResp) {
	if locationsResp.Next != nil {
		cfg.nextUrl = *locationsResp.Next
	} else {
		cfg.nextUrl = ""
	}

	if locationsResp.Previous != nil {
		cfg.previousUrl = *locationsResp.Previous
	} else {
		cfg.previousUrl = ""
	}
}
