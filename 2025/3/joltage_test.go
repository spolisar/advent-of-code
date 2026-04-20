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
	result := SumJoltage(inputStr)
	expected := 357
	if result != expected {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expected)
	}
}

func TestPartTwo(t *testing.T) {
	inputFilename := "example.txt"

	b, err := os.ReadFile(inputFilename)
	if err != nil {
		fmt.Println(err)
	}

	inputStr := string(b)
	result := SumJoltage2(inputStr)
	expected := 3121910778619
	if result != expected {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expected)
	}
}
