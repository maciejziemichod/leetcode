package main

// https://leetcode.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	l := 0

	for i, char := range strs[0] {
		for _, str := range strs {
			if i >= len(str) || rune(str[i]) != char {
				return str[:l]
			}
		}

		l++
	}

	return strs[0]
}
