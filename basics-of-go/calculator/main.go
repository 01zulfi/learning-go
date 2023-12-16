package main

import "fmt"

func main() {
	var operation string
	var n1 int
	var n2 int

	fmt.Println("Calculator	")
	fmt.Println("------------------")
	fmt.Println("Select operation:")
	fmt.Scanf("%s", &operation)
	fmt.Println("Enter first number:")
	fmt.Scanf("%d", &n1)
	fmt.Println("Enter second number:")
	fmt.Scanf("%d", &n2)

	switch operation {
	case "+":
		fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
	case "-":
		fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
	case "*":
		fmt.Printf("%d * %d = %d", n1, n2, n1*n2)
	case "/":
		fmt.Printf("%d / %d = %d", n1, n2, n1/n2)
	default:
		fmt.Println("Invalid operation")
	}
}
