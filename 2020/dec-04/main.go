package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var hexColor = regexp.MustCompile("^#[0-9a-f]{6}$")
var pidDigits = regexp.MustCompile("^[0-9]{9}$")

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n\n")
	passports := []map[string]string{
		make(map[string]string),
	}
	passports = passports[:len(passports)-1]
	for i, s := range split {
		s = strings.ReplaceAll(s, "\n", " ")
		passports = append(passports, make(map[string]string))
		for _, entry := range strings.Split(s, " ") {
			keyValueParts := strings.Split(entry, ":")
			if len(keyValueParts) != 2 {
				fmt.Printf("Failed to parse %s", entry)
				return
			}
			passports[i][keyValueParts[0]] = keyValueParts[1]
		}
		if err != nil {
			fmt.Printf("Failed to parse %s\n", split)
		}
	}
	partOneValidPassports := 0
outer:
	for _, passport := range passports {
		for _, k := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
			if passport[k] == "" {
				continue outer
			}
		}
		partOneValidPassports++
	}
	fmt.Println("Part one", partOneValidPassports)
	partTwoValidPassports := 0
	for _, passport := range passports {
		byr, err := strconv.Atoi(passport["byr"])
		if err != nil || byr < 1920 || byr > 2002 {
			continue
		}
		iyr, err := strconv.Atoi(passport["iyr"])
		if err != nil || iyr < 2010 || iyr > 2020 {
			continue
		}
		eyr, err := strconv.Atoi(passport["eyr"])
		if err != nil || eyr < 2020 || eyr > 2030 {
			continue
		}
		if len(passport["hgt"]) <= 3 {
			continue
		}
		hgt, err := strconv.Atoi(passport["hgt"][:len(passport["hgt"])-2])
		units := passport["hgt"][len(passport["hgt"])-2:]
		if err != nil {
			continue
		}
		switch units {
		case "in":
			if hgt < 59 || hgt > 76 {
				continue
			}
		case "cm":
			if hgt < 150 || hgt > 193 {
				continue
			}
		default:
			continue
		}
		if !hexColor.MatchString(passport["hcl"]) {
			continue
		}
		switch passport["ecl"] {
		case "amb":
		case "blu":
		case "brn":
		case "gry":
		case "grn":
		case "hzl":
		case "oth":
		default:
			continue
		}
		if !pidDigits.MatchString(passport["pid"]) {
			continue
		}
		partTwoValidPassports++
	}
	fmt.Println("Part two", partTwoValidPassports)
}
