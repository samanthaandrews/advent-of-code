package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Entry []string

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n\n")
	var customsForms []Entry
	uniqueYes := 0
	for _, s := range split {
		s = strings.ReplaceAll(s, "\n", "")
		entries := strings.Split(s, "")
		customsForms = append(customsForms, entries)
		keys := make(map[string]bool)
		unique := []string{}
		for _, entry := range entries {
			if _, value := keys[entry]; !value {
				keys[entry] = true
				unique = append(unique, entry)
			}
		}
		uniqueYes = uniqueYes + len(unique)
	}
	fmt.Println("Part one", uniqueYes)
}
