package main

import (
	"fmt"
	"strings"

	"github.com/gstoaldo/advent-of-code-2024/utils"
)

func parse(f string) (grid [][]string, i0, j0 int, dir0 string) {
	for _, line := range utils.ReadLines(f) {
		grid = append(grid, strings.Split(line, ""))
	}

	for i, row := range grid {
		for j := range row {
			if grid[i][j] != "." && grid[i][j] != "#" {
				i0, j0, dir0 = i, j, grid[i][j]
				grid[i][j] = "." // set initial guard position as a free space
			}
		}
	}

	return grid, i0, j0, dir0
}

var directions = map[string][]int{
	"^": {-1, 0},
	">": {0, 1},
	"v": {1, 0},
	"<": {0, -1},
}

type guard struct {
	i   int
	j   int
	dir []int
}

func (g *guard) step(grid [][]string) bool {
	ni, nj := g.i+g.dir[0], g.j+g.dir[1]

	if !(ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[0])) {
		// out of bounds
		return false
	}

	if grid[ni][nj] == "#" {
		g.dir = []int{g.dir[1], -g.dir[0]} // rotated vector = {j,−i}
	}

	if grid[ni][nj] == "." {
		g.i, g.j = ni, nj
	}

	return true
}

func p1(grid [][]string, i0, j0 int, dir0 string) int {
	g := guard{i0, j0, directions[dir0]}

	type pos struct{ i, j int }
	visited := map[pos]bool{
		{g.i, g.j}: true, // set starting position as visited
	}

	for g.step(grid) {
		visited[pos{g.i, g.j}] = true
	}

	return len(visited)
}

func p2(grid [][]string, i0, j0 int, dir0 string) (result int) {
	type pos struct{ i, j, di, dj int }

	for i := range grid {
		for j := range grid[i] {
			if i == i0 && j == j0 {
				continue
			}

			if grid[i][j] == "#" {
				continue
			}

			grid[i][j] = "#" // new obstacle

			g := guard{i0, j0, directions[dir0]}

			visited := map[pos]bool{
				{g.i, g.j, g.dir[0], g.dir[1]}: true,
			}

			for g.step(grid) {
				key := pos{g.i, g.j, g.dir[0], g.dir[1]}

				if visited[key] {
					// if the guard have been in the same position facing the same direction,
					// then he is in a loop
					result++
					break
				}

				visited[key] = true
			}

			grid[i][j] = "." // reset grid
		}
	}

	return result
}

func main() {
	fmt.Println(p1(parse(utils.Filepath())))
	fmt.Println(p2(parse(utils.Filepath())))
}
