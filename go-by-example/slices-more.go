package main

import (
	"fmt"
	"iotuil"
	"regexp"
)

func main() {
	dig := FindDigits("./slices-more.go")
	fmt.Println("digits:", dig)
}


func FindDigits(filename string) []byte {
	fmt.Println("filename:", filename)
	var digitRegexp = regexp.MustCompile("[0-9]+")
    b, _ := iotuil.ReadFile(filename)
	fmt.Println("b:", b)
    return digitRegexp.Find(b)
}
