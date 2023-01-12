package hangman_web

func Game(data string, word string) string {
	var hangman HangmanData
	var randomLetter = string(word[len(word)/2-1])
	for !IsHangmanComplete(hangman.Word, hangman.RandomLetter) {
		hangman.RandomLetter += data
		Display(word, randomLetter)
	}
	return Display(word, randomLetter)
}
