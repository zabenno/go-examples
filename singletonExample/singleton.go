package main

import "fmt"

var myint *int

func GetSingleInt() *int {
	if myint == nil {
		x := 1
		myint = &x
	}
	return myint
}

func main() {
	first := GetSingleInt()
	fmt.Println(*first)

	second := GetSingleInt()
	fmt.Println(*second)

	fmt.Println(first)
	fmt.Println(second)

	*first = 2
	fmt.Println(*first)
	fmt.Println(*second)
}
