package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func cleanInput(text string) []string {

	lowerCase := strings.ToLower(text)
	cleanedInput := strings.Fields(lowerCase)

	return cleanedInput
}

func startRepl(cfg *config) {

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
				err := value.callback(cfg, cleaned[1:])
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					logToFile("errors.log", err) // A helper function to append errors to a file
				}
			} else {
				fmt.Println("Unknown command")
			}
		}

	}
}
func logToFile(filename string, err error) {
	// Open or create the file, and allow appending to it
	f, errOpen := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if errOpen != nil {
		fmt.Printf("Error logging to file: %v\n", errOpen)
		return
	}
	defer f.Close() // Ensure the file is properly closed when function exits

	// Write the error message with a timestamp
	logMessage := fmt.Sprintf("%v: %v\n", time.Now().Format(time.RFC3339), err)
	_, errWrite := f.WriteString(logMessage)
	if errWrite != nil {
		fmt.Printf("Error writing to log file: %v\n", errWrite)
	}
}
