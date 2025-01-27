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
	fmt.Println("Welcome to the Pokédex!")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			cleaned := cleanInput(scanner.Text())
			if len(cleaned) == 0 {
				continue
			}
			fmt.Println("Your command was:", cleaned[0])
		}

	}
}
