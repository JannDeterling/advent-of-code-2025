package day9

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/JannDeterling/advent-of-code-2025/internal/util"
)

type Point struct {
	X int
	Y int
}

type Rectangle struct {
	P1   Point
	P2   Point
	Area float64
}
type Rectangles []Rectangle

func (r Rectangles) sortByArea() {
	slices.SortFunc(r, func(a, b Rectangle) int {
		return int(b.Area - a.Area)
	})
}

type Points []Point

func parsePoints(input []string) Points {
	points := make(Points, len(input))
	for i, line := range input {
		split := strings.Split(line, ",")
		points[i] = Point{X: util.ToInt(split[0]), Y: util.ToInt(split[1])}
	}
	return points
}

func buildPossibleRectangles(points Points) Rectangles {
	rectangles := Rectangles{}
	memo := make(map[string]bool)
	for _, p1 := range points {
		for _, p2 := range points {
			if p1 != p2 && p1.X != p2.X && p1.Y != p2.Y {
				key := fmt.Sprintf("%+v-%+v", p1, p2)
				reverseKey := fmt.Sprintf("%+v-%+v", p2, p1)
				if !memo[key] && !memo[reverseKey] {
					line1 := math.Abs(float64(p1.X)-float64(p2.X)) + 1 // +1 because edge points
					line2 := math.Abs(float64(p1.Y)-float64(p2.Y)) + 1
					rectangles = append(rectangles, Rectangle{P1: p1, P2: p2, Area: line1 * line2})
					memo[key] = true
					memo[reverseKey] = true
				}
			}
		}
	}
	return rectangles
}

func RunPart1() {
	input := util.ReadInput("./input/day9-test.txt")
	points := parsePoints(input)
	rectangles := buildPossibleRectangles(points)
	rectangles.sortByArea()
	fmt.Printf("%f\n", rectangles[0].Area)
	fmt.Printf("%+v\n", rectangles)
}
