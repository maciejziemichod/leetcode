package main

// https://leetcode.com/problems/length-of-last-word/
func lengthOfLastWordV1(s string) int {
	p, l := 0, 0

	for _, char := range s {
		if char == ' ' {
			if l != 0 {
				p = l
			}

			l = 0
		} else {
			l++
		}
	}

	if l != 0 {
		return l
	}

	return p
}

func lengthOfLastWordV2(s string) int {
	l := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			l++
		} else if l > 0 {
			break
		}
	}

	return l
}
