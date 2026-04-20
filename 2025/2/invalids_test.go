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
	result := SumInvalidIDs(inputStr)
	expected := 1227775554
	if result != expected {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPartTwo(t *testing.T) {
	inputFilename := "example.txt"

	b, err := os.ReadFile(inputFilename)
	if err != nil {
		fmt.Println(err)
	}

	inputStr := string(b)
	result := SumInvalidIDs2(inputStr)
	expected := 4174379265
	if result != expected {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, expected)
	}
}
// func TestExampleCross(t *testing.T) {
// 	inputFilename := "example.txt"
//
// 	b, err := os.ReadFile(inputFilename)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
//
// 	inputStr := string(b)
// 	result := RotateCountCross(inputStr)
// 	expected := 6
// 	if result != expected {
// 		t.Errorf("Result was incorrect, got: %d, want: %d.", result, expected)
// 	}
// }
