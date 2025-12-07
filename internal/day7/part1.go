package day7

import (
	"fmt"
	"strings"

	"github.com/JannDeterling/advent-of-code-2025/internal/util"
)

func RunPart1() {
	input := util.ReadInput("./input/day7.txt")

	indices := []int{}
	startIndex := 0
	splits := 0
	result := []string{}
	for i, line := range input {
		if i == 0 {
			startIndex = FindStartPoint(line)
			indices = append(indices, FindStartPoint(line))
			fmt.Println(line)
		} else {
			newLine, newIndices, newSplits := MoveBeams(line, indices)
			fmt.Println(newLine)
			result = append(result, newLine)
			indices = newIndices
			splits += newSplits
		}
	}
	fmt.Println(splits)
	fmt.Println(FindAllPossibleBeamPaths(result, startIndex))
}

func FindStartPoint(input string) int {
	return strings.Index(input, "S")
}

func MoveBeams(currentLine string, beamsAbove []int) (string, []int, int) {
	result := currentLine
	splits := 0
	for _, beam := range beamsAbove {
		newLine, newSplits := placeBeam(currentLine, beam)
		result = mergeLines(result, newLine)
		splits += newSplits
	}
	return result, findBeams(result), splits
}

func findBeams(currentLine string) []int {
	indices := []int{}
	for i, c := range currentLine {
		if c == '|' {
			indices = append(indices, i)
		}
	}
	return indices
}

func placeBeam(currentLine string, beamIndex int) (string, int) {
	runes := []rune(currentLine)
	location := runes[beamIndex]
	switch location {
	case '.':
		runes[beamIndex] = '|'
		return string(runes), 0
	case '|':
		return currentLine, 0
	case '^':
		left, leftSplits := placeBeam(currentLine, beamIndex-1)
		right, righSplits := placeBeam(currentLine, beamIndex+1)
		return mergeLines(left, right), 1 + leftSplits + righSplits
	default:
		panic(fmt.Errorf("This should not happen!"))
	}
}

func mergeLines(leftLine string, rightLine string) string {
	left := []rune(leftLine)
	right := []rune(rightLine)
	merged := make([]rune, len(left))
	copy(merged, left)
	for i := 0; i < len(left); i++ {
		if merged[i] != right[i] && right[i] == '|' {
			merged[i] = right[i]
		}
	}
	return string(merged)
}
