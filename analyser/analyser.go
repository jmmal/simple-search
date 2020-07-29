package analyser

import (
	"github.com/jmmal/simple-search/filters"
	"strings"
)

// Tokenize takes a string value and returns a list of terms
func Tokenize(value string) []string {
	filtered := filters.Lowercase(value)
	tokens := strings.Fields(filtered)
	nonStops := filters.RemoveEnglishStopWords(tokens)

	return nonStops
}