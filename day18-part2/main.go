package main

/*
	- An open acre will become filled with trees if three or more adjacent acres contained trees.
		Otherwise, nothing happens.
	- An acre filled with trees will become a lumberyard if three or more adjacent acres were lumberyards.
		Otherwise, nothing happens.
	- An acre containing a lumberyard will remain a lumberyard if it was adjacent to at least one other lumberyard
		and at least one acre containing trees. Otherwise, it becomes open.
*/

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const landOpen = '.'
const landTree = '|'
const landLumberyard = '#'

const generations = 1000000000
const testGenerations = 10000
const sampleSize = 100

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func clamp(n int, min int, max int) int {
	if n < min {
		return min
	}
	if n > max {
		return max
	}
	return n
}

type LoggingZone struct {
	land [][]rune
}

func (lz *LoggingZone) ParseMap(data string) {
	data = strings.TrimSpace(data)
	rows := strings.Split(data, "\n")
	lz.land = make([][]rune, len(rows))
	for i, row := range rows {
		lz.land[i] = []rune(row)
	}
}

func (lz *LoggingZone) Next() {
	newLand := make([][]rune, len(lz.land))
	for y := range lz.land {
		newLand[y] = make([]rune, len(lz.land[y]))
		for x := range lz.land[y] {
			newLand[y][x] = lz.ProcessLand(x, y)
		}
	}
	lz.land = newLand
}

func (lz LoggingZone) ProcessLand(x int, y int) rune {
	currentLand := lz.land[y][x]
	switch currentLand {
	case landOpen:
		treeCount := 0
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}
				if lz.Land(x+i, y+j) == landTree {
					treeCount++
				}
			}
		}
		if treeCount >= 3 {
			return landTree
		}
	case landTree:
		lumberyardCount := 0
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}
				if lz.Land(x+i, y+j) == landLumberyard {
					lumberyardCount++
				}
			}
		}
		if lumberyardCount >= 3 {
			return landLumberyard
		}
	case landLumberyard:
		lumberyardFound := false
		treeFound := false
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}
				if lz.Land(x+i, y+j) == landLumberyard {
					lumberyardFound = true
				} else if lz.Land(x+i, y+j) == landTree {
					treeFound = true
				}
			}
		}
		if lumberyardFound && treeFound {
			return landLumberyard
		}
		return landOpen
	}
	return currentLand
}

func (lz LoggingZone) Land(x int, y int) rune {
	if y != clamp(y, 0, len(lz.land)-1) {
		return landOpen
	}
	if x != clamp(x, 0, len(lz.land[y])-1) {
		return landOpen
	}
	return lz.land[y][x]
}

func (lz LoggingZone) PrintAll() {
	for _, row := range lz.land {
		fmt.Println(string(row))
	}
}

func (lz LoggingZone) Value() int {
	treeCount := 0
	lumberyardCount := 0
	for _, row := range lz.land {
		for _, land := range row {
			if land == landTree {
				treeCount++
			} else if land == landLumberyard {
				lumberyardCount++
			}
		}
	}
	return treeCount * lumberyardCount
}

func main() {
	data, err := ioutil.ReadFile("../input/d18-input.txt")
	check(err)

	lz := LoggingZone{}
	lz.ParseMap(string(data))
	fmt.Println("Logging Zone", len(lz.land))
	fmt.Println("Initial State")
	lz.PrintAll()
	values := make([]int, sampleSize)
	for i := 1; i <= (testGenerations + sampleSize - 1); i++ {
		lz.Next()
		fmt.Println(" ")
		value := lz.Value()
		fmt.Println("After Minute", i, value)
		if i >= testGenerations {
			values[i-testGenerations] = value
		}
	}
	fmt.Println("Values", values)
	size := 0
	for i := 1; i < len(values); i++ {
		if values[i] == values[0] {
			size = i
			break
		}
	}
	fmt.Println("Size", size)
	index := (generations - testGenerations) % size
	fmt.Println("Value", values[index])

	fmt.Println("Done")
}
