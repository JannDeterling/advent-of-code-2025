package day6

import (
	"reflect"
	"testing"
)

func TestPart1ParseInput(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []MathTask
	}{
		{"case 1", []string{"1 2 3", "2 3 1", "3 4 1", "+ * +"}, []MathTask{{Values: []int{3, 2, 1}, Operator: "+", Result: 6}, {Values: []int{4, 3, 2}, Operator: "*", Result: 24}, {Values: []int{1, 1, 3}, Operator: "+", Result: 5}}},
		{"case 2", []string{"4", "2", "3", "*"}, []MathTask{{Values: []int{3, 2, 4}, Operator: "*", Result: 24}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseMathTask(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
