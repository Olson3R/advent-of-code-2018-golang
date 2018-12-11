package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Coordinate struct {
	id int
	x  int
	y  int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func strip(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func parseRow(row string) (int, int) {
	parts := strings.FieldsFunc(row, func(r rune) bool {
		switch r {
		case ',':
			return true
		}
		return false
	})
	x, err := strconv.Atoi(parts[0])
	check(err)
	y, err := strconv.Atoi(parts[1])
	check(err)
	return x, y
}

func main() {
	file, err := os.Open("input/d06-input.txt")
	check(err)
	defer file.Close()

	var coordinates []Coordinate
	i := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strip(scanner.Text())
		x, y := parseRow(row)
		coordinates = append(coordinates, Coordinate{i, x, y})
		i += 1
	}

	fmt.Println("Coordinates", coordinates)
}
