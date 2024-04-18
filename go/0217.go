package main

// https://leetcode.com/problems/contains-duplicate/
func containsDuplicate(nums []int) bool {
	numsMap := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := numsMap[num]; ok {
			return true
		}

		numsMap[num] = struct{}{}
	}

	return false
}
