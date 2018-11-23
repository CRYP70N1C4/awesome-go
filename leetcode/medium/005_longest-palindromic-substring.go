package main

import "fmt"

func expand(str string, left, right int) (int, string) {
	for left >= 0 && right < len(str) && str[left] == str[right] {
		left--
		right++
	}
	if left == right {
		return 1, string(str[left])
	}
	return right - left - 1, str[left+1 : right]
}

func longestPalindrome(s string) string {
	max := 0
	var str string
	for i := 0; i < len(s); i++ {
		if length, tmp := expand(s, i, i); max < length {
			str, max = tmp, length
		}
		if i < len(s)-1 {
			if length, tmp := expand(s, i, i+1); max < length {
				str, max = tmp, length
			}
		}
	}
	return str
}

func main() {
	fmt.Println(longestPalindrome("bb"))
}
