package day4

import "testing"

func TestFindAccessibleRolls(t *testing.T) {
	//given
	expected := 13
	input := []string{
		"..@@.@@@@.",
		"@@@.@.@.@@",
		"@@@@@.@.@@",
		"@.@@@@..@.",
		"@@.@@@@.@@",
		".@@@@@@@.@",
		".@.@.@.@@@",
		"@.@@@.@@@@",
		".@@@@@@@@.",
		"@.@.@@@.@.",
	}
	//when
	result := FindAccessibleRolls(input)

	//then
	if result != expected {
		t.Errorf("got %d, wwanted %d", result, expected)
	}

}
