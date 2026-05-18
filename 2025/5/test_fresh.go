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
	result := CountFresh(inputStr)
	expected := 3
	if result != expected {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expected)
	}
}

