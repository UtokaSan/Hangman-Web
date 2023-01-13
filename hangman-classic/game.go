package hangman_web

import "fmt"

func Game(data string, word string) string {
	var hangman HangmanData
	hangman.Word = word
	hangman.RandomLetter = string(hangman.Word[len(hangman.Word)/2-1])
	if !IsHangmanComplete(hangman) {
		fmt.Println("OK")
		hangman.RandomLetter += data
		return Display(hangman.Word, hangman.RandomLetter)
	}
	return Display(hangman.Word, hangman.RandomLetter)
}
