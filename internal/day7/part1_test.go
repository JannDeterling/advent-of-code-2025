package day7

import (
	"slices"
	"testing"
)

func TestMoveBeams(t *testing.T) {
	tests := []struct {
		name           string
		beamsFromAbove []int
		input          string
		expected       string
		expectedBeams []int
		expectedSplits int
	}{
		{"case 1", []int{7}, "...............", ".......|.......",  []int{7}, 0},
		{"case 2", []int{7}, ".......^.......", "......|^|......",  []int{6,8}, 1},
		{"case 3", []int{6, 8}, "......^.^......", ".....|^|^|.....", []int{5,7,9}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, beams, splits := MoveBeams(tt.input, tt.beamsFromAbove)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
			if !slices.Equal(beams, tt.expectedBeams) {
				t.Errorf("got beams %v, want %v", beams, tt.expectedBeams)
			}
			if splits != tt.expectedSplits {
				t.Errorf("got splits %v, want %v", beams, tt.expectedBeams)
			}
		})
	}
}

func TestFindStartPoint(t *testing.T) {
	//given
	input := ".......S......."
	expected := 7
	//when
	result := FindStartPoint(input)
	//then
	if result != expected {
		t.Errorf("got %v, want %v", result, expected)
	}

}
