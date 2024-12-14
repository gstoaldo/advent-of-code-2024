package main

import (
	"fmt"

	"github.com/gstoaldo/advent-of-code-2024/utils"
)

func parse(f string) (diskmap []int) {
	for _, s := range utils.ReadFile(f) {
		diskmap = append(diskmap, utils.ToInt(string(s)))
	}

	return diskmap
}

func fileID(index int) int {
	return index / 2
}

func p1(diskmap []int) int {
	// left and right pointers
	l, r := 0, len(diskmap)-1

	expanded := []int{}

	for l <= r {
		if diskmap[l] == 0 {
			l++
			continue
		}

		if diskmap[r] == 0 || r%2 != 0 {
			r--
			continue
		}

		// even positions in diskmap are files, odd are free space

		// left pointer is in a file
		if l%2 == 0 {
			expanded = append(expanded, fileID(l))
			diskmap[l]--
			continue
		}

		// left pointer is in a free space
		expanded = append(expanded, fileID(r))
		diskmap[l]--
		diskmap[r]--
	}

	checksum := 0

	for i, v := range expanded {
		checksum += i * v
	}

	return checksum
}

func main() {
	fmt.Println(p1(parse(utils.Filepath())))
}
