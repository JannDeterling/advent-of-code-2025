package day4

import (
	"fmt"
	"strings"

	"github.com/JannDeterling/advent-of-code-2025/internal/util"
)

func RunPart1() {
	input := util.ReadInput("./input/day4.txt")
	result := FindAccessibleRolls(input)
	fmt.Println(result)
}

func FindAccessibleRolls(input []string) int {
	accessible := 0
	for y := 0; y < len(input); y++ {
		for x, roll := range input[y] {
			if roll == '@' {
				if (x == 0 && y == 0) || (x == 0 && y == len(input)-1) || (x == len(input[y])-1 && y == 0) || (x == len(input[y])-1 && y == len(input)-1) {
					accessible += 1
				} else {
					if countRollsInPosition(input, x, y) < 4 {
						accessible += 1
					}
				}
			} else {
				continue
			}
		}
	}
	return accessible
}

func countRollsInPosition(input []string, x int, y int) int {
	if y == 0 {
		//top
		//check x-1 x+1
		//check y+1 x-1 x x+2
		area := input[y][x-1:x+2] + input[y+1][x-1:x+2] // +2 since +1 would not be inclusive
		return strings.Count(area, "@") - 1             // -1 due to x being an @
	} else if y == len(input)-1 {
		// bottom
		//check x-1 x+1
		//check y-1 x-1 x x+1
		area := input[y][x-1:x+2] + input[y-1][x-1:x+2]
		return strings.Count(area, "@") - 1 // -1 due to x being an @
	} else if x == 0 {
		// left
		//check x+1
		//check y-1 x x+1
		//check y+1 x x+1
		area := input[y][x:x+2] + input[y-1][x:x+2] + input[y+1][x:x+2]
		return strings.Count(area, "@") - 1 // -1 due to x being an @
	} else if x == len(input[y])-1 {
		// right
		//check x-1
		//check y-1 x x-1
		//check y+1 x x-1
		area := input[y][x-1:x+1] + input[y-1][x-1:x+1] + input[y+1][x-1:x+1]
		return strings.Count(area, "@") - 1 // -1 due to x being an @
	} else {
		// 8 arround
		area := input[y][x-1:x+2] + input[y-1][x-1:x+2] + input[y+1][x-1:x+2]
		return strings.Count(area, "@") - 1 // -1 due to x being an @
	}
}
