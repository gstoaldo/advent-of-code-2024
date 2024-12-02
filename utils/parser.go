package utils

import (
	"os"
	"strconv"
	"strings"
)

func Filepath() string {
	filepath := "input.txt"

	if len(os.Args) == 2 {
		filepath = os.Args[1]
	}

	return filepath
}

func ReadFile(filepath string) string {
	file, err := os.ReadFile(filepath)
	if err != nil {
		panic("error reading file")
	}

	return string(file)
}

func ReadLines(filepath string) []string {
	file := ReadFile(filepath)

	return strings.Split(string(file), "\n")
}

func ToInt(s string) int {
	v, err := strconv.Atoi(s)

	if err != nil {
		panic("could not convert string to int")
	}

	return v
}
