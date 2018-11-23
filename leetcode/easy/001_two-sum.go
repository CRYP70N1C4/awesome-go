package main

import "fmt"

//easy

func twoSum(nums []int, target int) (int, int) {
	dict := make(map[int]int)
	for id, num := range nums {
		dict[target-num] = id
	}
	for id, num := range nums {
		if id2, ok := dict[num]; ok && id != id2 {
			return id, id2
		}
	}

	return -1, -1
}

func main() {
	nums := []int{2, 3, 7, 22, 12, 9, 12}
	fmt.Println(twoSum(nums, 10))
}
