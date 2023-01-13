package hangman_web

import (
	"strings"
)

func IsHangmanComplete(hangman HangmanData) bool {
	for _, letter := range hangman.Word {
		if !strings.Contains(hangman.RandomLetter, string(letter)) {
			return false
		}
	}
	return true
}

func Display(word string, letterRandom string) string {
	result := ""
	for _, letterOfWord := range word {
		if strings.Contains(letterRandom, string(letterOfWord)) {
			result += string(letterOfWord)
		} else {
			result += "_"
		}
	}
	return result
}

//LATER
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
