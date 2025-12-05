package day5

import (
	"testing"
)

func TestPart2CalculateDeltaOfAllRanges(t *testing.T) {
	expected := 14
	ranges := []Range{
		{Start: 3, End: 5},
		{Start: 10, End: 14},
		{Start: 16, End: 20},
		{Start: 12, End: 18},
	}

	result := calculateDeltaOfAllRanges(ranges)
	if result != expected {
		t.Errorf("got %v, want %v", result, expected)
	}

}

func TestPart2IsOverlapping(t *testing.T) {
	tests := []struct {
		name     string
		prev     Range
		new      Range
		expected bool
	}{
		{"case 1", Range{Start: 16, End: 20}, Range{Start: 12, End: 18}, true},
		{"case 2", Range{Start: 1, End: 4}, Range{Start: 2, End: 8}, true},
		{"case 3", Range{Start: 3, End: 5}, Range{Start: 10, End: 14}, false},
		{"case 4", Range{Start: 10, End: 14}, Range{Start: 8, End: 9}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isOverlapping(tt.prev, tt.new)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
