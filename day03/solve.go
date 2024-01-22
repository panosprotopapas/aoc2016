package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file_contents := getInput()
	part1 := part1(file_contents)
	fmt.Println("Part 1: ", part1)
	part2 := part2(file_contents)
	fmt.Println("Part 2: ", part2)
}

func part1(instructions [][]int) int {
	counter := 0
	for _, x := range instructions {
		if (x[0]+x[1] > x[2]) && (x[0]+x[2] > x[1]) && (x[1]+x[2] > x[0]) {
			counter += 1
		}
	}
	return counter
}

func part2(instructions [][]int) int {
	counter := 0
	for x := 0; x < len(instructions); x += 3 {
		for _, y := range []int{0, 1, 2} {
			if (instructions[x][y]+instructions[x+1][y] > instructions[x+2][y]) && 
			(instructions[x][y]+instructions[x+2][y] > instructions[x+1][y]) && 
			(instructions[x+1][y]+instructions[x+2][y] > instructions[x][y]) {
				counter += 1
			}
		}
	}
	return counter
}

func getInput() [][]int {
	// Reads and returns instructions from "input.txt".

	var instructions [][]int
	readFile, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		parts := strings.Fields(line)
		ints := make([]int, len(parts))
		for i, part := range parts {
			intVal, _ := strconv.Atoi(part)
			ints[i] = intVal
		}
		instructions = append(instructions, ints)
	}
	readFile.Close()

	return instructions
}
