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

func match(mask string, currMask string) bool {
	for i, c := range mask {
		if c == '*' {
			continue
		}

		if string(currMask[i]) != string(c) {
			return false
		}
	}

	return true
}

func p2(lines []string) (result int) {
	masks := []string{
		"M*M:*A*:S*S",
		"M*S:*A*:M*S",
		"S*M:*A*:S*M",
		"S*S:*A*:M*M",
	}

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			maxI := i + 2
			maxJ := j + 2

			if !(maxI >= 0 && maxI < len(lines) && maxJ >= 0 && maxJ < len(lines[0])) {
				// out of bounds
				continue
			}

			currentMask := fmt.Sprintf("%s:%s:%s", lines[i][j:j+3], lines[i+1][j:j+3], lines[i+2][j:j+3])

			for _, mask := range masks {
				if match(mask, currentMask) {
					result++
					break
				}
			}
		}
	}

	return result
}

func main() {
	lines := utils.ReadLines(utils.Filepath())
	fmt.Println(p1(lines))
	fmt.Println(p2(lines))
}
