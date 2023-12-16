package main

import (
	"basicsofgo/io/data"
	"fmt"
)

var url = 'i'

func calcTax(price float32) (float32, float32){
	return price+1,price*2.0
}

func calcTaxNamedOutputs(price float32) (state float32, city float32){
	state = price+1
	city = price*2.0
	return
}

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
	fmt.Println('a')

	var start float32 = 100
	state, city := calcTax(start)
	fmt.Println(state,city)

	age := 22
	addAge(&age)
	addAge(&age)
	fmt.Println(age)
}

func addAge(age *int) {
	if *age < 100 {
		panic("Age is less than 100")
	}
	*age++
}