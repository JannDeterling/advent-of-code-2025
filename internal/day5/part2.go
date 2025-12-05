package day5

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/JannDeterling/advent-of-code-2025/internal/util"
)

func RunPart2() {
	input := util.ReadInput("./input/day5.txt")
	ranges := []Range{}
	for _, line := range input {
		if line == "" {
			break
		}
		split := strings.Split(line, "-")
		start, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		ranges = append(ranges, Range{Start: start, End: end})
	}
	count := calculateDeltaOfAllRanges(ranges)
	fmt.Printf("all fresh IDs: %d\n", count)
}

type Range struct {
	Start int
	End   int
}

func calculateDeltaOfAllRanges(ranges []Range) int {
	var prevRange Range
	totalDelta := 0
	slices.SortFunc(ranges, func(a, b Range) int {
		if a.Start != b.Start {
			return a.Start - b.Start
		}
		if a.End != b.End {
			return a.End - b.End
		}
		return 0
	})
	for i, aRange := range ranges {

		if i == 0 {
			prevRange = aRange
		}

		if isOverlapping(prevRange, aRange) {
			if prevRange.Start > aRange.Start {
				prevRange.Start = aRange.Start
			}
			if prevRange.End < aRange.End {
				prevRange.End = aRange.End
			}
		} else {
			totalDelta += (prevRange.End - prevRange.Start) + 1
			prevRange = aRange
		}
	}
	totalDelta += (prevRange.End - prevRange.Start) + 1
	return totalDelta
}

func isOverlapping(prev Range, new Range) bool {
	if (prev.Start >= new.Start && prev.Start <= new.End) || (prev.Start <= new.Start && new.Start <= prev.End) {
		return true
	} else if (new.End >= prev.Start && new.End <= prev.End) || (prev.End >= new.Start && prev.End <= new.End) {
		return true
	} else {
		return false
	}
}
