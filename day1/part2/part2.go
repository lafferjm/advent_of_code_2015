package part2

import (
	"fmt"
	"os"
)

func Part2() {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("error reading file")
		os.Exit(1)
	}

	fileContents := string(file)
	floor := 0
	floor_counter := 0

	for i := 0; i < len(fileContents); i++ {
		if fileContents[i] == '(' {
			floor += 1
		} else if fileContents[i] == ')' {
			floor -= 1
		}

		floor_counter += 1

		if floor < 0 {
			break
		}
	}

	fmt.Println(floor_counter)
}
