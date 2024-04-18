package main

// https://leetcode.com/problems/two-sum/
func twoSumV1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		num := nums[i]

		for j := i + 1; j < len(nums); j++ {
			if num+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}

func twoSumV2(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums))

	for currentIndex, num := range nums {
		if index, ok := numsMap[num]; ok {
			return []int{currentIndex, index}
		}

		numsMap[target-num] = currentIndex
	}

	return []int{-1, -1}
}
