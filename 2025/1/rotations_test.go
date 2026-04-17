package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	// exampleFilename := "example.txt"

	result := RotateCount("")
	expected := 3
	if result != expected {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 3)
	}
}
