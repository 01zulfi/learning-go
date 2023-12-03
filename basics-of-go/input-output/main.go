package main

import (
	"basicsofgo/io/data"
	"fmt"
)

var url = 'i'

func main() {
	const pi float64 = 3.14
	var message string
	var price = 45.1
	message = "first"

	// PrintData()
	fmt.Println(data.Countries)
	fmt.Println(data.CSlice)
	fmt.Println(data.CMap)
	print(message, price, pi)
}