package main

import (
	"fmt"
	"strings"

	"github.com/gstoaldo/advent-of-code-2024/utils"
)

type loc struct{ i, j int }

func parse(f string) ([][]string, map[string][]loc) {
	M := [][]string{}
	anthenas := map[string][]loc{}

	for _, line := range utils.ReadLines(f) {
		M = append(M, strings.Split(line, ""))
	}

	for i := range M {
		for j := range M[i] {
			if M[i][j] == "." {
				continue
			}

			anthenas[M[i][j]] = append(anthenas[M[i][j]], loc{i, j})
		}
	}

	return M, anthenas
}

func inBounds(M [][]string, l loc) bool {
	return l.i >= 0 && l.i < len(M) && l.j >= 0 && l.j < len(M[0])
}

func p1(M [][]string, anthenas map[string][]loc) int {
	antinodes := map[loc]bool{}

	for _, locs := range anthenas {
		// pairwise
		for a := 0; a < len(locs)-1; a++ {
			for b := a + 1; b < len(locs); b++ {
				di := locs[b].i - locs[a].i
				dj := locs[b].j - locs[a].j

				antinodeA := loc{locs[a].i + 2*di, locs[a].j + 2*dj}
				if inBounds(M, antinodeA) {
					antinodes[antinodeA] = true
				}

				antinodeB := loc{locs[b].i - 2*di, locs[b].j - 2*dj}
				if inBounds(M, antinodeB) {
					antinodes[antinodeB] = true
				}
			}
		}
	}

	return len(antinodes)
}

func main() {
	fmt.Println(p1(parse(utils.Filepath())))
}
