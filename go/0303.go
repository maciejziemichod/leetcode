package main

// https://leetcode.com/problems/range-sum-query-immutable/
type NumArray struct {
	nums []int
}

func ConstructorV1(nums []int) NumArray {
	return NumArray{nums}
}

func (this *NumArray) SumRangeV1(left int, right int) int {
	s := 0

	for i := left; i <= right; i++ {
		s += this.nums[i]
	}

	return s
}

func ConstructorV2(nums []int) NumArray {
	c, s := make([]int, len(nums)), 0

	for i, v := range nums {
		s += v
		c[i] = s
	}

	return NumArray{c}
}

func (this *NumArray) SumRangeV2(left int, right int) int {
	if left == 0 {
		return this.nums[right]
	}

	return this.nums[right] - this.nums[left-1]
}
