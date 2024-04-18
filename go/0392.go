package main

import (
	"strings"
)

// https://leetcode.com/problems/is-subsequence/
func isSubsequenceV1(s string, t string) bool {
	lastIndex := -1

	for _, char := range s {
		r := lastIndex + 1
		if r >= len(t) {
			return false
		}

		newIndex := strings.IndexRune(t[r:], char)
		if newIndex == -1 {
			return false
		}

		lastIndex += 1 + newIndex
	}

	return true
}

func isSubsequenceV2(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	i := 0

	for j := 0; j < len(t); j++ {
		if t[j] == s[i] {
			i++
		}

		if i == len(s) {
			return true
		}
	}

	return false
}
