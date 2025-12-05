package day5

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/JannDeterling/advent-of-code-2025/internal/util"
)

func RunPart1() {
	input := util.ReadInput("./input/day5.txt")
	values := []string{}
	ranges := []string{}
	isValues := false
	for _, line := range input {

		if line == "" {
			isValues = true
			continue
		}

		if isValues {
			values = append(values, line)
		} else {
			ranges = append(ranges, line)
		}
	}
	count := 0
	for _, value := range values {
		for _, aRange := range ranges {
			if IsInRange(aRange, value) {
				count++
				break
			}
		}
	}
	fmt.Printf("fresh IDs: %d\n", count)
}

func IsInRange(aRange string, input string) bool {
	split := strings.Split(aRange, "-")
	start, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}
	end, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}
	inputToCheck, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	if start <= inputToCheck && inputToCheck <= end {
		return true
	}
	return false
}
