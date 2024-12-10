package main

import (
	"fmt"
	"strings"

	"github.com/gstoaldo/advent-of-code-2024/utils"
)

type rule struct{ x, y int }

func parse(f string) (rules []rule, updates [][]int) {
	text := utils.ReadFile(f)

	parts := strings.Split(text, "\n\n")

	for _, line := range strings.Split(parts[0], "\n") {
		s := strings.Split(line, "|")

		rules = append(rules, rule{utils.ToInt(s[0]), utils.ToInt(s[1])})
	}

	for _, line := range strings.Split(parts[1], "\n") {
		pages := []int{}
		for _, s := range strings.Split(line, ",") {
			pages = append(pages, utils.ToInt(s))
		}

		updates = append(updates, pages)
	}

	return rules, updates
}

func valid(r rule, pages []int, i int) (bool, int) {
	target := pages[i]

	left := pages[:i]
	right := pages[i+1:]

	if target == r.x {
		for j, v := range left {
			if v == r.y {
				return false, j
			}
		}
	}

	if target == r.y {
		for j, v := range right {
			if v == r.x {
				return false, j + i + 1
			}
		}
	}

	return true, 0
}

func updateIsValid(rules []rule, pages []int) bool {
	for _, r := range rules {
		for i := range pages {
			isValid, _ := valid(r, pages, i)

			if !isValid {
				return false
			}
		}
	}

	return true
}

func sort(rules []rule, pages []int) {
	for target := 0; target < len(pages); target++ {
		rIndex := 0
		for rIndex < len(rules) {
			r := rules[rIndex]

			isValid, pivot := valid(r, pages, target)

			if !isValid {
				temp := pages[target]

				pages[target] = pages[pivot]
				pages[pivot] = temp
				rIndex = 0 // go back to first rule
			}

			if isValid {
				rIndex++
			}
		}
	}
}

func p1(rules []rule, updates [][]int) (result int) {
	for _, pages := range updates {
		if updateIsValid(rules, pages) {
			result += pages[(len(pages)-1)/2]
		}
	}

	return result
}

func p2(rules []rule, updates [][]int) (result int) {
	for _, pages := range updates {
		if !updateIsValid(rules, pages) {
			sort(rules, pages)
			result += pages[(len(pages)-1)/2]
		}
	}

	return result
}

func main() {
	rules, updates := parse(utils.Filepath())

	fmt.Println(p1(rules, updates))
	fmt.Println(p2(rules, updates))
}
