package day6

import (
	"fmt"
	"regexp"
	"strconv"
	"github.com/JannDeterling/advent-of-code-2025/internal/util"
)

type MathTask struct {
	Values   []int
	Operator string
	Result   int
}

func RunPart1() {
	input := util.ReadInput("./input/day6.txt")
	tasks := parseMathTask(input)
	result :=0
	for _,task := range tasks {
		result += task.Result
	}
	fmt.Println(result)
}

func parseMathTask(input []string) []MathTask {
	lenght := len(input)
	operator := ""
	tasks := []MathTask{}
	re := regexp.MustCompile(`\s+`)
	for y := lenght - 1; y >= 0; y-- {
		split := re.Split(input[y], -1)
		for x, xInput := range split {
			if y == lenght-1 {
				tasks = append(tasks, MathTask{
					Operator: xInput,
					Values:   []int{},
					Result:   0,
				})
			} else {
				value, err := strconv.Atoi(xInput)
				if err != nil {
					panic(err)
				}
				tasks[x].Values = append(tasks[x].Values, value)
				if y == lenght-2 {
					tasks[x].Result = value
				} else {
					switch tasks[x].Operator {
					case "+":
						tasks[x].Result += value
					case "*":
						tasks[x].Result *= value
					default:
						panic(fmt.Errorf("unsurported operator %s", operator))
					}
				}
			}
		}
	}
	return tasks
}
