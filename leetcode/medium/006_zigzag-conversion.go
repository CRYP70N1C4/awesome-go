package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	var data [][]string
	for i := 0; i < numRows; i++ {
		data = append(data, make([]string, 0))
	}
	h := 0
	down := true
	for _, ch := range s {
		data[h] = append(data[h], string(ch))
		if down {
			if h < numRows-1 {
				h++
			} else {
				h--
				down = !down
			}

		} else {
			if h > 0 {
				h--
			} else {
				h++
				down = !down
			}
		}
	}
	res := ""

	for _, arr := range data {
		res += strings.Join(arr, "")
		//for _, ch := range arr {
		//	res += string(ch)
		//}
	}

	return res
}

func main() {
	fmt.Println(convert("abcdef", 3))
}
