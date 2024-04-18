package main

// https://leetcode.com/problems/contains-duplicate-ii/
func containsNearbyDuplicate(nums []int, k int) bool {
	numsMap := make(map[int]int)

	for i, num := range nums {
		if j, ok := numsMap[num]; ok {
			if i-j <= k {
				return true
			}
		}

		numsMap[num] = i
	}

	return false
}
