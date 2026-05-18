package main

import (
	"fmt"
	"os"
	"strings"
)

type Grid [][]rune

func (g Grid) String() string {
	var sb strings.Builder
	for _, row := range g {
		for _, c := range row {
			sb.WriteRune(c)
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func convertGrid(inputStr string) Grid {
	lines := strings.Split(inputStr, "\n")

	var result Grid

	for row, line := range lines {
		if len(line) > 0 {
			result = append(result, []rune{})
		}
		for _, c := range line {
			result[row] = append(result[row], c)
		}
	}
	return result
}

func countAdjacents(g Grid, r, c int) (count int) {
	// '.' for empty, '@' for occupied
	for i := r - 1; i <= r+1; i++ {
		for j := c - 1; j <= c+1; j++ {
			if i < 0 || j < 0 || i >= len(g) || j >= len(g[0]) {
				continue
			} else if i == r && j == c {
				continue
			} else if g[i][j] == '@' {
				count++
			}
		}
	}
	// fmt.Printf("adjacents: %d\n", count)
	return count
}

func CountRemovable(inputStr string) (removableCount int) {
	grid := convertGrid(inputStr)

	for r := range grid {
		for c := range grid[r] {
			// '.' for empty, '@' for occupied
			if grid[r][c] == '.' {
				continue
			}
			adjacents := countAdjacents(grid, r, c)
			if adjacents < 4 {
				removableCount++
			}
		}
	}
	return removableCount
}

func CountRemovable2(inputStr string) (removableCount int) {
	grid := convertGrid(inputStr)

	for removed := true; removed; {
		removed = false
		for r := range grid {
			for c := range grid[r] {
				// '.' for empty, '@' for occupied
				if grid[r][c] == '.' {
					continue
				}
				adjacents := countAdjacents(grid, r, c)
				if adjacents < 4 {
					grid[r][c] = '.'
					removableCount++
					removed = true
				}
			}
		}
	}
	return removableCount
}

func main() {
	inputFilename := "example.txt"
	// inputFilename := "input.txt"
	b, err := os.ReadFile(inputFilename)
	if err != nil {
		fmt.Println(err)
	}
	inputStr := string(b)
	result := CountRemovable2(inputStr)
	fmt.Printf("result: %d\n", result)
}
