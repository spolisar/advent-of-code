package main

import (
	"fmt"
	"os"
)

func calcJunctions(inputStr string) int {
	return 0
}

func main() {
	inputFilename := "example.txt"
	// inputFilename := "input.txt"
	b, err := os.ReadFile(inputFilename)
	if err != nil {
		fmt.Println(err)
	}
	inputStr := string(b)
	result := calcJunctions(inputStr)
	fmt.Printf("result: %d\n", result)
}
