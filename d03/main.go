package main

import (
	"fmt"
	"regexp"

	"github.com/gstoaldo/advent-of-code-2024/utils"
)

func p1(text string) int {
	result := 0

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	for _, match := range re.FindAllStringSubmatch(text, -1) {
		result += utils.ToInt(match[1]) * utils.ToInt(match[2])
	}

	return result
}

func main() {
	input := utils.ReadFile(utils.Filepath())
	fmt.Println(p1(input))
}
