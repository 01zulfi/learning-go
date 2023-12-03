package main

import "fmt"

var global = "global"

func PrintData() {
	var n = 10.4
	nAsInt := int(n)
	fmt.Println("n", nAsInt)
	fmt.Print(global)
	fmt.Println(url)
}