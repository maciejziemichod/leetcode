package main

// https://leetcode.com/problems/pascals-triangle-ii/
func getRowV1(rowIndex int) []int {
	t := [][]int{{1}, {1, 1}}

	if rowIndex < 2 {
		return t[rowIndex]
	}

	for i := 0; i < rowIndex-1; i++ {
		r := []int{1}
		last := t[len(t)-1]

		for j := 0; j < len(last)-1; j++ {
			r = append(r, last[j]+last[j+1])
		}

		r = append(r, 1)
		t = append(t, r)
	}

	return t[rowIndex]
}

func getRowV2(rowIndex int) []int {
	if rowIndex <= 0 {
		return []int{1}
	}

	if rowIndex == 1 {
		return []int{1, 1}
	}

	last := []int{1, 1}

	for i := 0; i < rowIndex-1; i++ {
		r := []int{1}
		for j := 0; j < len(last)-1; j++ {
			r = append(r, last[j]+last[j+1])
		}

		last = append(r, 1)
	}

	return last
}
