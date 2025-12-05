package day5

import "testing"

func TestPart1IsInRange(t *testing.T) {
	tests := []struct {
		name     string
		aRange   string
		input    string
		expected bool
	}{
		{"case 1", "11-13", "11", true},
		{"case 2", "11-13", "12", true},
		{"case 3", "11-13", "13", true},
		{"case 4", "11-13", "14", false},
		{"case 5", "11-13", "10", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsInRange(tt.aRange, tt.input)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
