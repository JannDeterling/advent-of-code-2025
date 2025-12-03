package day3

import (
	"fmt"
	"strconv"

	"github.com/JannDeterling/advent-of-code-2025/internal/util"
)

func RunPart1() {
	input := util.ReadInput("./input/day3.txt")
	sum := 0
	for _, bank := range input {
		sum += FindHighestJoltage(bank)
	}
	fmt.Println(sum)
}

func FindHighestJoltage(bank string) int {

	highestCharge := 0
	secondCharge := 0
	highestChargeIndex := 0
	for i, battery := range bank {
		batteryCharge, err := strconv.Atoi(string(battery))
		if err != nil {
			panic(err)
		}
		if highestCharge < batteryCharge && i < (len(bank)-1) {
			highestCharge = batteryCharge
			highestChargeIndex = i
		}
	}
	for _, battery := range bank[highestChargeIndex+1:] {
		batteryCharge, err := strconv.Atoi(string(battery))
		if err != nil {
			panic(err)
		}
		if secondCharge < batteryCharge {
			secondCharge = batteryCharge
		}
	}
	totalCharge := fmt.Sprintf("%d%d", highestCharge, secondCharge)
	result, err := strconv.Atoi(totalCharge)
	if err != nil {
		panic(err)
	}
	return result
}
