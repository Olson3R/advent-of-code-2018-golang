package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"unicode"
)

type Coord struct {
	x int
	y int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func stripWhitespace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func parseRow(row string) (int, int, int, int) {
	parts := strings.FieldsFunc(row, func(r rune) bool {
		switch r {
		case ':':
			return true
		case ',':
			return true
		case '@':
			return true
		case 'x':
			return true
		}
		return false
	})
	x, err := strconv.Atoi(parts[1])
	check(err)
	y, err := strconv.Atoi(parts[2])
	check(err)
	width, err := strconv.Atoi(parts[3])
	check(err)
	height, err := strconv.Atoi(parts[4])
	check(err)
	return x, y, width, height
}

func main() {
	file, err := os.Open("d3-data.txt")
	check(err)
	defer file.Close()

	coords := make(map[Coord]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := stripWhitespace(scanner.Text())
		x, y, width, height := parseRow(row)
		fmt.Println("x: ", x, ", y: ", y, ", width: ", width, ", height: ", height)
		for i := x + 1; i <= x + width; i++ {
			for j := y + 1; j <= y + height; j++ {
				coord := Coord{i, j}
				count, ok := coords[coord]
				if ok {
					coords[coord] = count + 1
				} else {
					coords[coord] = 1
				}
			}
		}
	}

	overlapCount := 0
	for _, value := range coords {
		if value > 1 {
			fmt.Println("V: ", value)
			overlapCount += 1
		}
	}
	fmt.Println("Total: ", len(coords))
	fmt.Println("Overlap: ", overlapCount)
}
