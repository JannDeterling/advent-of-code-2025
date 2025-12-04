package day4

import (
	"fmt"
	"github.com/JannDeterling/advent-of-code-2025/internal/util"
)

func RunPart2() {
	input := util.ReadInput("./input/day4.txt")
	rollsAccessible := 0
	for {
		result, newInput := FindAccessibleRollsAgain(input)
		fmt.Printf("%d\n", result)
		if result == 0 {
			break
		}
		rollsAccessible += result
		input = newInput
	}
	fmt.Println(rollsAccessible)
}

func FindAccessibleRollsAgain(input []string) (int, []string) {
	accessible := 0
	newInput := make([]string, len(input))
	copy(newInput, input)
	for y := 0; y < len(input); y++ {
		for x, roll := range input[y] {
			if roll == '@' {
				if (x == 0 && y == 0) || (x == 0 && y == len(input)-1) || (x == len(input[y])-1 && y == 0) || (x == len(input[y])-1 && y == len(input)-1) {
					accessible += 1
					newInput[y] = replaceRoll(newInput[y], x)
				} else {
					if countRollsInPosition(input, x, y) < 4 {
						accessible += 1
						newInput[y] = replaceRoll(newInput[y], x)
					}
				}
			} else {
				continue
			}
		}
	}
	return accessible, newInput
}

func replaceRoll(input string, x int) string {
	runes := []rune(input)
	runes[x] = '.'
	return string(runes)
}
