package day9

import (
	"slices"
	"testing"
)

func TestPart1ParsePoints(t *testing.T) {
	expected := Points{
		{X: 10, Y: 9},
		{X: 12, Y: 18},
		{X: 15, Y: 25},
	}
	input := []string{
		"10,9",
		"12,18",
		"15,25",
	}

	result := parsePoints(input)
	if !slices.Equal(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}

}

func TestPart1BuildRectangles(t *testing.T) {
	input := Points{
		{X: 10, Y: 9},
		{X: 10, Y: 10},
		{X: 12, Y: 18},
		{X: 15, Y: 25},
	}
	expected := Rectangles{
		{P1: Point{X: 10, Y: 9}, P2: Point{X: 12, Y: 18}, Area: float64(3 * 10)},
		{P1: Point{X: 10, Y: 9}, P2: Point{X: 15, Y: 25}, Area: float64(6 * 17)},
		{P1: Point{X: 10, Y: 10}, P2: Point{X: 12, Y: 18}, Area: float64(3 * 9)},
		{P1: Point{X: 10, Y: 10}, P2: Point{X: 15, Y: 25}, Area: float64(6 * 16)},
		{P1: Point{X: 12, Y: 18}, P2: Point{X: 15, Y: 25}, Area: float64(4 * 8)},
	}

	result := buildPossibleRectangles(input)
	if !slices.Equal(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestPart1SortRectanglesByArea(t *testing.T) {
	input := Rectangles{
		{P1: Point{X: 10, Y: 9}, P2: Point{X: 12, Y: 18}, Area: float64(2 * 9)},
		{P1: Point{X: 10, Y: 9}, P2: Point{X: 15, Y: 25}, Area: float64(5 * 16)},
		{P1: Point{X: 12, Y: 18}, P2: Point{X: 15, Y: 25}, Area: float64(3 * 7)},
	}
	expected := Rectangles{
		{P1: Point{X: 10, Y: 9}, P2: Point{X: 15, Y: 25}, Area: float64(5 * 16)},
		{P1: Point{X: 12, Y: 18}, P2: Point{X: 15, Y: 25}, Area: float64(3 * 7)},
		{P1: Point{X: 10, Y: 9}, P2: Point{X: 12, Y: 18}, Area: float64(2 * 9)},
	}
	input.sortByArea()
	result := input
	if !slices.Equal(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}

}
