package utils

import (
	"fmt"
	"time"
)

func PrintSolution(p1, p2 interface{}) {
	fmt.Printf("Part 1: %v\nPart 2: %v\n", p1, p2)
}

func Timer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("total time: %v\n", time.Since(start))
	}
}
