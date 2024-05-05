package main

// https://leetcode.com/problems/find-pivot-index/
func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, sum(nums)

	for i, v := range nums {
		if i > 0 {
			left += nums[i-1]
		}
		right -= v

		if left == right {
			return i
		}
	}

	return -1
}

func sum(nums []int) int {
	sum := 0

	for _, v := range nums {
		sum += v
	}

	return sum
}
