package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	length := len(nums)
	for i, a := range nums {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		l, r := i+1, length-1
		for l < r {
			val := nums[l] + nums[r] + a
			if val == 0 {
				result = append(result, []int{nums[l], nums[r], a})
				l++
				r--
				for nums[l] == nums[l-1] && l <= r {
					l++
				}
				for nums[r] == nums[r+1] && l <= r {
					r--

				}
			} else if val > 0 {
				r--
			} else {
				l++
			}
		}

	}
	return result
}

func main() {
	fmt.Println(threeSum([]int{0, 0, 0}))
}
