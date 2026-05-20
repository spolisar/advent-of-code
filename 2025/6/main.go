package main

import (
	"fmt"
	"os"
)

func CalcTotal(inputStr string) int {
	total := 0
	return total
}

func main() {
	inputFilename := "example.txt"
	// inputFilename := "input.txt"
	b, err := os.ReadFile(inputFilename)
	if err != nil {
		fmt.Println(err)
	}
	inputStr := string(b)
	result := CalcTotal(inputStr)
	fmt.Printf("result: %d\n", result)
}
