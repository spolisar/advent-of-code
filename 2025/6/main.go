package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convInput(inputStr string) ([][]int, []string) {
	var elements [][]string
	lineScanner := bufio.NewScanner(strings.NewReader(inputStr))
	for lineScanner.Scan() {
		elements = append(elements, strings.Fields(lineScanner.Text()))
	}
	var nums [][]int
	for r, row := range elements[:len(elements)-1] {
		nums = append(nums, []int{})
		for _, numStr := range row {
			x, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println(err)
				return nil, nil
			}
			nums[r] = append(nums[r], x)
		}
	}
	return nums, elements[len(elements)-1]
}

func CalcTotal(inputStr string) int {
	total := 0

	nums, ops := convInput(inputStr)
	if nums == nil {
		fmt.Println("error converting input")
		return -1
	}
	// start with first row
	results := append(make([]int, 0, len(ops)), nums[0]...)
	for _, row := range nums[1:] {
		for c, x := range row {
			switch ops[c] {
			case "+":
				results[c] += x
			case "*":
				results[c] *= x
			case "-":
				results[c] -= x
			}
		}
	}
	for _, x := range results {
		total += x
	}
	return total
}

func convInput2(inputStr string) ([][]int, []string) {
	lineScanner := bufio.NewScanner(strings.NewReader(inputStr))
	var rows []string
	for lineScanner.Scan() {
		rows = append(rows, lineScanner.Text())
	}
	if len(rows) == 0 {
		fmt.Println("got no rows in convInput2")
		return nil, nil
	}
	var nums [][]int
	nums = append(nums, []int{})
	probIndex := 0
	for c := range rows[0] {
		numBuilder := strings.Builder{}
		spaceCount := 0
		for r := range len(rows) - 1 {
			if rows[r][c] == ' ' {
				spaceCount++
				continue
			}
			numBuilder.WriteByte(rows[r][c])
		}
		if spaceCount == len(rows)-1 {
			probIndex++
			nums = append(nums, []int{})
			continue
		}
		x, err := strconv.Atoi(numBuilder.String())
		if err != nil {
			fmt.Printf("error converting %s\n%s\n", numBuilder.String(), err)
		}
		nums[probIndex] = append(nums[probIndex], x)
	}
	ops := strings.Fields(rows[len(rows)-1])
	return nums, ops
}

func CalcTotal2(inputStr string) int {
	total := 0

	nums, ops := convInput2(inputStr)
	if nums == nil {
		fmt.Println("error converting input")
		return -1
	}
	var results []int
	for probIndex, op := range ops {
		initial := nums[probIndex][len(nums[probIndex])-1]
		results = append(results, initial)
		for i := len(nums[probIndex])-2; i >= 0; i-- {
			switch op {
			case "+":
				results[probIndex] += nums[probIndex][i]
			case "*":
				results[probIndex] *= nums[probIndex][i]
			case "-":
				results[probIndex] -= nums[probIndex][i]
			}
		}
	}
	for _, x := range results {
		total += x
	}
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
	result := CalcTotal2(inputStr)
	fmt.Printf("result: %d\n", result)
}
