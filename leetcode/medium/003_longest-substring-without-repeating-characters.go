package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	size, i, j := 0, 0, 0
	dict := make(map[uint8]int)
	for ; j < len(s); j++ {
		key := s[j]
		if _, ok := dict[key]; ok {
			i = max(dict[key]+1, i) // cannot jump back
		}
		dict[key] = j
		size = max(size, j-i+1)
	}
	return size
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabdcbb"))
}
