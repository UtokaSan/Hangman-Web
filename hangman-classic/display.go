package main

import (
	"fmt"
	"strings"
)

func IsHangmanComplete(word string, letters string) bool {
	for _, letter := range word {
		if !strings.Contains(letters, string(letter)) {
			return false
		}
	}
	return true
}

func Display(word string, letters string) {
	for _, letter := range word {
		if strings.Contains(letters, string(letter)) {
			fmt.Print(string(letter))
		} else {
			fmt.Print("_")
		}
	}
}
func IsInputValid(word string, input string) bool {
	if len(input) == 1 {
		for _, letter := range word {
			if string(letter) == input {
				return true
			}
		}
	} else {
		return input == word
	}
	return false
}
