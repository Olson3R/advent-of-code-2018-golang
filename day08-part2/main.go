package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseRow(row string) int {
	parts := strings.Split(row, " ")
	sum, _ := processNode(parts)
	return sum
}

func processNode(parts []string) (int, []string) {
	sum := 0
	numChildren, _ := strconv.Atoi(parts[0])
	numMetadata, _ := strconv.Atoi(parts[1])
	parts = parts[2:]
	children := make([]int, numChildren)
	fmt.Println("PN", numChildren, numMetadata, len(parts))
	for c := 0; c < numChildren; c++ {
		childSum, childParts := processNode(parts)
		children[c] = childSum
		parts = childParts
	}

	for m := 0; m < numMetadata; m++ {
		value, _ := strconv.Atoi(parts[0])
		if numChildren > 0 {
			if value <= numChildren {
				sum += children[value-1]
			}
		} else {
			sum += value
		}
		parts = parts[1:]
	}

	return sum, parts
}

func main() {
	file, err := os.Open("input/d08-input.txt")
	check(err)
	defer file.Close()

	// parse parent/child steps
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		sum := parseRow(row)
		fmt.Println("SUM", sum)
	}

	fmt.Println("Done")
}
