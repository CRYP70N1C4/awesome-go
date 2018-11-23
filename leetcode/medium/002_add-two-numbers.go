package main

import "fmt"

func add(nums []int, nums2 []int) []int {
	size := len(nums)

	if len(nums2) > size {
		size = len(nums2)
		nums, nums2 = nums2, nums
	}
	carry := 0
	fmt.Println(nums, nums2)
	var result []int
	for i := 0; i < size; i++ {
		var tmp int
		if i < len(nums2) {
			tmp = nums[i] + carry + nums2[i]
		} else {
			tmp = nums[i] + carry
		}
		if tmp >= 10 {
			carry = tmp / 10
			tmp = tmp % 10
		} else {
			carry = 0
		}
		fmt.Println(tmp)
		result = append(result, tmp)
	}
	for ; carry > 0; carry = carry / 10 {
		result = append(result, carry)
	}

	return result
}

func main() {
	fmt.Println(add([]int{2, 4, 3}, []int{5, 66, 4, 1}))
}
