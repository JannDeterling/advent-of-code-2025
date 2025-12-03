package day3

import (
	"fmt"
	"strconv"

	"github.com/JannDeterling/advent-of-code-2025/internal/util"
)

func RunPart2() {
	input := util.ReadInput("./input/day3.txt")
	sum := 0
	for _, bank := range input {
		sum += FindHighestJoltageNoLimits(bank)
	}
	fmt.Println(sum)
}

func FindHighestJoltageNoLimits(bank string) int {

	charges := ""
	chargeIndex := -1
	for i := range 12 {
		highestCharge := 0
		for batteryIndex, battery := range bank {
			batteryCharge, err := strconv.Atoi(string(battery))
			if err != nil {
				panic(err)
			}
			if highestCharge < batteryCharge && batteryIndex < (len(bank)-(11-i)) && chargeIndex < batteryIndex {
				highestCharge = batteryCharge
				chargeIndex = batteryIndex
			}
		}
		charges = charges + fmt.Sprintf("%d", highestCharge)
	}

	result, err := strconv.Atoi(charges)
	if err != nil {
		panic(err)
	}
	return result
}
