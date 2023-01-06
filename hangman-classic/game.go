package hangman_web

import (
	"fmt"
)

func Game() string {
	var hangman HangmanData
	hangman.RandomLetter = string(hangman.Word[len(hangman.Word)/2-1])
	for !IsHangmanComplete(hangman.Word, hangman.RandomLetter) {
		fmt.Print(hangman.Word)
		fmt.Print("\nChoose : \n")
		var input string
		fmt.Scanf("%s", &input)
		hangman.RandomLetter += input
		Display(hangman.Word, hangman.RandomLetter)
	}
	fmt.Print("\n\nCongrats !")
	return Display(hangman.Word, hangman.RandomLetter)
}
