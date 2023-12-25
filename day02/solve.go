package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	keypad1 := create_keypad_part1()
	keypad2 := create_keypad_part2()
	file_contents := getInput()
	fmt.Println("Part 1: ", process(keypad1, file_contents))
	fmt.Println("Part 2: ", process(keypad2, file_contents))
}

func process(keypad map[int]*Key, instructions []string) string {
	key := keypad[5]
	code := ""
	for _, instruction := range instructions {
		for _, direction := range instruction {
			key = key.move(string(direction))
		}
		code += button_to_key(key.val)
	}
	return code
}

func button_to_key(button int) string {
	switch button {
	case 10:
		return "A"
	case 11:
		return "B"
	case 12:
		return "C"
	case 13:
		return "D"
	default:
		return strconv.Itoa(button)
	}
}

func create_keypad_part1() map[int]*Key {
	kp := make(map[int]*Key)
	for i := 1; i < 10; i++ {
		kp[i] = &Key{val: i}
	}

	kp[1].up, kp[1].down, kp[1].left, kp[1].right = kp[1], kp[4], kp[1], kp[2]
	kp[2].up, kp[2].down, kp[2].left, kp[2].right = kp[2], kp[5], kp[1], kp[3]
	kp[3].up, kp[3].down, kp[3].left, kp[3].right = kp[3], kp[6], kp[2], kp[3]
	kp[4].up, kp[4].down, kp[4].left, kp[4].right = kp[1], kp[7], kp[4], kp[5]
	kp[5].up, kp[5].down, kp[5].left, kp[5].right = kp[2], kp[8], kp[4], kp[6]
	kp[6].up, kp[6].down, kp[6].left, kp[6].right = kp[3], kp[9], kp[5], kp[6]
	kp[7].up, kp[7].down, kp[7].left, kp[7].right = kp[4], kp[7], kp[7], kp[8]
	kp[8].up, kp[8].down, kp[8].left, kp[8].right = kp[5], kp[8], kp[7], kp[9]
	kp[9].up, kp[9].down, kp[9].left, kp[9].right = kp[6], kp[9], kp[8], kp[9]

	return kp
}

func create_keypad_part2() map[int]*Key {
	kp := make(map[int]*Key)
	for i := 1; i < 14; i++ {
		kp[i] = &Key{val: i}
	}

	kp[1].up, kp[1].down, kp[1].left, kp[1].right = kp[1], kp[3], kp[1], kp[1]
	kp[2].up, kp[2].down, kp[2].left, kp[2].right = kp[2], kp[6], kp[2], kp[3]
	kp[3].up, kp[3].down, kp[3].left, kp[3].right = kp[1], kp[7], kp[2], kp[4]
	kp[4].up, kp[4].down, kp[4].left, kp[4].right = kp[4], kp[8], kp[3], kp[4]
	kp[5].up, kp[5].down, kp[5].left, kp[5].right = kp[5], kp[5], kp[5], kp[6]
	kp[6].up, kp[6].down, kp[6].left, kp[6].right = kp[2], kp[10], kp[5], kp[7]
	kp[7].up, kp[7].down, kp[7].left, kp[7].right = kp[3], kp[11], kp[6], kp[8]
	kp[8].up, kp[8].down, kp[8].left, kp[8].right = kp[4], kp[12], kp[7], kp[9]
	kp[9].up, kp[9].down, kp[9].left, kp[9].right = kp[9], kp[9], kp[8], kp[9]
	kp[10].up, kp[10].down, kp[10].left, kp[10].right = kp[6], kp[10], kp[10], kp[11]
	kp[11].up, kp[11].down, kp[11].left, kp[11].right = kp[7], kp[13], kp[10], kp[12]
	kp[12].up, kp[12].down, kp[12].left, kp[12].right = kp[8], kp[12], kp[11], kp[12]
	kp[13].up, kp[13].down, kp[13].left, kp[13].right = kp[11], kp[13], kp[13], kp[13]

	return kp
}

type Key struct {
	val   int
	up    *Key
	down  *Key
	left  *Key
	right *Key
}

func (input_key *Key) move(direction string) *Key {
	switch direction {
	case "U":
		return input_key.up
	case "D":
		return input_key.down
	case "L":
		return input_key.left
	case "R":
		return input_key.right
	default:
		return input_key
	}
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
