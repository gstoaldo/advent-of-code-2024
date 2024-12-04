package main

import (
	"fmt"
	"regexp"

	"github.com/gstoaldo/advent-of-code-2024/utils"
)

func mult(text string, ignoreDont bool) int {
	result := 0

	reMult := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	reDoMult := regexp.MustCompile(`do\(\)`)
	reDontMult := regexp.MustCompile(`don't\(\)`)

	multIndexes := reMult.FindAllStringIndex(text, -1)

	doMultIndexes := append(
		[][]int{{0, 0}}, // "do" is enabled by default at the beggining
		reDoMult.FindAllStringIndex(text, -1)...,
	)
	dontMultIndexes := reDontMult.FindAllStringIndex(text, -1)

	for i, match := range reMult.FindAllStringSubmatch(text, -1) {
		matchIndex := multIndexes[i][0]

		distDo := 999999
		for _, v := range doMultIndexes {
			dist := matchIndex - v[1]

			// ignore if do/dont is after mult
			if dist < 0 {
				break
			}

			distDo = utils.Min(distDo, dist)
		}

		distDont := 999999
		for _, v := range dontMultIndexes {
			dist := matchIndex - v[1]

			if dist < 0 {
				break
			}

			distDont = utils.Min(distDont, dist)
		}

		if (distDo < distDont) || ignoreDont {
			result += utils.ToInt(match[1]) * utils.ToInt(match[2])
		}
	}

	return result
}

func main() {
	input := utils.ReadFile(utils.Filepath())
	fmt.Println(mult(input, true))
	fmt.Println(mult(input, false))
}
