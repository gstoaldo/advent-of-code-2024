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

func p(M [][]string, anthenas map[string][]loc, harmonics bool) int {
	antinodes := map[loc]bool{}

	for _, locs := range anthenas {
		// pairwise
		for a := 0; a < len(locs)-1; a++ {
			for b := a + 1; b < len(locs); b++ {
				di := locs[b].i - locs[a].i
				dj := locs[b].j - locs[a].j

				i := -1
				if !harmonics {
					i = 0
				}

				for {
					antinodeA := loc{locs[a].i + di*(2+i), locs[a].j + dj*(2+i)}
					if !inBounds(M, antinodeA) {
						break
					}
					antinodes[antinodeA] = true
					i++

					if !harmonics {
						break
					}
				}

				i = -1

				if !harmonics {
					i = 0
				}

				for {
					antinodeB := loc{locs[b].i - di*(2+i), locs[b].j - dj*(2+i)}
					if !inBounds(M, antinodeB) {
						break
					}
					antinodes[antinodeB] = true
					i++

					if !harmonics {
						break
					}
				}
			}
		}
	}

	return len(antinodes)
}

func main() {
	M, anthenas := parse(utils.Filepath())
	fmt.Println(p(M, anthenas, false))
	fmt.Println(p(M, anthenas, true))

	// 358 too low
}
