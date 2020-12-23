package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getPassportStrs(filename string) []string {
	content, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	text := string(content)
	passportStrs := strings.Split(text, "\n\n")
	return passportStrs
}

func getValidPassports(passportStrs []string) int {
	validKeys := []string{
		"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid",
	}
	numValidPassports := 0
	for _, passport := range passportStrs {
		keyValuePairs := strings.Fields(passport)
		passportMap := make(map[string]string)
		for _, keyValuePair := range keyValuePairs {
			keyValueArr := strings.Split(keyValuePair, ":")
			passportMap[keyValueArr[0]] = keyValueArr[1]
		}
		isValidPassport := true
		for _, validKey := range validKeys {
			_, isInMap := passportMap[validKey]
			if !isInMap {
				isValidPassport = false
				break
			}
		}
		if isValidPassport {
			numValidPassports++
		}
	}
	return numValidPassports
}

func main() {
	passportStrs := getPassportStrs("../input.txt")
	numValidPassports := getValidPassports(passportStrs)
	fmt.Println("Day 04 - Part 1 - # of Valid Passports:", numValidPassports)
}
