package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	inputs := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.Trim(line, "\n")
		num, err := strconv.Atoi(trimmedLine)
		if err != nil {
			panic(err)
		}
		inputs = append(inputs, num)
	}

	for _, n := range inputs {
		for _, n2 := range inputs {
			if n+n2 == 2020 {
				fmt.Println("Part One", n*n2)
			}
		}
	}

	for _, n := range inputs {
		for _, n2 := range inputs {
			for _, n3 := range inputs {
				if n+n2+n3 == 2020 {
					fmt.Println("Part Two", n*n2*n3)
				}
			}
		}
	}
}
