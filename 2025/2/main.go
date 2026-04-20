package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// assuming IDs have at least 2 digits

func isInvalidID(x int) bool {
	s := strconv.Itoa(x)
	if len(s)%2 != 0 {
		return false
	}
	midInd := len(s) / 2
	return s[:midInd] == s[midInd:]
}

func isInvalidID2(x int) bool {
	s := strconv.Itoa(x)
	doubled := s + s
	iOf := strings.Index(doubled[1:], s)
	return iOf+1 != len(s)
}

func splitAt(substring []byte) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Source - https://stackoverflow.com/a/78962169
	// Posted by Rase
	// Retrieved 2026-04-19, License - CC BY-SA 4.0

	// creates a splitting function for a scanner
	// intended to be used for scanner.split() to split at a delimiter
	// example for splitting at commas:
	// scanner.Split(splitAt([]byte(",")))

	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {

		// Return nothing if at the end of the file and no data passed
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		// Find the index of the input of the separator substring
		if i := bytes.Index(data, substring); i >= 0 {
			return i + len(substring), data[0:i], nil
		}

		// If at the end of the file with data, return the data
		if atEOF {
			return len(data), data, nil
		}

		return
	}
}

func SumInvalidIDs(input string) int {
	sum := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(splitAt([]byte(",")))
	for scanner.Scan() {
		rangeStr := scanner.Text()
		lowStr, highStr := func() (string, string) {
			rangeVals := strings.Split(rangeStr, "-")
			return rangeVals[0], rangeVals[1]
		}()
		// last highStr may have a trailing \n
		highStr = strings.TrimSuffix(highStr, "\n")

		low, err := strconv.Atoi(lowStr)
		if err != nil {
			fmt.Printf("Error converting str to int: %s\n", lowStr)
			fmt.Println(err.Error())
		}
		high, err := strconv.Atoi(highStr)
		if err != nil {
			fmt.Printf("Error converting str to int: %s\n", highStr)
			fmt.Println(err.Error())
		}

		// simplest approach would just be to convert numbers in the range to a string and compare slices
		for x := low; x <= high; x++ {
			if isInvalidID(x) {
				sum += x
			}
		}
	}
	return sum
}

func SumInvalidIDs2(input string) int {
	sum := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(splitAt([]byte(",")))
	for scanner.Scan() {
		rangeStr := scanner.Text()
		lowStr, highStr := func() (string, string) {
			rangeVals := strings.Split(rangeStr, "-")
			return rangeVals[0], rangeVals[1]
		}()
		// last highStr may have a trailing \n
		highStr = strings.TrimSuffix(highStr, "\n")

		low, err := strconv.Atoi(lowStr)
		if err != nil {
			fmt.Printf("Error converting str to int: %s\n", lowStr)
			fmt.Println(err.Error())
		}
		high, err := strconv.Atoi(highStr)
		if err != nil {
			fmt.Printf("Error converting str to int: %s\n", highStr)
			fmt.Println(err.Error())
		}

		for x := low; x <= high; x++ {
			if isInvalidID2(x) {
				sum += x
			}
		}
	}
	return sum
}

func main() {
	// inputFilename := "example.txt"
	inputFilename := "input.txt"
	b, err := os.ReadFile(inputFilename)
	if err != nil {
		fmt.Println(err)
	}
	inputStr := string(b)
	invalidSum := SumInvalidIDs2(inputStr)
	fmt.Printf("sum of invalid IDs: %d\n", invalidSum)
}
