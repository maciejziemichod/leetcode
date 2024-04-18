package main

// https://leetcode.com/problems/pascals-triangle/
func generate(numRows int) [][]int {
	t := [][]int{{1}, {1, 1}}

	if numRows < 3 {
		return t[:numRows]
	}

	for i := 0; i < numRows-2; i++ {
		r := []int{1}
		last := t[len(t)-1]

		for j := 0; j < len(last)-1; j++ {
			r = append(r, last[j]+last[j+1])
		}

		r = append(r, 1)
		t = append(t, r)
	}

	return t
}
