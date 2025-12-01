package day1

import (
	"fmt"
	"strconv"
)

func RunPart2() {
	inputs := readInput("./input/day1.txt")
	fmt.Printf("Input Lenght %d\n", len(inputs))
	var currentLockIndex int = 50
	password := 0
	for i, input := range inputs {
		newIndex, fullRounds := moveLockCarefully(input, currentLockIndex)
		password += fullRounds
		currentLockIndex = newIndex
		if currentLockIndex == 0 {
			password += 1
		}
		fmt.Printf("Current Lock Index at step %d (%s): %d\n", i, input, currentLockIndex)
	}
	fmt.Printf("Password is %d\n", password)
}

func moveLockCarefully(input string, currentLockIndex int) (int, int) {
	var newIndex int
	var fullRoundCount int = 0
	direction := input[0:1]
	turns, err := strconv.Atoi(input[1:])
	if err != nil {
		panic(err)
	}
	if turns > 99 {
		fullRoundCount += turns / 100
		turns = turns % 100
	}
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
		if newIndex > 0 {
			fullRoundCount += 1
		}
	}
	if newIndex < 0 {
		if currentLockIndex > 0 {
			fullRoundCount += 1
		}
		newIndex = newIndex + 1 + 99
	}
	return newIndex, fullRoundCount
}
