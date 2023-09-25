package part1

import (
	"fmt"
	"os"
)

func Part1() {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("failed to open file")
		os.Exit(1)
	}

	contents := string(file)
	floor := 0

	for i := 0; i < len(contents); i++ {
		if contents[i] == '(' {
			floor = floor + 1
		} else if contents[i] == ')' {
			floor = floor - 1
		}
	}

	fmt.Println(floor)
}
