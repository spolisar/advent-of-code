package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RotateCount(input string) int {
	// initial position is 50
	position := 50
	count := 0
	numSlots := 100
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		text := scanner.Text()
		direction := text[0]
		distance, err := strconv.Atoi(text[1:])
		if err != nil {
			return -1
		}
		// fmt.Printf("%c %d\n", direction, distance)
		switch direction {
		case 'L':
			position = (position - distance + numSlots) % numSlots
			// if position < 0 {
			// 	position = numSlots + position
			// }
		case 'R':
			position = (position + distance) % numSlots
		}
		if position == 0 {
			count++
		}
		// fmt.Printf("position: %d\n", position)
	}
	return count
}

func main() {
	// exampleFilename := "example.txt"
	inputFilename := "input.txt"
	b, err := os.ReadFile(inputFilename)
	if err != nil {
		fmt.Println(err)
	}
	inputStr := string(b)
	// fmt.Println(inputStr)
	zeroCount := RotateCount(inputStr)
	fmt.Printf("zero count: %d", zeroCount)
}
