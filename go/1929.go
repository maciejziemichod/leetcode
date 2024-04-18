package main

// https://leetcode.com/problems/concatenation-of-array/
func getConcatenation(nums []int) []int {
	l := len(nums)
	c := make([]int, l*2)

	for i, v := range nums {
		c[i] = v
		c[i+l] = v
	}

	return c
}

func getConcatenationV2(nums []int) []int {
	l := len(nums)
	c := make([]int, l*2)

	copy(c, nums)

	for i, v := range nums {
		c[i+l] = v
	}

	return c
}

func getConcatenationV3(nums []int) []int {
	return append(nums, nums...)
}
