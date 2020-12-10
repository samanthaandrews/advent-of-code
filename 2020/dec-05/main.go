package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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
	var seats []int
	highestSeat := 0
	for _, line := range split {
		if len(line) != 7+3 {
			fmt.Printf("Invalid boarding pass length %s", line)
		}
		col := binarySearch(line[:7], 'F', 'B', 127)
		row := binarySearch(line[7:], 'L', 'R', 7)
		seatID := col*8 + row
		seats = append(seats, seatID)
		if seatID > highestSeat {
			highestSeat = seatID
		}
	}
	sort.Ints(seats)
	prev := seats[0]
	for _, id := range seats[1:] {
		if prev == id-1 {
			prev++
		} else {
			fmt.Println("Part two", prev+1)
			break
		}
	}

	fmt.Println("Part one", highestSeat)
}

func binarySearch(line string, up byte, down byte, top int) int {
	min := 0
	max := top
	for i := range line {
		half := (max + min) / 2
		if line[i] == up {
			max = half
		} else {
			min = half + 1
		}
	}
	return min
}
