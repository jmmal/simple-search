package filters

import (
	"strings"
)

// RemoveEnglishStopWords removes common terms like "and", "or" from
// a slice of tokens
func RemoveEnglishStopWords(tokens []string) []string {
	isNotStopWord := func(s string) bool {
		for _, word := range getEnglishStopWords() {
			if word == s {
				return false
			}
		}
		return true
	}

	return filter(tokens, isNotStopWord)
}

// Lowercase converts the value to lowercase
func Lowercase(v string) string {
	return strings.ToLower(v)
}

func filter(ss []string, test func(string) bool) (ret []string) {
    for _, s := range ss {
        if test(s) {
            ret = append(ret, s)
        }
    }
    return
}

func getEnglishStopWords() []string{
	return []string{
		"i",
		"me",
		"my",
		"myself",
		"we",
		"our",
		"ours",
		"ourselves",
		"you",
		"your",
		"yours",
		"yourself",
		"yourselves",
		"he",
		"him",
		"his",
		"himself",
		"she",
		"her",
		"hers",
		"herself",
		"it",
		"its",
		"itself",
		"they",
		"them",
		"their",
		"theirs",
		"themselves",
		"what",
		"which",
		"who",
		"whom",
		"this",
		"that",
		"these",
		"those",
		"am",
		"is",
		"are",
		"was",
		"were",
		"be",
		"been",
		"being",
		"have",
		"has",
		"had",
		"having",
		"do",
		"does",
		"did",
		"doing",
		"a",
		"an",
		"the",
		"and",
		"but",
		"if",
		"or",
		"because",
		"as",
		"until",
		"while",
		"of",
		"at",
		"by",
		"for",
		"with",
		"about",
		"against",
		"between",
		"into",
		"through",
		"during",
		"before",
		"after",
		"above",
		"below",
		"to",
		"from",
		"up",
		"down",
		"in",
		"out",
		"on",
		"off",
		"over",
		"under",
		"again",
		"further",
		"then",
		"once",
		"here",
		"there",
		"when",
		"where",
		"why",
		"how",
		"all",
		"any",
		"both",
		"each",
		"few",
		"more",
		"most",
		"other",
		"some",
		"such",
		"no",
		"nor",
		"not",
		"only",
		"own",
		"same",
		"so",
		"than",
		"too",
		"very",
		"s",
		"t",
		"can",
		"will",
		"just",
		"don",
		"should",
		"now",
	}
}