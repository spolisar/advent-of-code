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
	result := CountSplits(inputStr)
	expected := 21
	if result != expected {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expected)
	}
}

func TestExample2(t *testing.T) {
	inputFilename := "example.txt"

	b, err := os.ReadFile(inputFilename)
	if err != nil {
		fmt.Println(err)
	}

	inputStr := string(b)
	result := CountSplits2(inputStr)
	expected := 40
	if result != expected {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expected)
	}
}
