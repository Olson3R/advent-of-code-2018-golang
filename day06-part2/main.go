package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const STAR = "*"
const MAX_DISTANCE = 10000

type Location struct {
	name       string
	coordinate Coordinate
}

type Coordinate struct {
	x int
	y int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

func closeToAll(coordinate Coordinate, locations []Location) bool {
	distance := 0
	for _, location := range locations {
		distance += Abs(coordinate.x - location.coordinate.x)
		distance += Abs(coordinate.y - location.coordinate.y)
		if distance >= MAX_DISTANCE {
			return false
		}
	}
	return true
}

func findMinMax(v int, min int, max int) (int, int) {
	if min == -1 || min > v {
		min = v
	}
	if max == -1 || max < v {
		max = v
	}
	return min, max
}

func main() {
	file, err := os.Open("input/d06-input.txt")
	check(err)
	defer file.Close()

	var locations []Location
	name := 'a'

	minX := -1
	maxX := -1
	minY := -1
	maxY := -1

	// parse and save locations along with finding min/max X/Y
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strip(scanner.Text())
		x, y := parseRow(row)

		minX, maxX = findMinMax(x, minX, maxX)
		minY, maxY = findMinMax(y, minY, maxY)

		// save locations
		location := Location{string(name), Coordinate{x, y}}
		locations = append(locations, location)
		if name != 'z' {
			name += 1
		} else {
			name = 'A'
		}
	}

	size := 0
	// test each coordinate
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if closeToAll(Coordinate{x, y}, locations) {
				size += 1
			}
		}
	}

	fmt.Println("Locations", locations)
	fmt.Println("SIZE", size)
	fmt.Println("Done")
}
