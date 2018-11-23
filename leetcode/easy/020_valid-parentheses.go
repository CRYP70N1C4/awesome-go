package main

import "fmt"

func isMatch(l string, r string) bool {
	fmt.Println(l, r)
	return (l == "{" && r == "}") || (l == "[" && r == "]") || (l == "(" && r == ")")
}

func isValid(s string) bool {
	length := len(s)
	if length%2 != 0 {
		return false
	}
	var stack []string

	for _, ch := range s {
		if len(stack) > 0 && isMatch(stack[len(stack)-1], string(ch)) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, string(ch))
		}
	}
	return len(stack) == 0
}
