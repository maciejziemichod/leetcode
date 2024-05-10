package main

import "strings"

// https://leetcode.com/problems/word-pattern/
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	patternMap, usedWords := map[byte]string{}, map[string]struct{}{}

	for i := 0; i < len(pattern); i++ {
		char, word := pattern[i], words[i]

		matchedWord, exists := patternMap[char]
		_, used := usedWords[word]

		if !exists && used {
			return false
		} else if !exists {
			patternMap[char] = word
			usedWords[word] = struct{}{}
		} else if matchedWord != word {
			return false
		}
	}

	return true
}
