package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	partOneTally := 0
	partTwoTally := 0

	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	stringInput := string(input)
	stringSet := strings.Split(stringInput, "\n")

	for _, stringItem := range stringSet {

		firstSet := strings.Split(stringItem, ": ")
		remainderString, password := firstSet[0], firstSet[1]

		secondSet := strings.Split(remainderString, " ")
		numberRange, letter := secondSet[0], secondSet[1]

		numbers := strings.Split(numberRange, "-")
		startStr, stopStr := numbers[0], numbers[1]

		start, _ := strconv.Atoi(string(startStr))
		stop, _ := strconv.Atoi(string(stopStr))

		if checkCountInRange(password, letter, start, stop) {
			partOneTally++
		}
		if checkIndices(password, letter, start-1, stop-1) {
			partTwoTally++
		}
	}
	fmt.Println("Part one", partOneTally)
	fmt.Println("Part one", partTwoTally)
}

func checkCountInRange(password string, letter string, start int, stop int) bool {
	return strings.Count(password, letter) >= start && strings.Count(password, letter) <= stop
}

func checkIndices(password string, letter string, start int, stop int) bool {
	// make sure that the values at the start and stop positions are not equal
	return strings.Index(password, letter) == start || strings.Index(password, letter) == stop
}
