package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func ApostropheCheck(s string) string {
	if len(s) < 2 {
		return s
	}
	single := rune(39)
	rs := []rune(s)
	end := len(rs) - 1
	// 'word'
	if rs[0] == single && rs[end] == single {
		s = s[1:]
		end--
		s = s[:end]
		end--
	}
	rs = []rune(s)
	switch {
	case rs[0] == single && rs[end] != single:
		// 'word
		s = s[1:]
		end--
	case len(rs) > 0 && rs[0] != single && rs[end] == single:
		// word'
		s = s[:end]
		end--
	}
	return s
}

func WordCount(phrase string) Frequency {
	freq := Frequency{}
	filter := regexp.MustCompile(`[^a-zA-Z\d']`)              // exclude alphanumeric & single-quote
	wordList := regexp.MustCompile(`[\s,]`).Split(phrase, -1) // split at whitespace and/or comma
	for _, word := range wordList {
		key := ApostropheCheck(strings.ToLower(filter.ReplaceAllString(word, "")))
		if len(key) > 0 {
			freq[key]++
		}
	}
	return freq
}
