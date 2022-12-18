package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

func Top10(str string) []string {
	var pattern = regexp.MustCompile(`(\p{L})(\p{P}$)`)
	wordMap := make(map[string]int)
	var wordSlice []string

	for _, v := range strings.Fields(str) {
		word := pattern.ReplaceAllString(strings.ToLower(v), "$1")
		if word == "-" {
			continue
		}
		if value, ok := wordMap[word]; ok {
			wordMap[word] = value + 1
		} else {
			wordMap[word] = 1
			wordSlice = append(wordSlice, word)
		}
	}

	sort.Slice(wordSlice, func(i, j int) bool {
		if wordMap[wordSlice[i]] == wordMap[wordSlice[j]] {
			return wordSlice[i] < wordSlice[j]
		}
		return wordMap[wordSlice[i]] > wordMap[wordSlice[j]]
	})

	if len(wordSlice) > 10 {
		return wordSlice[:10]
	}
	return wordSlice
}
