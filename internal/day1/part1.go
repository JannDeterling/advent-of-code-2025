package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TestPart1() {
	inputs := readInput("./input/day1-test.txt")
	fmt.Printf("Input Lenght %d\n", len(inputs))
	var currentLockIndex int = 50
	password := 0
	for i, input := range inputs {
		currentLockIndex = moveLock(input, currentLockIndex)
		fmt.Printf("Current Lock Index at step %d: %d\n", i, currentLockIndex)
		if currentLockIndex == 0 {
			password += 1
		}
	}
	fmt.Printf("Password is %d\n", password)
}

func RunPart1() {

	inputs := readInput("./input/day1.txt")
	fmt.Printf("Input Lenght %d\n", len(inputs))
	var currentLockIndex int = 50
	password := 0
	for i, input := range inputs {
		currentLockIndex = moveLock(input, currentLockIndex)
		fmt.Printf("Current Lock Index at step (%s)  %d: %d\n", input, i, currentLockIndex)
		if currentLockIndex == 0 {
			password += 1
		}
	}
	fmt.Printf("Password is %d\n", password)

}

func moveLock(input string, currentLockIndex int) int {
	direction := input[0:1]
	turns, err := strconv.Atoi(input[1:])
	if err != nil {
		panic(err)
	}
	if turns > 99 {
		turns = turns % 100
	}
	var newIndex int
	switch direction {
	case "L":
		newIndex = currentLockIndex - turns
	case "R":
		newIndex = currentLockIndex + turns
	default:
		panic(fmt.Errorf("This direction is not allowed! %s", direction))
	}
	if newIndex > 99 {
		newIndex = newIndex % 100
	}
	if newIndex < 0 {
		newIndex = newIndex + 1 + 99
	}
	return newIndex
}

func readInput(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
