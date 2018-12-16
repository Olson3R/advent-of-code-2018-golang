package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func parseInitialRow(row string) string {
	re := regexp.MustCompile("initial state: ([\\.#]+)")
	parts := re.FindStringSubmatch(row)
	return parts[1]
}

func parseMaskRow(row string) (mask string, value string) {
	re := regexp.MustCompile("([\\.#]{5}) => ([\\.#])")
	parts := re.FindStringSubmatch(row)
	return parts[1], parts[2]
}

type PlantPredictor struct {
	state  string
	masks  map[string]string
	offset int
}

func (pp *PlantPredictor) Generations(number int) string {
	fmt.Println("0 :", pp.state)
	for g := 1; g <= number; g++ {
		newState := ""
		state := "...." + pp.state + "...."
		for i := 2; i < len(state)-2; i++ {
			plantValues := state[i-2 : i+3]
			value, ok := pp.masks[plantValues]
			if ok {
				newState += value
			} else {
				newState += "."
			}
		}
		offsetAdjustment := 2 - strings.Index(newState, "#")
		pp.state = strings.Trim(newState, ".")
		pp.offset += offsetAdjustment
		fmt.Println(g, ":", pp.state, pp.offset)
	}
	return pp.state
}

func (pp PlantPredictor) Sum() int {
	sum := 0
	for i, value := range pp.state {
		if value == '#' {
			sum += i - pp.offset
		}
	}
	return sum
}

func main() {
	data, err := ioutil.ReadFile("input/d12-input.txt")
	check(err)

	parts := strings.Split(string(data), "\n")
	initialState := parseInitialRow(parts[0])

	masks := make(map[string]string)
	for _, maskString := range parts[2 : len(parts)-1] {
		mask, value := parseMaskRow(maskString)
		masks[mask] = value
	}

	plantPredictor := PlantPredictor{initialState, masks, 0}
	state := plantPredictor.Generations(20)

	fmt.Println("STATE", state)
	fmt.Println("SUM", plantPredictor.Sum())

	fmt.Println("Done")
}
