package main

import (
	"strings"
)

func cleanInput(text string) []string {

	lowerCase := strings.ToLower(text)
	cleanedInput := strings.Fields(lowerCase)

	return cleanedInput
}
