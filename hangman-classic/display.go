package hangman_web

import (
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

func Display(word string, letters string) string {
	result := ""
	for _, letter := range word {
		if strings.Contains(letters, string(letter)) {
			result = string(letter)
		} else {
			result = "_"
		}
	}
	return result
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
