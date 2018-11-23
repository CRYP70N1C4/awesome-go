package main

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	return letterCombinations2(digits[1:], getLetter(string(digits[0])))
}

func getLetter(a string) []string {
	if a == "2" {
		return []string{"a", "b", "c"}
	} else if a == "3" {
		return []string{"d", "e", "f"}
	} else if a == "4" {
		return []string{"g", "h", "i"}
	} else if a == "5" {
		return []string{"j", "k", "l"}
	} else if a == "6" {
		return []string{"m", "n", "o"}
	} else if a == "7" {
		return []string{"p", "q", "r", "s"}
	} else if a == "8" {
		return []string{"t", "u", "v"}
	} else {
		return []string{"w", "x", "y", "z"}
	}
}

func letterCombinations2(digits string, prefix []string) []string {
	if len(digits) == 0 {
		return prefix
	} else {
		letter := string(digits[0])
		var result []string
		for _, a := range getLetter(letter) {
			for _, p := range prefix {
				result = append(result, p+a)
			}
		}
		digits = digits[1:]
		return letterCombinations2(digits, result)
	}
}

func main() {
	fmt.Println(letterCombinations("23"))
}
