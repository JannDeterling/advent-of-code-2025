package day2

import (
	"slices"
	"testing"
)

func TestPart1IsInvalidWhen(t *testing.T) {
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
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := IsValid(tt.input)
            if result != tt.expected {
                t.Errorf("got %v, want %v", result, tt.expected)
            }
        })
    }
}

func TestPart1IsValidWhen(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected bool
        }{
        {"case 1", "123321", true},
        {"case 2", "123", true},
        {"case 3", "222", true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := IsValid(tt.input)
            if result != tt.expected {
            	t.Errorf("got %v, want %v", result, tt.expected)
            }
        })
    }
}

func TestGetNumbersInRange(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected []string
    }{
        {"case 1", "1-6", []string{"1", "2", "3", "4", "5", "6"} },
        {"case 2", "1223-1227", []string{"1223", "1224", "1225", "1226", "1227"}},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := GetNumbersInRange(tt.input)
            if !slices.Equal(result, tt.expected){
            	t.Errorf("got %v, want %v", result, tt.expected)
            }
        })
    }
}