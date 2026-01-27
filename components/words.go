package components

import (
	_ "embed"
	"math/rand"
	"strings"
)

//go:embed english-snippets.txt
var englishSnippets string

//go:embed go-snippets.txt
var goSnippets string

//go:embed javascript-snippets.txt
var javascriptSnippets string

func GenerateWords(language string) string {
	var wordsFile string

	switch strings.ToLower(language) {
	case "go", "golang":
		wordsFile = goSnippets
	case "javascript", "js":
		wordsFile = javascriptSnippets
	case "english", "en":
		fallthrough
	default:
		wordsFile = englishSnippets
	}

	words := strings.Split(wordsFile, "\n")
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	return strings.Join(words, " ")
}
