package hangman_web

import (
	"fmt"
	"os"
)

func game() {
	var hangman HangmanData
	file := os.Args[1]
	hangman.Word = Dictionnary(file)
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
}
