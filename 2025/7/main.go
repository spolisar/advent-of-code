package main

import (
	"os"
	"fmt"
)

func CountSplits(inputStr string) int {
	count := 0
	return count
}

func main() {
	inputFilename := "example.txt"
	// inputFilename := "input.txt"
	b, err := os.ReadFile(inputFilename)
	if err != nil {
		fmt.Println(err)
	}
	inputStr := string(b)
	result := CountSplits(inputStr)
	fmt.Printf("result: %d\n", result)
}
