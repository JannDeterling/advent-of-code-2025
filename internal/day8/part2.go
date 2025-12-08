package day8

import (
	"fmt"
	"slices"

	"github.com/JannDeterling/advent-of-code-2025/internal/util"
)

func connectJunctionBoxesFullCircle(junctionBoxes JunctionBoxes) (Circuits, Distance) {
	circuitsLookUp := make(map[Coordinate]int)
	circuits := Circuits{}
	fullCircleDistance := Distance{}
	for i, junctionBox := range junctionBoxes {
		circuits = append(circuits, Circuit{junctionBox.Coordinate})
		circuitsLookUp[junctionBox.Coordinate] = i
	}
	distances := junctionBoxes.allDistinctDistances()
	for _, distance := range distances {
		circuitP1Index := circuitsLookUp[distance.P1]
		circuitP2Index := circuitsLookUp[distance.P2]
		if circuitP1Index >= 0 && circuitP2Index >= 0 && circuitP1Index != circuitP2Index {
			circuits[circuitP1Index] = append(circuits[circuitP1Index], circuits[circuitP2Index]...)
			for _, coord := range circuits[circuitP2Index] {
				circuitsLookUp[coord] = circuitP1Index
			}
			circuits[circuitP2Index] = nil
		}
		if allValuesEqual(circuitsLookUp) && fullCircleDistance == (Distance{}) {
			fullCircleDistance = distance
		}
	}
	return slices.DeleteFunc(circuits, func(circuit Circuit) bool {
		return circuit == nil
	}), fullCircleDistance
}

func allValuesEqual(m map[Coordinate]int) bool {
	if len(m) == 0 {
		return true
	}

	var firstValue int
	first := true

	for _, v := range m {
		if first {
			firstValue = v
			first = false
		} else if v != firstValue {
			return false
		}
	}

	return true
}

func RunPart2() {
	input := util.ReadInput("./input/day8.txt")
	coordinates := createCoordinates(input)
	junctionBoxes := createJunctionBoxes(coordinates)
	circuits, distance := connectJunctionBoxesFullCircle(junctionBoxes)
	slices.SortFunc(circuits, func(a, b Circuit) int {
		return len(b) - len(a)
	})
	fmt.Printf("----------------------------------------\n")
	fmt.Printf("%d\n", len(circuits[0]))
	fmt.Printf("%v\n", distance)
	fmt.Printf("%d\n", distance.P1.X*distance.P2.X)
	fmt.Printf("----------------------------------------\n")

}
