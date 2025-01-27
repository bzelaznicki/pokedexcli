package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commands map[string]cliCommand

func cleanInput(text string) []string {

	lowerCase := strings.ToLower(text)
	cleanedInput := strings.Fields(lowerCase)

	return cleanedInput
}

func startRepl() {
	commands = map[string]cliCommand{
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
	}
	fmt.Println("Welcome to the Pokedex!")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			cleaned := cleanInput(scanner.Text())
			if len(cleaned) == 0 {
				continue
			}
			if value, exists := commands[cleaned[0]]; exists {
				err := value.callback()
				if err != nil {
					fmt.Printf("Error: %v\n", err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}

	}
}

func commandHelp() error {
	fmt.Println("Usage:")

	for commandName, command := range commands {
		fmt.Printf("%s: %s\n", commandName, command.description)
	}
	return nil

}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
