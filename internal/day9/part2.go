package day9

import (
	"fmt"
	"math"

	"github.com/JannDeterling/advent-of-code-2025/internal/util"
)

type Line struct {
	P1 Point
	P2 Point
}

type Lines []Line

// create all boundries of the given points
// create all reactangles by checking if they intersect with the boundries
// P1 and P2 of the rectangle need to be mirrored so that I get all points and all lines that can intersect
// continue as in part1
func findBoundries(points Points) Lines {
	lines := Lines{}
	var p1, p2 Point
	for i := 0; i < len(points); i++ {
		p1 = points[i]
		if i < len(points)-1 {
			p2 = points[i+1]
		} else {
			p2 = points[0]
		}
		lines = append(lines, Line{P1: p1, P2: p2})
	}
	return lines
}

func (l Line) intersects(lines Lines) bool {
	for _, line := range lines {
		if line.intersect(l) {
			fmt.Printf("Intersected by: %v\n", line)
			return true
		}
	}
	return false
}

func (l Line) intersect(l2 Line) bool {

	o1 := orientation(l.P1, l.P2, l2.P1)
	o2 := orientation(l.P1, l.P2, l2.P2)
	o3 := orientation(l2.P1, l2.P2, l.P1)
	o4 := orientation(l2.P1, l2.P2, l.P2)

	// General case: segments intersect if orientations are different
	if o1 != o2 && o3 != o4 {
		return true
	}

	if (o1 == 0 && o2 == 0) || (o3 == 0 && o4 == 0) {
		return false
	}

	return false
}

func orientation(p, q, r Point) int {
	val := (q.Y-p.Y)*(r.X-q.X) - (q.X-p.X)*(r.Y-q.Y)
	if val == 0 {
		return 0 // collinear
	}
	if val > 0 {
		return 1 // clockwise
	}
	return 2 // counterclockwise
}

func onSegment(p, q, r Point) bool {
	return q.X <= max(p.X, r.X) && q.X >= min(p.X, r.X) &&
		q.Y <= max(p.Y, r.Y) && q.Y >= min(p.Y, r.Y)
}

/*
*
......
....P. (X 5 Y 1) -> (X 5 Y 3)  | (X 5 Y 3) -> (X 2 Y 3) | (X 2 Y 3) -> (X 2 Y 1) | (X 2 Y 1) -> (X 5 Y 1)
......
.P....
*/
func (r Rectangle) getLines() Lines {
	return Lines{
		Line{P1: r.P1, P2: Point{X: r.P1.X, Y: r.P2.Y}},
		Line{P1: Point{X: r.P1.X, Y: r.P2.Y}, P2: r.P2},
		Line{P1: r.P2, P2: Point{X: r.P2.X, Y: r.P1.Y}},
		Line{P1: Point{X: r.P2.X, Y: r.P1.Y}, P2: r.P1},
	}
}

func buildPossibleRectanglesWithBoundries(points Points, boundries Lines) Rectangles {
	rectangles := Rectangles{}
	memo := make(map[string]bool)
	for _, p1 := range points {
		for _, p2 := range points {
			if p1 != p2 && p1.X != p2.X && p1.Y != p2.Y {
				fmt.Printf("%+v\n", Rectangle{P1: p1, P2: p2})
				key := fmt.Sprintf("%+v-%+v", p1, p2)
				reverseKey := fmt.Sprintf("%+v-%+v", p2, p1)
				if !memo[key] && !memo[reverseKey] {
					optRect := Rectangle{P1: p1, P2: p2}
					//optRect.render(boundries)
					if optRect.isInsideBoundries(boundries) {
						line1 := math.Abs(float64(p1.X)-float64(p2.X)) + 1 // +1 because edge points
						line2 := math.Abs(float64(p1.Y)-float64(p2.Y)) + 1
						rectangles = append(rectangles, Rectangle{P1: p1, P2: p2, Area: line1 * line2})
						memo[key] = true
						memo[reverseKey] = true
					}
				}
			}
		}
	}
	return rectangles
}

func (r Rectangle) render(boundries Lines) {
	rectLines := r.getLines()
	for y := 0; y < 10; y++ {
		fmt.Printf("%d: ", y)
		for x := 0; x < 15; x++ {
			runeToPrint := "."
			for _, line := range rectLines {
				if line.P1.X == x && line.P1.Y == y || line.P2.X == x && line.P2.Y == y {
					runeToPrint = "@"
				} else if line.P1.X < x && line.P2.X > x && line.P1.Y == y || line.P2.X < x && line.P1.X > x && line.P2.Y == y {
					runeToPrint = "-"
				} else if line.P1.Y < y && line.P2.Y > y && line.P1.X == x || line.P2.Y < y && line.P1.Y > y && line.P2.X == x {
					runeToPrint = "|"
				}
			}
			for _, line := range boundries {
				if line.P1.X == x && line.P1.Y == y || line.P2.X == x && line.P2.Y == y {
					runeToPrint = "#"
				} else if line.P1.X < x && line.P2.X > x && line.P1.Y == y || line.P2.X < x && line.P1.X > x && line.P2.Y == y {
					runeToPrint = "~"
				} else if line.P1.Y < y && line.P2.Y > y && line.P1.X == x || line.P2.Y < y && line.P1.Y > y && line.P2.X == x {
					runeToPrint = "!"
				}
			}
			fmt.Printf("%s", runeToPrint)
		}
		fmt.Printf("\n")
	}
}

func (r Rectangle) isInsideBoundries(boundries Lines) bool {
	return r.P1.isInsideBoundries(boundries) && r.P2.isInsideBoundries(boundries)
}

func (p Point) isInsideBoundries(boundries Lines) bool {
	inside := false
	for _, boundry := range boundries {
		if p.Y > min(boundry.P1.Y, boundry.P2.Y) && p.Y < max(boundry.P1.Y, boundry.P2.Y) {
			if p.X <= max(boundry.P1.X, boundry.P2.X) {
				var xinters int
				if boundry.P1.Y != boundry.P2.Y {
					xinters = (p.Y-boundry.P1.Y)*(boundry.P1.X-boundry.P2.X)/(boundry.P2.Y-boundry.P1.Y) + boundry.P1.X
				}
				if boundry.P1.Y != boundry.P2.Y || p.X <= xinters {
					inside = !inside
				}
			}
		}
	}
	return inside
}

func RunPart2() {
	input := util.ReadInput("./input/day9-test.txt")
	points := parsePoints(input)
	boundries := findBoundries(points)
	fmt.Printf("Boundries: %v\n", boundries)
	rectangles := buildPossibleRectanglesWithBoundries(points, boundries)
	rectangles.sortByArea()
	fmt.Printf("%+v\n", rectangles)
	fmt.Printf("%f\n", rectangles[0].Area)
	rectangles[0].render(boundries)
}
