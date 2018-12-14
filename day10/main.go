package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func Abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInt(value string) int {
	intValue, err := strconv.Atoi(value)
	check(err)
	return intValue
}

func parseRow(row string) (x int, y int, dx int, dy int) {
	re := regexp.MustCompile("position=<\\s*(\\-?\\d+),\\s*(\\-?\\d+)> velocity=<\\s*(\\-?\\d+),\\s*(\\-?\\d+)>")
	parts := re.FindStringSubmatch(row)
	x = parseInt(parts[1])
	y = parseInt(parts[2])
	dx = parseInt(parts[3])
	dy = parseInt(parts[4])
	return x, y, dx, dy
}

type Point struct {
	x  int
	y  int
	dx int
	dy int
}

func (p *Point) Next() (x int, y int) {
	p.x = p.x + p.dx
	p.y = p.y + p.dy
	return p.x, p.y
}

type Message struct {
	points []*Point
}

func (m *Message) AddPoint(x int, y int, dx int, dy int) {
	m.points = append(m.points, &Point{x, y, dx, dy})
}

func (m *Message) Next() {
	for _, point := range m.points {
		point.Next()
	}
}

func (m Message) Show() {
	xMin, xMax, yMin, yMax := m.FindBounds()
	board := make([][]rune, yMax-yMin+1)
	for i := range board {
		board[i] = make([]rune, xMax-xMin+1)
		for j := range board[i] {
			board[i][j] = ' '
		}
	}

	for _, point := range m.points {
		board[point.y-yMin][point.x-xMin] = 'X'
	}

	for _, row := range board {
		fmt.Println(string(row))
	}
}

func (m Message) Size() int {
	xMin, xMax, yMin, yMax := m.FindBounds()
	return Abs(xMax-xMin) + Abs(yMax-yMin)
}

func (m Message) FindBounds() (xMin int, xMax int, yMin int, yMax int) {
	xMin = m.points[0].x
	xMax = m.points[0].x
	yMin = m.points[0].y
	yMax = m.points[0].y
	for _, point := range m.points {
		if xMin > point.x {
			xMin = point.x
		} else if xMax < point.x {
			xMax = point.x
		}

		if yMin > point.y {
			yMin = point.y
		} else if yMax < point.y {
			yMax = point.y
		}
	}
	return xMin, xMax, yMin, yMax
}

func main() {
	file, err := os.Open("input/d10-input.txt")
	check(err)
	defer file.Close()

	message := Message{[]*Point{}}

	// parse parent/child steps
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		x, y, dx, dy := parseRow(row)
		message.AddPoint(x, y, dx, dy)
	}

	prevSize := -1
	i := 0
	for {
		currentSize := message.Size()
		fmt.Println(i, ":", currentSize)
		if currentSize < 100 {
			message.Show()
			fmt.Println(" ")
			time.Sleep(3 * time.Second)
		}

		message.Next()

		if prevSize < currentSize && prevSize != -1 {
			break
		}

		prevSize = currentSize
		i++
	}

	fmt.Println("Done")
}
