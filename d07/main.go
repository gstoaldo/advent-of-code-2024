package main

import (
	"fmt"
	"regexp"

	"github.com/gstoaldo/advent-of-code-2024/utils"
)

func parse(f string) (equations [][]int) {
	re := regexp.MustCompile(`\d+`)

	for _, line := range utils.ReadLines(f) {
		equation := []int{}
		for _, v := range re.FindAllString(line, -1) {
			equation = append(equation, utils.ToInt(v))
		}

		equations = append(equations, equation)
	}

	return equations
}

func testEquation(equation []int) bool {
	target, numbers := equation[0], equation[1:]

	combinations := []int{numbers[0]}

	for _, n := range numbers[1:] {
		newCombinations := []int{}

		for _, c := range combinations {
			newCombinations = append(newCombinations, c+n)
			newCombinations = append(newCombinations, c*n)
		}
		combinations = newCombinations
	}

	for _, c := range combinations {
		if c == target {
			return true
		}
	}

	return false
}

func p1(equations [][]int) (result int) {
	for _, eq := range equations {
		if testEquation(eq) {
			result += eq[0]
		}
	}

	return result
}

func main() {
	fmt.Println(p1(parse(utils.Filepath())))
}
