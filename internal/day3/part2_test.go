package day3

import "testing"

func TestPart2FindHighestJoltageNoLimits(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"case 1", "987654321111111", 987654321111},
		{"case 2", "811111111111119", 811111111119},
		{"case 3", "234234234234278", 434234234278},
		{"case 4", "818181911112111", 888911112111},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindHighestJoltageNoLimits(tt.input)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
