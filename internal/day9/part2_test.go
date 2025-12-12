package day9

import (
	"slices"
	"testing"
)

func TestPart2FindBoundries(t *testing.T) {
	input := Points{
		{X: 10, Y: 9},
		{X: 12, Y: 9},
		{X: 12, Y: 25},
	}
	expected := Lines{
		{P1: Point{X: 10, Y: 9}, P2: Point{X: 12, Y: 9}},
		{P1: Point{X: 12, Y: 9}, P2: Point{X: 12, Y: 25}},
		{P1: Point{X: 12, Y: 25}, P2: Point{X: 10, Y: 9}},
	}

	result := findBoundries(input)
	if !slices.Equal(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}

}

func TestPart2Intersect(t *testing.T) {
	tests := []struct {
		name     string
		l1       Line
		l2       Line
		expected bool
	}{
		{"Case 1", Line{P1: Point{X: 1, Y: 2}, P2: Point{X: 5, Y: 2}}, Line{P1: Point{X: 3, Y: 1}, P2: Point{X: 3, Y: 3}}, true},
		{"Case 2", Line{P1: Point{X: 1, Y: 20}, P2: Point{X: 1, Y: 10}}, Line{P1: Point{X: -1, Y: 15}, P2: Point{X: 5, Y: 15}}, true},
		{"Case 3", Line{P1: Point{X: 1, Y: 2}, P2: Point{X: 5, Y: 2}}, Line{P1: Point{X: 3, Y: 3}, P2: Point{X: 3, Y: 3}}, false},
		{"Case 4", Line{P1: Point{X: 1, Y: 1}, P2: Point{X: 1, Y: 5}}, Line{P1: Point{X: 3, Y: 1}, P2: Point{X: 3, Y: 5}}, false},
		{"Case 5", Line{P1: Point{X: 1, Y: 1}, P2: Point{X: 1, Y: 5}}, Line{P1: Point{X: 1, Y: 0}, P2: Point{X: 5, Y: 0}}, false},
		{"Case 6", Line{P1: Point{X: 7, Y: 1}, P2: Point{X: 7, Y: 7}}, Line{P1: Point{X: 9, Y: 5}, P2: Point{X: 2, Y: 5}}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.l1.intersect(tt.l2)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPart2RectangleGetLines(t *testing.T) {
	tests := []struct {
		name     string
		r        Rectangle
		expected Lines
	}{
		{"Case 1", Rectangle{P1: Point{X: 10, Y: 9}, P2: Point{X: 12, Y: 18}, Area: float64(3 * 10)}, Lines{
			Line{P1: Point{X: 10, Y: 9}, P2: Point{X: 10, Y: 18}},
			Line{P1: Point{X: 10, Y: 18}, P2: Point{X: 12, Y: 18}},
			Line{P1: Point{X: 12, Y: 18}, P2: Point{X: 12, Y: 9}},
			Line{P1: Point{X: 12, Y: 9}, P2: Point{X: 10, Y: 9}},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.r.getLines()
			if !slices.Equal(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
