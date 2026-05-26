package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Set map[int]struct{}

func (s *Set) Add(x int) {
	(*s)[x] = struct{}{}
}

func (s *Set) Remove(x int) {
	delete(*s, x)
}

func (s *Set) Has(x int) bool {
	_, found := (*s)[x]
	return found
}

func readInput(inputStr string) []string {
	lineScanner := bufio.NewScanner(strings.NewReader(inputStr))
	var results []string
	for lineScanner.Scan() {
		results = append(results, lineScanner.Text())
	}
	return results
}

func CountSplits(inputStr string) int {
	count := 0
	lines := readInput(inputStr)
	if len(lines) == 0 {
		return count
	}

	beams := make(Set)

	for i, c := range lines[0] {
		if c == 'S' {
			beams.Add(i)
		}
	}

	for r := 1; r < len(lines); r++ {
		for col, char := range lines[r] {
			if char == '^' && beams.Has(col) {
				beams.Remove(col)
				beams.Add(min(col+1, len(lines[r])-1))
				beams.Add(max(col-1, 0))
				count++
			}
		}
	}

	return count
}

func CountSplits2(inputStr string) int {
	count := 0
	lines := readInput(inputStr)
	if len(lines) == 0 {
		return count
	}

	beams := make(map[int]int)

	for i, c := range lines[0] {
		if c == 'S' {
			beams[i] = 1
		}
	}

	for r := 1; r < len(lines); r++ {
		for col, char := range lines[r] {
			paths, has := beams[col]
			if char == '^' && has {
				delete(beams, col)
				left := max(col-1, 0)
				right := min(col+1, len(lines[r])-1)
				beams[left] += paths
				beams[right] += paths
			}
		}
	}
	for _, v := range beams {
		count += v
	}

	return count
}

func main() {
	inputFilename := "example.txt"
	// inputFilename := "input.txt"
	b, err := os.ReadFile(inputFilename)
	if err != nil {
		fmt.Println(err)
	}
	inputStr := string(b)
	result := CountSplits2(inputStr)
	fmt.Printf("result: %d\n", result)
}
