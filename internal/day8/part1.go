package day8

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/JannDeterling/advent-of-code-2025/internal/util"
)

type Circuit []Coordinate

type Circuits []Circuit

type JunctionBox struct {
	Coordinate Coordinate
	distances  Distances
}

type JunctionBoxes []JunctionBox

func (b JunctionBoxes) allDistinctDistances() Distances {
	memo := make(map[string]bool)
	allDistances := Distances{}
	for _, junctionBox := range b {
		for _, distance := range junctionBox.distances {
			key := fmt.Sprintf("%+v-%+v", distance.P1, distance.P2)
			reverseKey := fmt.Sprintf("%+v-%+v", distance.P2, distance.P1)
			if !memo[key] || !memo[reverseKey] {
				allDistances = append(allDistances, distance)
				memo[key] = true
				memo[reverseKey] = true
			}
		}
	}
	allDistances.sort()
	return allDistances
}

type Coordinate struct {
	X int
	Y int
	Z int
}

func (c Coordinate) isEqual(other Coordinate) bool {
	return c.X == other.X && c.Y == other.Y && c.Z == other.Z
}

type Distance struct {
	P1       Coordinate
	P2       Coordinate
	Distance float64
}

type Distances []Distance

func (d Distances) sort() {
	slices.SortFunc(d, func(a, b Distance) int {
		return int(a.Distance - b.Distance)
	})
}

func connectJunctionBoxes(junctionBoxes JunctionBoxes, limit int) Circuits {
	circuitsLookUp := make(map[Coordinate]int)
	circuits := Circuits{}
	for i, junctionBox := range junctionBoxes {
		circuits = append(circuits, Circuit{junctionBox.Coordinate})
		circuitsLookUp[junctionBox.Coordinate] = i
	}
	distances := junctionBoxes.allDistinctDistances()
	for i, distance := range distances {
		if i == limit {
			break
		}
		circuitP1Index := circuitsLookUp[distance.P1]
		circuitP2Index := circuitsLookUp[distance.P2]
		if circuitP1Index >= 0 && circuitP2Index >= 0 && circuitP1Index != circuitP2Index {
			circuits[circuitP1Index] = append(circuits[circuitP1Index], circuits[circuitP2Index]...)
			for _, coord := range circuits[circuitP2Index] {
				circuitsLookUp[coord] = circuitP1Index
			}
			circuits[circuitP2Index] = nil
		}
	}
	return slices.DeleteFunc(circuits, func(circuit Circuit) bool {
		return circuit == nil
	})
}

func createJunctionBoxes(coordinates []Coordinate) []JunctionBox {
	var junctionBoxes []JunctionBox
	for _, coordinate := range coordinates {
		junctionBoxes = append(junctionBoxes, distanceToOtherJunctionBoxes(coordinate, coordinates))
	}
	return junctionBoxes
}

func distanceToOtherJunctionBoxes(coordinate Coordinate, coordinates []Coordinate) JunctionBox {
	var distances []Distance
	for _, otherCoordinate := range coordinates {
		if coordinate != otherCoordinate {
			distances = append(distances, Distance{P1: coordinate, P2: otherCoordinate, Distance: distanceBetween(coordinate, otherCoordinate)})
		}
	}
	return JunctionBox{Coordinate: coordinate, distances: distances}
}

func distanceBetween(p1 Coordinate, p2 Coordinate) float64 {
	return math.Sqrt(
		math.Pow(float64(p1.X-p2.X), 2) +
			math.Pow(float64(p1.Y-p2.Y), 2) +
			math.Pow(float64(p1.Z-p2.Z), 2),
	)
}

func toInt(in string) int {
	value, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return value
}

func createCoordinates(input []string) []Coordinate {
	coordinates := []Coordinate{}
	for _, line := range input {
		split := strings.Split(line, ",")
		coordinates = append(coordinates, Coordinate{
			X: toInt(split[0]),
			Y: toInt(split[1]),
			Z: toInt(split[2]),
		})
	}
	return coordinates
}

func RunPart1() {
	input := util.ReadInput("./input/day8.txt")
	coordinates := createCoordinates(input)
	junctionBoxes := createJunctionBoxes(coordinates)
	circuits := connectJunctionBoxes(junctionBoxes, 1000)
	slices.SortFunc(circuits, func(a, b Circuit) int {
		return len(b) - len(a)
	})
	fmt.Printf("----------------------------------------\n")
	fmt.Printf("%d\n", len(circuits[0]))
	fmt.Printf("%d\n", len(circuits[1]))
	fmt.Printf("%d\n", len(circuits[2]))
	fmt.Printf("total: %d\n", len(circuits[0])*len(circuits[1])*len(circuits[2]))
	fmt.Printf("----------------------------------------\n")
}
