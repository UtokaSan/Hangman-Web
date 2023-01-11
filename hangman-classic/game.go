package hangman_web

import (
	"fmt"
)

func Game(data string) string {
	var hangman HangmanData
	hangman.RandomLetter = string(hangman.Word[len(hangman.Word)/2-1])
	for !IsHangmanComplete(hangman.Word, hangman.RandomLetter) {
		fmt.Print(hangman.Word)
		hangman.RandomLetter += data
		Display(hangman.Word, hangman.RandomLetter)
	}
	fmt.Print("\n\nCongrats !")
	return Display(hangman.Word, hangman.RandomLetter)
}
