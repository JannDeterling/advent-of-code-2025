package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/JannDeterling/advent-of-code-2025/internal/util"
)

func RunPart1() {
	ranges := util.ReadCommaSeparatedInput("./input/day2.txt")
	result := 0
	for _, arange := range ranges {
		for _, number := range GetNumbersInRange(arange) {
			if !IsValid(number) {
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

func IsValid(id string) bool {
	isValid := true
	if id[0:1] == "0" {
		return false
	}
	duplicated := ""
	lenght := len(id)
	for i, char := range id {
		if i < lenght/2 {
			duplicated += string(char)
			if duplicated == id[i+1:] {
				isValid = false
			}
		}
	}
	return isValid
}

func GetNumbersInRange(aRange string) []string {
	allNumbers := []string{}
	parts := strings.Split(aRange, "-")
	start, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	end, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	for i := start; i < end; i++ {
		allNumbers = append(allNumbers, fmt.Sprint(i))
	}
	allNumbers = append(allNumbers, fmt.Sprint(end))
	return allNumbers
}
