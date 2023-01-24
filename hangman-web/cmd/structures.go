package hangman_web

type HangmanWeb struct {
	Word      string
	Life      int
	Display   string
	Input     string
	InputUse  bool
	LetterUse string
}

type Account struct {
	Mail     string
	Password string
}
type LeaderBoard struct {
	name string
}
