package main

import (
	"fmt"
	"strings"

	"github.com/gstoaldo/advent-of-code-2024/utils"
)

func parse(f string) (M [][]int) {
	for _, line := range utils.ReadLines(f) {
		row := []int{}
		for _, s := range strings.Split(line, "") {
			row = append(row, utils.ToInt(s))
		}
		M = append(M, row)
	}

	return M
}

type loc struct{ i, j int }

func neighbors(l loc) (result []loc) {
	for _, delta := range [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		result = append(result, loc{l.i + delta[0], l.j + delta[1]})
	}

	return result
}

func inside(M [][]int, l loc) bool {
	return l.i >= 0 && l.i < len(M) && l.j >= 0 && l.j < len(M[0])
}

func p1(M [][]int) int {
	totalScore := 0

	for i := range M {
		for j := range M[i] {
			visited := map[loc]bool{}

			// not a trailhead
			if M[i][j] != 0 {
				continue
			}

			stack := []loc{
				{i, j},
			}

			for len(stack) > 0 {
				curr := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				visited[curr] = true

				for _, n := range neighbors(curr) {
					if inside(M, n) && M[n.i][n.j]-M[curr.i][curr.j] == 1 && !visited[n] {
						stack = append(stack, n)

						if M[n.i][n.j] == 9 {
							totalScore++
						}
					}
				}
			}
		}
	}

	return totalScore
}

func main() {
	fmt.Println(p1(parse(utils.Filepath())))
}
