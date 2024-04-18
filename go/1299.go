package main

// https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/
func replaceElementsV1(arr []int) []int {
	r := make([]int, len(arr))
	r[len(arr)-1] = -1

	for i := 0; i < len(arr)-1; i++ {
		biggest := arr[i+1]

		for j := i + 2; j < len(arr); j++ {
			if arr[j] > biggest {
				biggest = arr[j]
			}
		}

		r[i] = biggest
	}

	return r
}

func replaceElementsV2(arr []int) []int {
	r := make([]int, len(arr))
	h := 0

	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] > h {
			h = arr[i]
		}

		r[i-1] = h
	}

	r[0] = h
	r[len(arr)-1] = -1

	return r
}

func replaceElementsV3(arr []int) []int {
	max := -1
	for i := len(arr) - 1; i >= 0; i-- {
		newMax := max
		if arr[i] > max {
			newMax = arr[i]
		}

		arr[i] = max
		max = newMax
	}

	return arr
}
