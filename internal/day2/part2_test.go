package day2

import (
	"testing"
)

func TestPart2IsInvalidWhen(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"case 1", "11", false},
		{"case 2", "2222", false},
		{"case 3", "123123", false},
		{"case 4", "10101010", false},
		{"case 5", "0123", false},
		{"case 6", "222", false},
		{"case 7", "2121212121", false},
		{"case 8", "123123123", false},
		{"case 9", "565656", false},
		{"case 10", "824824824", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsStrictlyValid(tt.input)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPart2IsValidWhen(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"case 1", "123321", true},
		{"case 2", "123", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsStrictlyValid(tt.input)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
