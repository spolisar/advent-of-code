package main

import (
	"fmt"
	"os"
)

func RotateCount(input string) int {
	// initial position is 50
	position := 50
	return 0
}

func main() {
	exampleFilename := "example.txt"
	b, err := os.ReadFile(exampleFilename)
	if err != nil {
		fmt.Println(err)
	}
	inputStr := string(b)
	fmt.Println(inputStr)
}
