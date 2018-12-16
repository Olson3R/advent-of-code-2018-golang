package main

import (
	"fmt"
	"math"
)

const GRID_SERIAL_NUMBER = 7672
const SIZE = 300

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

func (fcg FuelCellGrid) FindMax() (xMax int, yMax int, sizeMax int) {
	sumMax := fcg.PodPowerLevel(0, 0, SIZE)
	sizeMax = SIZE
	for size := 1; size < SIZE; size++ {
		for x := 0; x < len(fcg.fuelCells)-size; x++ {
			for y := 0; y < len(fcg.fuelCells[x])-size; y++ {
				sum := fcg.PodPowerLevel(x, y, size)
				fmt.Println("MAX?", sum, sumMax, size)
				if sum > sumMax {
					xMax = x
					yMax = y
					sumMax = sum
					sizeMax = size
				}
			}
		}
	}
	return xMax + 1, yMax + 1, sizeMax
}

func (fcg FuelCellGrid) PodPowerLevel(x int, y int, size int) int {
	sum := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sum += fcg.fuelCells[x+i][y+j]
		}
	}
	return sum
}

func newFuelCellGrid(serialNumber int) FuelCellGrid {
	grid := FuelCellGrid{serialNumber, make([][]int, SIZE)}
	for i := range grid.fuelCells {
		grid.fuelCells[i] = make([]int, SIZE)
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
	x, y, size := grid.FindMax()
	fmt.Println("MAX", x, ",", y, ",", size)
	fmt.Println("Done")
}
