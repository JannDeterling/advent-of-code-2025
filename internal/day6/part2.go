package day6

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/JannDeterling/advent-of-code-2025/internal/util"
)

func RunPart2() {
	input := util.ReadInput("./input/day6.txt")
	tasks := parseColumnBasedMathTask(input)
	result :=0
	for _,task := range tasks {
		result += task.Result
	}
	fmt.Println(result)
}

type MathTaskExtended struct {
	Values   [][]rune
	Operator string
	MaxNumberSize int
	Result   int
}

func parseColumnBasedMathTask(input []string) []MathTask {
	lenght := len(input)
	tasksExtended := []MathTaskExtended{}
	operatorLineRegxp := regexp.MustCompile(`([\+\*]+\s*)\s`)
	matches := operatorLineRegxp.FindAllStringSubmatch(input[lenght-1], -1)
	for _, match := range matches {
		maxSize := len(match[1])
		operator := strings.Trim(match[1], " ")
		tasksExtended = append(tasksExtended, MathTaskExtended{
			MaxNumberSize: maxSize,
			Operator: operator,
			Values: make([][]rune, maxSize),
			Result: 0,
		})
	}
	
	for y, line := range input {
		if y == lenght -1 {
			continue
		} else {
			prevEnd := 0
			for i, task := range tasksExtended {
				if prevEnd + task.MaxNumberSize < len(line)-1 {
					columns := line[prevEnd:prevEnd+task.MaxNumberSize]
					for c, rune := range columns {
						tasksExtended[i].Values[c] = append(tasksExtended[i].Values[c], rune)
					}
				} else {
					columns := line[prevEnd:]
					for c, rune := range columns {
						tasksExtended[i].Values[c] = append(tasksExtended[i].Values[c], rune)
					}
				}
				prevEnd += task.MaxNumberSize +1
			}
		}
	}

	tasks := []MathTask{}
	for _,exTask := range tasksExtended {
		convValues := []int{}
		result := 0 
		for i,value := range exTask.Values {
			convValue, err := strconv.Atoi(strings.Trim(string(value), " "))
			if err != nil {
				panic(err)
			}
			convValues = append(convValues, convValue)
			if i == 0 {
					result = convValue
				} else {
					switch exTask.Operator {
					case "+":
						result += convValue
					case "*":
						result *= convValue
					default:
						panic(fmt.Errorf("unsurported operator %s", exTask.Operator))
					}
				}
		}

		tasks = append(tasks, MathTask{
			Values: convValues,
			Operator: exTask.Operator,
			Result: result,
		})
	}

	return tasks
}
