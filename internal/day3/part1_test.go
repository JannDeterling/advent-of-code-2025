package day3

import "testing"

func TestPart1FindHighestJoltage(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"case 1", "987654321111111", 98},
		{"case 2", "811111111111119", 89},
		{"case 3", "234234234234278", 78},
		{"case 4", "818181911112111", 92},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindHighestJoltage(tt.input)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
