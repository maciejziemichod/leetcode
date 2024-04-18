package main

import (
	"sort"
)

// https://leetcode.com/problems/valid-anagram/
func isAnagramV1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	ss := []rune(s)
	ts := []rune(t)

	sort.Slice(ss, func(i int, j int) bool { return ss[i] < ss[j] })
	sort.Slice(ts, func(i int, j int) bool { return ts[i] < ts[j] })

	return string(ss) == string(ts)
}

func isAnagramV2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sm := make(map[rune]int)

	for _, char := range s {
		count, ok := sm[char]
		if !ok {
			sm[char] = 1
		} else {
			sm[char] = count + 1
		}
	}

	for _, char := range t {
		count, ok := sm[char]
		if !ok || count <= 0 {
			return false
		} else {
			sm[char] = count - 1
		}
	}

	return true
}
