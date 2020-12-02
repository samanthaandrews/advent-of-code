package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	stringInput := string(input)

	stringSet := strings.Split(stringInput, "\n")

	set := []int{}

	for _, stringItem := range stringSet {
		number, _ := strconv.Atoi(string(stringItem))
		set = append(set, number)
	}
	fmt.Println("Part 1", partOne(set))
	fmt.Println("Part 2", partTwo(set))
}

func partOne(array []int) int {
	for _, i := range array {
		for _, j := range array {
			if i+j == 2020 {
				return i * j
			}
		}
	}
	return 0
}

func partTwo(array []int) int {
	for _, i := range array {
		for _, j := range array {
			for _, k := range array {
				if i+j+k == 2020 {
					return i * j * k
				}
			}
		}
	}
	return 0
}
