package main

// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
func findDisappearedNumbersV1(nums []int) []int {
	n := make([]int, len(nums))

	for _, num := range nums {
		n[num-1] = 1
	}

	d := []int{}
	for i, v := range n {
		if v == 0 {
			d = append(d, i+1)
		}
	}

	return d
}

func findDisappearedNumbersV2(nums []int) []int {
	abs := func(num int) int {
		if num < 0 {
			return -num
		}
		return num
	}

	for _, num := range nums {
		i := abs(num) - 1
		nums[i] = -abs(nums[i])
	}

	d := []int{}
	for i, v := range nums {
		if v > 0 {
			d = append(d, i+1)
		}
	}

	return d
}
