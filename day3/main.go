package main

import (
	"fmt"
	"os"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) move(direction string) Coordinate {
	newCoordinate := c

	if direction == ">" {
		newCoordinate.X += 1
	} else if direction == "<" {
		newCoordinate.X -= 1
	} else if direction == "^" {
		newCoordinate.Y += 1
	} else if direction == "v" {
		newCoordinate.Y -= 1
	}

	return *newCoordinate
}

func getVisitedHouses(data string) map[Coordinate]bool {
	position := Coordinate{0, 0}
	coordSet := make(map[Coordinate]bool)

	coordSet[position] = true
	for _, i := range data {
		coord := position.move(string(i))
		coordSet[coord] = true
	}

	return coordSet
}

func main() {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := string(file)

	partOneSanta := getVisitedHouses(data)

	var santasDirections []string
	var robosDirections []string
	for i := 0; i < len(data)-1; i += 2 {
		santasDirections = append(santasDirections, string(data[i]))
		robosDirections = append(robosDirections, string(data[i+1]))
	}

	partTwoSanta := getVisitedHouses(strings.Join(santasDirections[:], ""))
	partTwoRobo := getVisitedHouses(strings.Join(robosDirections[:], ""))

	for k := range partTwoRobo {
		partTwoSanta[k] = true
	}

	fmt.Printf("Part One: %d\n", len(partOneSanta))
	fmt.Printf("Part Two: %d\n", len(partTwoSanta))
}
