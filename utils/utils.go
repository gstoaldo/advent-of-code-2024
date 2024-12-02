package utils

import (
	"fmt"
	"time"
)

func Timer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("total time: %v\n", time.Since(start))
	}
}
