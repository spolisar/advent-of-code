package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func convertRanges(rangeStr string) (mergedRanges []Range) {
	rangeStrs := strings.Split(rangeStr, "\n")
	initialRanges := []Range{}
	for _, r := range rangeStrs {
		splitRange := strings.Split(r, "-")
		start, err := strconv.Atoi(splitRange[0])
		if err != nil {
			fmt.Println(err)
			return mergedRanges
		}
		end, err := strconv.Atoi(splitRange[1])
		if err != nil {
			fmt.Println(err)
			return mergedRanges
		}
		initialRanges = append(initialRanges, Range{start: start, end: end})
	}

	rangeCmp := func(a, b Range) int {
		return cmp.Compare(a.start, b.start)
	}
	slices.SortFunc(initialRanges, rangeCmp)

	prev := initialRanges[0]
	for _, r := range initialRanges[1:] {
		if prev.end < r.start {
			mergedRanges = append(mergedRanges, prev)
			prev = r
			continue
		}
		prev.start = min(r.start, prev.start)
		prev.end = max(r.end, prev.end)
	}
	mergedRanges = append(mergedRanges, prev)
	return mergedRanges
}

func CountFresh(inputStr string) int {
	count := 0
	splitStrs := strings.Split(inputStr, "\n\n")
	ranges := convertRanges(splitStrs[0])
	idScanner := bufio.NewScanner(strings.NewReader(splitStrs[1]))
	for idScanner.Scan() {
		id, err := strconv.Atoi(idScanner.Text())
		if err != nil {
			fmt.Println(err)
			return -1
		}
		for _, r := range ranges {
			if r.start <= id && id <= r.end {
				count++
				break
			}
		}
	}
	return count
}
func CountFresh2(inputStr string) int {
	count := 0
	splitStrs := strings.Split(inputStr, "\n\n")
	ranges := convertRanges(splitStrs[0])
	for _, r := range ranges {
		numInRange := r.end - r.start + 1
		count += numInRange
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
	result := CountFresh2(inputStr)
	fmt.Printf("result: %d\n", result)
}
