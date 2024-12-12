package main

import (
	"fmt"
	"regexp"
	"strconv"

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

func testEquation(equation []int, operations []func(a, b int) int) bool {
	target, numbers := equation[0], equation[1:]

	combinations := []int{numbers[0]}

	for _, n := range numbers[1:] {
		newCombinations := []int{}

		for _, c := range combinations {
			for _, op := range operations {
				newCombinations = append(newCombinations, op(c, n))
			}
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
	operations := []func(a, b int) int{
		func(a, b int) int { return a + b },
		func(a, b int) int { return a * b },
	}

	for _, eq := range equations {
		if testEquation(eq, operations) {
			result += eq[0]
		}
	}

	return result
}

func p2(equations [][]int) (result int) {
	operations := []func(a, b int) int{
		func(a, b int) int { return a + b },
		func(a, b int) int { return a * b },
		func(a, b int) int { return utils.ToInt(strconv.Itoa(a) + strconv.Itoa(b)) },
	}

	for _, eq := range equations {
		if testEquation(eq, operations) {
			result += eq[0]
		}
	}

	return result
}

func main() {
	fmt.Println(p1(parse(utils.Filepath())))
	fmt.Println(p2(parse(utils.Filepath())))
}
