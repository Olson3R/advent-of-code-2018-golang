package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"unicode"
)

type Box struct {
	id int
	xBegin int
	xEnd int
	yBegin int
	yEnd int
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
		if r == '#' {
			return -1
		}
		return r
	}, str)
}

func overlap(a Box, b Box) bool {
	if a.xBegin <= b.xEnd && a.xEnd >= b.xBegin &&
			a.yBegin <= b.yEnd && a.yEnd >= b.yBegin {
		return true
	}
	return false
}

func parseRow(row string) (int, int, int, int, int) {
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
	id, err := strconv.Atoi(parts[0])
	check(err)
	x, err := strconv.Atoi(parts[1])
	check(err)
	y, err := strconv.Atoi(parts[2])
	check(err)
	width, err := strconv.Atoi(parts[3])
	check(err)
	height, err := strconv.Atoi(parts[4])
	check(err)
	return id, x + 1, x + width, y + 1, y + height
}

func main() {
	file, err := os.Open("d3-data.txt")
	check(err)
	defer file.Close()

	var boxes []Box
	boxOverlaps := make(map[int]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strip(scanner.Text())
		id, xBegin, xEnd, yBegin, yEnd := parseRow(row)
		boxes = append(boxes, Box{id, xBegin, xEnd, yBegin, yEnd})
		boxOverlaps[id] = false
	}

	for boxIndex, box := range boxes {
		for _, otherBox := range boxes[boxIndex + 1:] {
			if overlap(box, otherBox) {
				boxOverlaps[box.id] = true
				boxOverlaps[otherBox.id] = true
			}
		}
	}


	for id, overlaps := range boxOverlaps {
		if !overlaps {
			fmt.Println("No overlap: ", id)
		}
	}
}
