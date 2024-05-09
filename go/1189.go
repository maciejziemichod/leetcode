package main

// https://leetcode.com/problems/maximum-number-of-balloons/
func maxNumberOfBalloons(text string) int {
	c := map[rune]int{'b': 0, 'a': 0, 'l': 0, 'o': 0, 'n': 0}

	for _, char := range text {
		c[char]++
	}

	b, a, l, o, n := c['b'], c['a'], c['l'], c['o'], c['n']

	return min(b, a, l/2, o/2, n)
}
