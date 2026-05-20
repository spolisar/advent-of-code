package main

import (
	"fmt"
	"os"
	"testing"
)


func TestExample(t *testing.T) {
	inputFilename := "example.txt"

	b, err := os.ReadFile(inputFilename)
	if err != nil {
		fmt.Println(err)
	}

	inputStr := string(b)
	result := CalcTotal(inputStr)
	expected := 4277556
	if result != expected {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expected)
	}
}
