package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {

	lowerCase := strings.ToLower(text)
	cleanedInput := strings.Fields(lowerCase)

	return cleanedInput
}

func startRepl() {

	fmt.Println("Welcome to the Pokedex!")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			cleaned := cleanInput(scanner.Text())
			if len(cleaned) == 0 {
				continue
			}
			command := cleaned[0]
			if value, exists := getCommands()[command]; exists {
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
