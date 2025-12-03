package util

import (
	"bufio"
	"os"
	"strings"
)

func ReadInput(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func ReadCommaSeparatedInput(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, value := range strings.Split(scanner.Text(), ",") {
			lines = append(lines, value)
		}
	}
	return lines
}
