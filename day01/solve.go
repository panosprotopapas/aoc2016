package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	Direction string
	Number    int
}

type Coordinate struct {
	x int
	y int
}

func main() {

	relative_directions := getInput()
	absolute_directions := constructAbsoluteDirections(relative_directions)

	// Part 1
	var x, y int
	for _, pair := range absolute_directions {
		if pair.Direction == "N" {
			y += pair.Number
		} else if pair.Direction == "S" {
			y -= pair.Number
		} else if pair.Direction == "E" {
			x += pair.Number
		} else if pair.Direction == "W" {
			x -= pair.Number
		}
	}
	absX := int(math.Abs(float64(x)))
	absY := int(math.Abs(float64(y)))
	fmt.Println("Part 1:", absX+absY)

	// Part 2
	var walkedCoordinates []Coordinate
	walkedCoordinates = append(walkedCoordinates, Coordinate{x: 0, y: 0})
	for _, direction := range absolute_directions {
		walkedCoordinates = addCoordinates(walkedCoordinates, direction, walkedCoordinates[len(walkedCoordinates)-1])
		checkForDuplicates(walkedCoordinates)
	}
}

func checkForDuplicates(coordinates []Coordinate) {
	// Checks if there are any duplicate coordinates in the given slice.
	for i := 0; i < len(coordinates)-1; i++ {
		for j := i + 1; j < len(coordinates); j++ {
			if coordinates[i] == coordinates[j] {
				absX := int(math.Abs(float64(coordinates[i].x)))
				absY := int(math.Abs(float64(coordinates[i].y)))
				fmt.Println("Part 2:", absX+absY)
				os.Exit(0)
			}
		}
	}
}

func addCoordinates(coordinates []Coordinate, direction Pair, start Coordinate) []Coordinate {
	// Adds coordinates to the given slice that have been walked on according to the direction.
	if direction.Direction == "N" {
		for i := 1; i <= direction.Number; i++ {
			coordinates = append(coordinates, Coordinate{x: start.x, y: start.y + i})
		}
	} else if direction.Direction == "S" {
		for i := 1; i <= direction.Number; i++ {
			coordinates = append(coordinates, Coordinate{x: start.x, y: start.y - i})
		}
	} else if direction.Direction == "E" {
		for i := 1; i <= direction.Number; i++ {
			coordinates = append(coordinates, Coordinate{x: start.x + i, y: start.y})
		}
	} else if direction.Direction == "W" {
		for i := 1; i <= direction.Number; i++ {
			coordinates = append(coordinates, Coordinate{x: start.x - i, y: start.y})
		}
	}

	return coordinates
}

func rotate(direction string, points [4]string) [4]string {
	// Rotates the points based on the given direction.
	if direction == "R" {
		first := points[0]
		copy(points[:], points[1:])
		points[len(points)-1] = first
	} else if direction == "L" {
		last := points[len(points)-1]
		copy(points[1:], points[:])
		points[0] = last
	}

	return points
}

func constructAbsoluteDirections(relative_directions []string) []Pair {
	// Constructs absolute directions from relative directions.
	var absolute_directions []Pair
	fourPoints := [4]string{"N", "E", "S", "W"}
	for _, instruction := range relative_directions {
		letter := instruction[0]
		number, _ := strconv.Atoi(instruction[1:])
		fourPoints = rotate(string(letter), fourPoints)
		absolute_directions = append(absolute_directions, Pair{fourPoints[0], number})
	}

	return absolute_directions
}

func getInput() []string {
	// Reads and returns instructions from "input.txt".

	var instructions []string
	readFile, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		parts := strings.Split(line, ",")
		for _, part := range parts {
			instructions = append(instructions, strings.TrimSpace(part))
		}
	}
	readFile.Close()

	return instructions
}
