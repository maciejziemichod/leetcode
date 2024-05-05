package main

import "slices"

// https://leetcode.com/problems/next-greater-element-i/
func nextGreaterElementV1(nums1 []int, nums2 []int) []int {
	n := make([]int, len(nums1))

	for i, num1 := range nums1 {
		n[i] = -1
		found := false

		for _, num2 := range nums2 {
			if found && num2 > num1 {
				n[i] = num2
				break
			}

			if !found && num2 == num1 {
				found = true
			}
		}
	}

	return n
}

func nextGreaterElementV2(nums1 []int, nums2 []int) []int {
	n, s, o := make([]int, len(nums1)), []int{}, []int{}

	for _, num := range nums2 {
		c := 0
		for _, sv := range s {
			if sv < num {
				c++
			}
		}

		for i := len(s) - c; i < len(s); i++ {
			n[o[i]] = num
		}
		s = s[:len(s)-c]
		o = o[:len(o)-c]

		index := slices.Index(nums1, num)
		if index != -1 {
			s = append(s, nums1[index])
			o = append(o, index)
			n[index] = -1
		}
	}

	return n
}

func nextGreaterElementV3(nums1 []int, nums2 []int) []int {
	m, s := make(map[int]int, len(nums2)), []int{}
	for _, v := range nums2 {
		for len(s) > 0 && v > s[len(s)-1] {
			m[s[len(s)-1]] = v
			s = s[:len(s)-1]
		}

		s = append(s, v)
	}

	r := make([]int, len(nums1))
	for i, v := range nums1 {
		if g, ok := m[v]; ok {
			r[i] = g
		} else {
			r[i] = -1
		}
	}

	return r
}
