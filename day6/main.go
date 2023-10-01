package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getFileLines(filePath string) ([]string, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()

	return lines, nil
}

func partOne(lines []string) int {
	grid := make([][]bool, 1000)
	for i := range grid {
		grid[i] = make([]bool, 1000)
	}

	for _, line := range lines {
		var startRow, startColumn, stopRow, stopColumn int
		if strings.HasPrefix(line, "turn on") {
			fmt.Sscanf(line, "turn on %d,%d through %d,%d", &startRow, &startColumn, &stopRow, &stopColumn)
			for i := startRow; i <= stopRow; i++ {
				for j := startColumn; j <= stopColumn; j++ {
					grid[i][j] = true
				}
			}
		}

		if strings.HasPrefix(line, "turn off") {
			fmt.Sscanf(line, "turn off %d,%d through %d,%d", &startRow, &startColumn, &stopRow, &stopColumn)
			for i := startRow; i <= stopRow; i++ {
				for j := startColumn; j <= stopColumn; j++ {
					grid[i][j] = false
				}
			}
		}

		if strings.HasPrefix(line, "toggle") {
			fmt.Sscanf(line, "toggle %d,%d through %d,%d", &startRow, &startColumn, &stopRow, &stopColumn)
			for i := startRow; i <= stopRow; i++ {
				for j := startColumn; j <= stopColumn; j++ {
					grid[i][j] = !grid[i][j]
				}
			}
		}
	}

	lightsOn := 0
	for _, row := range grid {
		for _, column := range row {
			if column {
				lightsOn += 1
			}
		}
	}

	return lightsOn
}

func partTwo(lines []string) int {
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	for _, line := range lines {
		var startRow, startColumn, stopRow, stopColumn int
		if strings.HasPrefix(line, "turn on") {
			fmt.Sscanf(line, "turn on %d,%d through %d,%d", &startRow, &startColumn, &stopRow, &stopColumn)
			for i := startRow; i <= stopRow; i++ {
				for j := startColumn; j <= stopColumn; j++ {
					grid[i][j] = grid[i][j] + 1
				}
			}
		}

		if strings.HasPrefix(line, "turn off") {
			fmt.Sscanf(line, "turn off %d,%d through %d,%d", &startRow, &startColumn, &stopRow, &stopColumn)
			for i := startRow; i <= stopRow; i++ {
				for j := startColumn; j <= stopColumn; j++ {
					if grid[i][j] > 0 {
						grid[i][j] = grid[i][j] - 1
					}
				}
			}
		}

		if strings.HasPrefix(line, "toggle") {
			fmt.Sscanf(line, "toggle %d,%d through %d,%d", &startRow, &startColumn, &stopRow, &stopColumn)
			for i := startRow; i <= stopRow; i++ {
				for j := startColumn; j <= stopColumn; j++ {
					grid[i][j] = grid[i][j] + 2
				}
			}
		}
	}

	totalBrightness := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			totalBrightness += grid[i][j]
		}
	}

	return totalBrightness
}
func main() {
	lines, err := getFileLines("./input.txt")
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	lightsOn := partOne(lines)
	totalBrightness := partTwo(lines)

	fmt.Println(lightsOn)
	fmt.Println(totalBrightness)
}
