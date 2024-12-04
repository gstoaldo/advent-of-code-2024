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

func permutations(report []int) [][]int {
	result := [][]int{}

	for i := range report {
		cp := append([]int{}, report...)
		result = append(result, append(cp[:i], cp[i+1:]...))
	}

	return result
}

func isSafe(report []int) bool {
	deltas := []int{}

	for i := 1; i < len(report); i++ {
		deltas = append(deltas, report[i]-report[i-1])
	}

	for _, v := range deltas {
		if utils.Abs(v) > 3 || v == 0 {
			return false
		}
	}

	// check if it is always increasing or decreasing
	for i := 1; i < len(deltas); i++ {
		if deltas[i]*deltas[i-1] <= 0 {
			return false
		}
	}

	return true
}

func p1(reports [][]int) int {
	result := 0

	for _, report := range reports {
		if isSafe(report) {
			result++
		}
	}

	return result
}

func p2(reports [][]int) int {
	result := 0

	for _, report := range reports {
		anyIsSafe := false

		// Brute force: generate every possible permutation by removing one
		// element from the slice and check if any permutation is safe.
		// The size of the reports in the input is pretty small (max 8 elements),
		// so it's not a big deal to use a brute force approach.
		for _, permutation := range permutations(report) {
			anyIsSafe = anyIsSafe || isSafe(permutation)
		}

		if anyIsSafe {
			result++
		}
	}

	return result
}

func main() {
	reports := parse(utils.Filepath())

	fmt.Println(p1(reports))
	fmt.Println(p2(reports))
}
