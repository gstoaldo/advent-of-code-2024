package main

import (
	"fmt"
	"regexp"
	"slices"

	"github.com/gstoaldo/advent-of-code-2024/utils"
)

func parse(filepath string) ([]int, []int) {
	l1, l2 := []int{}, []int{}

	re := regexp.MustCompile(`\d+`)

	for _, line := range utils.ReadLines(filepath) {
		matches := re.FindAllString(line, -1)

		l1 = append(l1, utils.ToInt(matches[0]))
		l2 = append(l2, utils.ToInt(matches[1]))
	}

	return l1, l2
}

func p1(l1, l2 []int) int {
	slices.Sort(l1)
	slices.Sort(l2)

	result := 0

	for i := range l1 {
		result += utils.Abs(l1[i] - l2[i])
	}

	return result
}

func p2(l1, l2 []int) int {
	result := 0

	l2Count := map[int]int{}

	for _, n := range l2 {
		l2Count[n] += 1
	}

	for _, n := range l1 {
		result += n * l2Count[n]
	}

	return result
}

func main() {
	l1, l2 := parse(utils.Filepath())

	fmt.Println(p1(l1, l2))
	fmt.Println(p2(l1, l2))
}
