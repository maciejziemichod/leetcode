package main

// https://leetcode.com/problems/majority-element/
func majorityElementV1(nums []int) int {
	c := make(map[int]int)

	for _, num := range nums {
		_, ok := c[num]
		if !ok {
			c[num] = 0
		}
		c[num]++
	}

	m, mv := nums[0], 0
	for k, v := range c {
		if v > mv {
			m = k
			mv = v
		}
	}

	return m
}

func majorityElementV2(nums []int) int {
	max, score := nums[0], 0

	for _, num := range nums {
		if num == max {
			score++
		} else {
			score--
		}

		if score < 0 {
			max = num
			score = 0
		}
	}

	return max
}

// same as v3 but does the algorithm correctly (not needed for the challenge as there is always a majority element)
func majorityElementV3(nums []int) int {
	max, score := nums[0], 1

	for _, num := range nums {
		if num == max {
			score++
		} else {
			score--
		}

		if score == 0 {
			max = num
			score = 1
		}
	}

	count := 0

	for _, num := range nums {
		if num == max {
			count++
		}
	}

	if count <= len(nums)/2 {
		return -1
	}

	return max
}
