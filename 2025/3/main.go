package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SumJoltage(input string) int {
	sum := 0
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		leftMaxInd := 0
		line := scanner.Text()
		for i := 0; i < len(line)-1; i++ {
			if line[i] > line[leftMaxInd] {
				leftMaxInd = i
			}
		}
		rightMaxInd := leftMaxInd + 1
		for i := rightMaxInd; i < len(line); i++ {
			if line[i] > line[rightMaxInd] {
				rightMaxInd = i
			}
		}
		s := fmt.Sprintf("%c%c", line[leftMaxInd], line[rightMaxInd])
		joltage, err := strconv.Atoi(string(s))
		if err != nil {
			fmt.Println(line)
			fmt.Println(s)
			fmt.Println(err.Error())
			return -1
		}
		sum += joltage
	}

	return sum
}

func SumJoltage2(input string) int {
	sum := 0
	const numDigits = 12
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()
		maxInds := [numDigits]int{}

		for digitInd := range numDigits {
			initialInd := digitInd
			if digitInd > 0 {
				initialInd = maxInds[digitInd-1] + 1
				maxInds[digitInd] = initialInd
			}
			end := len(line) - numDigits + digitInd + 1
			for i := initialInd; i < end; i++ {
				if line[i] > line[maxInds[digitInd]] {
					maxInds[digitInd] = i
				}
			}
		}
		joltage := 0
		for i := range numDigits {
			digit := int(line[maxInds[i]] - '0')
			joltage = (joltage * 10) + digit
		}
		sum += joltage
	}

	return sum
}

func main() {
	inputFilename := "example.txt"
	// inputFilename := "input.txt"
	b, err := os.ReadFile(inputFilename)
	if err != nil {
		fmt.Println(err)
	}
	inputStr := string(b)
	joltageSum := SumJoltage2(inputStr)
	fmt.Printf("sum of joltages: %d\n", joltageSum)
}
