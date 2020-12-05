package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")
	// split = split[:len(split)-1]
	trees := make([][]bool, len(split))
	for i, s := range split {
		trees[i] = make([]bool, len(s))
		for j, c := range s {
			trees[i][j] = (c == '#')
		}
		if err != nil {
			fmt.Printf("Failed to parse %s\n", s)
		}
	}
	fmt.Println("Part one", checkSlope(1, 3, trees))
	fmt.Println("Part two", checkSlope(1, 3, trees)*checkSlope(1, 1, trees)*checkSlope(1, 5, trees)*checkSlope(1, 7, trees)*checkSlope(2, 1, trees))
}

func checkSlope(down, right int, trees [][]bool) int {
	hit := 0
	for time := 0; time*down < len(trees); time++ {
		traversed := time * down
		column := (time * right) % len(trees[traversed])
		if trees[traversed][column] {
			hit++
		}
	}
	return hit
}
