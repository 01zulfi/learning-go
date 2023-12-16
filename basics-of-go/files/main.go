package main

import (
	"files/fileutils"
	"fmt"
	"os"
)

func main() {
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/data/test.txt"
	s, err := fileutils.ReadTextFile(filePath)

	if err == nil {
		write := fmt.Sprintf("Original content: %s\n Double content: %s%s", s, s, s)
		fileutils.WriteToFile(filePath+".output.txt", write)
		fmt.Println(s)

	} else {
		fmt.Println(err)
	}
}
