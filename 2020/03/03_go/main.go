package main

import (
	"bufio"
	"fmt"
	"os"
)

func getLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	inputs := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		inputs = append(inputs, line)
	}

	return inputs
}

func numOfTreesP1(inputs []string) int {
	lineSize := len(inputs[0])
	numLines := len(inputs)
	currentRow := 0
	currentCol := 0
	numOfTrees := 0
	for {
		currentCol += 3
		currentRow += 1
		if currentRow >= numLines {
			break
		}
		if currentCol >= lineSize {
			currentCol = currentCol - lineSize
		}
		if string(inputs[currentRow][currentCol]) == "#" {
			numOfTrees++
		}
	}
	return numOfTrees
}

func main() {
	inputs := getLines("../input.txt")
	numTreesP1 := numOfTreesP1(inputs)
	fmt.Println("Day 03 - Part One - Num Of Trees:", numTreesP1)
}
