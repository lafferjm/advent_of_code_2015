package main

import (
	"bufio"
	"fmt"
	"lafferjm.com/day2/box"
	"os"
	"strconv"
	"strings"
)

func getLinesInFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	file.Close()

	return fileLines, nil
}

func main() {
	lines, err := getLinesInFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	var paperNeeded int
	var ribbonNeeded int
	for _, line := range lines {
		parts := strings.Split(line, "x")
		length, _ := strconv.Atoi(parts[0])
		width, _ := strconv.Atoi(parts[1])
		height, _ := strconv.Atoi(parts[2])

		b := box.Box{
			Length: length,
			Width:  width,
			Height: height,
		}

		paperNeeded += b.GetSmallestArea() + b.GetSquareFeet()
		ribbonNeeded += b.GetBowFeet() + b.GetRibbonFeet()
	}

	fmt.Printf("Paper Needed: %d\n", paperNeeded)
	fmt.Printf("Ribbon Needed: %d\n", ribbonNeeded)
}
