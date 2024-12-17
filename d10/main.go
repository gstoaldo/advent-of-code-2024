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
			if s == "." {
				row = append(row, -1)
			} else {
				row = append(row, utils.ToInt(s))
			}
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

func copyMap(original map[loc]bool) map[loc]bool {
	cp := map[loc]bool{}
	for k, v := range original {
		cp[k] = v
	}

	return cp
}

func p2(M [][]int) int {
	totalRating := 0

	type trail struct {
		curr    loc
		visited map[loc]bool
	}

	for i := range M {
		for j := range M[i] {
			// not a trailhead
			if M[i][j] != 0 {
				continue
			}

			stack := []trail{
				{
					curr:    loc{i, j},
					visited: map[loc]bool{},
				},
			}

			for len(stack) > 0 {
				currTrail := stack[len(stack)-1]
				curr := currTrail.curr
				stack = stack[:len(stack)-1]

				currTrail.visited[curr] = true

				for _, n := range neighbors(curr) {
					if inside(M, n) && M[n.i][n.j]-M[curr.i][curr.j] == 1 && !currTrail.visited[n] {
						stack = append(stack, trail{curr: n, visited: copyMap(currTrail.visited)})

						if M[n.i][n.j] == 9 {
							totalRating++
						}
					}
				}
			}
		}
	}

	return totalRating
}

func main() {
	fmt.Println(p1(parse(utils.Filepath())))
	fmt.Println(p2(parse(utils.Filepath())))
}
