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
	children, _ := strconv.Atoi(parts[0])
	metadata, _ := strconv.Atoi(parts[1])
	parts = parts[2:]
	fmt.Println("PN", children, metadata, len(parts))
	for c := 0; c < children; c++ {
		childSum, childParts := processNode(parts)
		sum += childSum
		parts = childParts
	}

	for m := 0; m < metadata; m++ {
		metadataValue, _ := strconv.Atoi(parts[0])
		sum += metadataValue
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
