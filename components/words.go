package components

import (
	_ "embed"
	"math/rand"
	"strings"
)

//go:embed 1-1000.txt
var wordsFile string

func GenerateWords() string {
	words := strings.Split(wordsFile, "\n")
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	return strings.Join(words, " ")
}
