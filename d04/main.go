package main

import (
	"fmt"

	"github.com/gstoaldo/advent-of-code-2024/utils"
)

func p1(lines []string) (result int) {
	directions := []struct{ di, dj int }{
		{0, 1},   // ->
		{0, -1},  // <-
		{1, 0},   // V
		{-1, 0},  // ^
		{1, 1},   // V ->
		{1, -1},  // V <-
		{-1, 1},  // ^ ->
		{-1, -1}, // ^ <-
	}

	// brute force: iterate over every element and check if there is a match in
	// any direction
	for _, dir := range directions {
		for i := 0; i < len(lines); i++ {
			for j := 0; j < len(lines[0]); j++ {
				maxI := i + 3*dir.di
				maxJ := j + 3*dir.dj

				if !(maxI >= 0 && maxI < len(lines) && maxJ >= 0 && maxJ < len(lines[0])) {
					// out of bounds
					continue
				}

				match := true
				for idx, c := range "XMAS" {
					if string(lines[i+idx*dir.di][j+idx*dir.dj]) != string(c) {
						match = false
					}
				}

				if match {
					result++
				}
			}
		}
	}

	return result
}

func main() {
	lines := utils.ReadLines(utils.Filepath())
	fmt.Println(p1(lines))
}
