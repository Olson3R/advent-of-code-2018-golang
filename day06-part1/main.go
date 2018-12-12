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

func findClosest(coordinate Coordinate, locations []Location) string {
	name := STAR
	distance := -1
	for _, location := range locations {
		locationDistance := Abs(coordinate.x - location.coordinate.x)
		locationDistance += Abs(coordinate.y - location.coordinate.y)
		if distance == -1 || distance > locationDistance {
			distance = locationDistance
			name = location.name
		} else if distance == locationDistance {
			name = STAR
		}
	}
	return name
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

	grid := make(map[Coordinate]string)
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
		grid[location.coordinate] = location.name
		if name != 'z' {
			name += 1
		} else {
			name = 'A'
		}
	}


	// find the closest location for each coordinate
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			coordinate := Coordinate{x, y}
			_, ok := grid[coordinate]
			if !ok {
				grid[coordinate] = findClosest(coordinate, locations)
			}
		}
	}

	// find infinite locations to ignore by getting locations touching the edges
	infiniteLocations := make(map[string]bool)
	for x := minX; x <= maxX; x++ {
		// top row
		name := grid[Coordinate{x, minY}]
		if name != STAR {
			infiniteLocations[name] = true
		}

		// bottom row
		name = grid[Coordinate{x, maxY}]
		if name != STAR {
			infiniteLocations[name] = true
		}
	}

	for y := minY; y <= maxY; y++ {
		// left column
		name := grid[Coordinate{y, minX}]
		if name != STAR {
			infiniteLocations[name] = true
		}

		// right column
		name = grid[Coordinate{y, maxX}]
		if name != STAR {
			infiniteLocations[name] = true
		}
	}

	// compute sizes of each location
	sizes := make(map[string]int)
	for _, location := range locations {
		sizes[location.name] = 0
	}

	for _, name := range grid {
		if name != STAR {
			sizes[name] += 1
		}
	}

	// find largest location
	largest := 0
	for name, size := range sizes {
		_, ok := infiniteLocations[name]
		if !ok && largest < size {
			largest = size
			fmt.Println("NEW LARGEST", name, size)
		}
	}

	fmt.Println("Locations", locations)
	fmt.Println("Grid", len(grid))
	fmt.Println("Infinite locations", infiniteLocations)
	fmt.Println("LARGEST", largest)
	fmt.Println("Done")
}
