package main

import (
	"fmt"
	"strings"

	"github.com/gstoaldo/advent-of-code-2024/utils"
)

func parse(filepath string) (reports [][]int) {
	for _, line := range utils.ReadLines(filepath) {
		report := []int{}

		for _, v := range strings.Split(line, " ") {
			report = append(report, utils.ToInt(v))
		}

		reports = append(reports, report)
	}

	return reports
}

func p1(reports [][]int) int {
	result := 0

	for _, report := range reports {
		if report[1]-report[0] == 0 {
			continue
		}

		// sign = +1 for increasing, -1 for decreasing
		sign := (report[1] - report[0]) / utils.Abs(report[1]-report[0])

		safe := true

		for i := 1; i < len(report); i++ {
			value := (report[i] - report[i-1]) * sign

			if value <= 0 || value > 3 {
				safe = false
				break
			}
		}

		if safe {
			result++
		}
	}

	return result
}

func main() {
	reports := parse(utils.Filepath())

	fmt.Println(p1(reports))
}
