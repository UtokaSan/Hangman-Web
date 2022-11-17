package main

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

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
