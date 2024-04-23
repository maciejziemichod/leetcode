package main

// https://leetcode.com/problems/can-place-flowers/
func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	if len(flowerbed) == 1 {
		return flowerbed[0] == 0 && n == 1
	}

	c := 0

	if flowerbed[0] == 0 && flowerbed[1] == 0 {
		c++
		flowerbed[0] = 1
	}

	for i := 1; i < len(flowerbed)-1; i++ {
		if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			c++
		}
	}

	if flowerbed[len(flowerbed)-1] == 0 && flowerbed[len(flowerbed)-2] == 0 {
		c++
	}

	return c >= n
}
