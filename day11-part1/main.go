package main

import (
	"fmt"
	"math"
)

const GRID_SERIAL_NUMBER = 7672
const GRID_SEARCH_SIZE = 3
const COLUMNS = 300
const ROWS = 300

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type FuelCellGrid struct {
	serialNumber int
	fuelCells    [][]int
}

func (fcg FuelCellGrid) CalculateCellPower(x int, y int) int {
	x += 1 // account for 0 offset
	y += 1 // account for 0 offset
	rackId := x + 10
	powerLevel := rackId*y + fcg.serialNumber
	powerLevel = powerLevel * rackId
	powerLevel = digit(powerLevel, 3)
	return powerLevel - 5
}

func (fcg FuelCellGrid) FindMax() (xMax int, yMax int) {
	sumMax := fcg.PodPowerLevel(xMax, yMax)
	for x := 0; x < len(fcg.fuelCells)-3; x++ {
		for y := 0; y < len(fcg.fuelCells[x])-3; y++ {
			sum := fcg.PodPowerLevel(x, y)
			if sum > sumMax {
				xMax = x
				yMax = y
				sumMax = sum
			}
		}
	}
	return xMax + 1, yMax + 1
}

func (fcg FuelCellGrid) PodPowerLevel(x int, y int) int {
	sum := 0
	for i := 0; i < GRID_SEARCH_SIZE; i++ {
		for j := 0; j < GRID_SEARCH_SIZE; j++ {
			sum += fcg.fuelCells[x+i][y+j]
		}
	}
	return sum
}

func newFuelCellGrid(serialNumber int) FuelCellGrid {
	grid := FuelCellGrid{serialNumber, make([][]int, ROWS)}
	for i := range grid.fuelCells {
		grid.fuelCells[i] = make([]int, COLUMNS)
		for j := range grid.fuelCells[i] {
			grid.fuelCells[i][j] = grid.CalculateCellPower(i, j)
		}
	}
	return grid
}

func digit(number int, place int) int {
	remainder := number % int(math.Pow(10, float64(place)))
	return remainder / int(math.Pow(10, float64(place-1)))
}

func main() {
	grid := newFuelCellGrid(GRID_SERIAL_NUMBER)
	x, y := grid.FindMax()
	fmt.Println("MAX", x, ",", y)
	fmt.Println("Done")
}
