package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type HangmanData struct {
	Word         string
	RandomLetter string
	InputMenu    int
}

func Random(nbr int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(nbr)
}

func Dictionnary(fileAccess string) string {
	data, _ := ioutil.ReadFile(fileAccess)
	str := string(data)
	split := strings.Split(str, "\n")
	random := Random(len(split))
	word := split[random]
	return word
}

func PrintMenu() {
	fmt.Println("---- [Menu] ----")
	fmt.Println("1. Play\n2. Rules")
	fmt.Print("Choose : \n")
}

func main() {
	if len(os.Args) > 2 {
		fmt.Print("Too many arguments")
		return
	}
	var hangman HangmanData
	PrintMenu()
	for hangman.InputMenu != 1 {
		hangman.InputMenu = 0
		fmt.Scanf("%d", &hangman.InputMenu)
		if hangman.InputMenu == 1 {
			game()
		}
		if hangman.InputMenu == 2 {
			rules()
			PrintMenu()
		}
	}
}

func rules() {
	fmt.Println("---- [Rules] ----")
	fmt.Println("Hangman is a simple word guessing game. Players try to figure out an unknown word by guessing letters. ")
	fmt.Println("If too many letters which do not appear in the word are guessed, the player is hanged (and loses)")
}
