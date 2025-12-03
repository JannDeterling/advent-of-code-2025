package day2

import (
	"fmt"
	"github.com/JannDeterling/advent-of-code-2025/internal/util"
	"strconv"
)

func RunPart2() {
	ranges := util.ReadCommaSeparatedInput("./input/day2.txt")
	result := 0
	for _, arange := range ranges {
		for _, number := range GetNumbersInRange(arange) {
			if !IsStrictlyValid(number) {
				intNumber, err := strconv.Atoi(number)
				if err != nil {
					panic(err)
				}
				result += intNumber
			}
		}
	}
	fmt.Println(result)
}

func IsStrictlyValid(id string) bool {

	if id[0:1] == "0" {
		return false
	}
	pattern := ""

	for i, char := range id {
		if i < len(id)/2 {
			pattern += string(char)
			if containsOnly(id, pattern) {
				return false
			}
		}
	}
	return true
}

func containsOnly(source string, pattern string) bool {
	sourceLenght := len(source)
	patternLenght := len(pattern)
	if sourceLenght%patternLenght != 0 {
		return false
	} else {
		for i := 0; i < sourceLenght; i += patternLenght {
			if pattern != source[i:i+patternLenght] {
				return false
			}
		}
	}
	return true
}
