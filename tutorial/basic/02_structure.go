package main

import "fmt"

func main() {
	map3()
}

func map3() {
	var dict map[string]int
	fmt.Println(dict)
	dict2 := make(map[string]int)
	dict2["one"] = 1
	dict2["two"] = 2
	fmt.Println(dict2)
	delete(dict2, "one")
	delete(dict2, "one")
	fmt.Println(dict2)

}

func slice2() {
	slice := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	a := slice[2:4]
	fmt.Println(a)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)
	slice = append(slice, 11, 22)
	fmt.Println(slice)

}

func array1() {
	var arr [10]int
	arr[0] = 11
	arr[3] = 111
	fmt.Println(arr)

	a := [3]int{1, 2, 3}
	b := [10]int{1, 2, 3}
	c := [...]int{4, 5, 6}

	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{1, 2, 3, 4}}
	easyArray := [2][4]int{{1, 2, 3, 4}, {4, 4}}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(doubleArray)
	fmt.Println(easyArray)

}
