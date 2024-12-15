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

func isFile(index int) bool {
	return index%2 == 0
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

func p2(diskmap []int) int {
	type file struct {
		id   int
		i0   int
		size int
	}

	type space struct {
		i0   int
		size int
	}

	// sorted by decreasing id
	files := []file{}
	for i := len(diskmap) - 1; i >= 0; i-- {
		if !isFile(i) {
			continue
		}

		i0 := 0
		for j := 0; j < i; j++ {
			i0 += diskmap[j]
		}

		files = append(files, file{
			id:   fileID(i),
			i0:   i0,
			size: diskmap[i],
		})
	}

	// sorted from left to right
	spaces := []space{}
	for i := 0; i < len(diskmap); i++ {
		if isFile(i) {
			continue
		}

		i0 := 0
		for j := 0; j < i; j++ {
			i0 += diskmap[j]
		}

		spaces = append(spaces, space{
			i0:   i0,
			size: diskmap[i],
		})
	}

	for f := range files {
		for s := range spaces {
			if files[f].size <= spaces[s].size && files[f].i0 > spaces[s].i0 {
				files[f].i0 = spaces[s].i0
				spaces[s].size -= files[f].size
				spaces[s].i0 += files[f].size
				break
			}
		}

	}

	checksum := 0
	for _, f := range files {
		for i := 0; i < f.size; i++ {
			checksum += f.id * (f.i0 + i)
		}
	}

	return checksum
}

func main() {
	fmt.Println(p1(parse(utils.Filepath())))
	fmt.Println(p2(parse(utils.Filepath())))
}
