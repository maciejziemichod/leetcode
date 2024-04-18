package main

// https://leetcode.com/problems/remove-element/
func removeElementV1(nums []int, val int) int {
	r := 0
	for _, num := range nums {
		if num != val {
			nums[r] = num
			r++
		}
	}

	return r
}

func removeElementV2(nums []int, val int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		if nums[left] == val {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		} else {
			left++
		}
	}

	return left
}
