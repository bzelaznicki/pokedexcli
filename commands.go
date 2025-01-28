package main

import (
	"fmt"
	"os"

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
	}
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
	if len(params) != 1 {
		return fmt.Errorf("you must provide a location name")
	}

	locationName := params[0]
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
