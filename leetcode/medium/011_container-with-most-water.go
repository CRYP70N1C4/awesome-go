package main

import "fmt"

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	area, l, r := 0, 0, len(height)-1
	for l < r {
		area = Max(area, (r-l)*Min(height[r], height[l]))
		if height[r] < height[l] {
			r--
		} else {
			l++
		}
	}
	return area
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
