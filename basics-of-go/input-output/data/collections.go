package data

import "fmt"

var Countries [10]string
var CSlice []string
var CMap map[string]string

func addSubtract(a int, b int) (int,int) {
	return a+b,a-b
}

func init() {
	add, sub := addSubtract(2,2)
	fmt.Println(add, sub)

	fmt.Println(addSubtract(2,2))
}