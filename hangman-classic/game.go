package hangman_web

func Game(data string, word string) string {
	var hangman HangmanData
	hangman.Word = word
	hangman.RandomLetter = string(hangman.Word[len(hangman.Word)/2-1])
	hangman.RandomLetter += data
	return Display(hangman.Word, hangman.RandomLetter)
}
