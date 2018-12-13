package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseRow(row string) (numPlayers int, maxMarble int) {
	re := regexp.MustCompile("(\\d+) players; last marble is worth (\\d+) points")
	parts := re.FindStringSubmatch(row)
	numPlayers, err1 := strconv.Atoi(parts[1])
	check(err1)
	maxMarble, err2 := strconv.Atoi(parts[2])
	check(err2)
	return numPlayers, maxMarble * 100
}

const STEP_FORWARD = 2
const STEP_BACKWARD = -7

type MarbleGame struct {
	numPlayers int
	maxMarble  int
	index      int
	marbles    []int
}

func (mg *MarbleGame) NextIndex(increment int) {
	if increment > 0 {
		mg.index = (mg.index + increment) % mg.Len()
		if mg.index == 0 {
			mg.index = mg.Len()
		}
	} else {
		mg.index = mg.index + increment
		if mg.index < 0 {
			mg.index = mg.Len() + mg.index
		}
	}
}

func (mg MarbleGame) Play() {
	scores := make([]int, mg.numPlayers)
	for v := 0; v <= mg.maxMarble; v++ {
		player := ((v - 1) % mg.numPlayers)
		if v > 0 && v%23 == 0 {
			mg.NextIndex(STEP_BACKWARD)
			scored := v + mg.CurrentMarble()
			scores[player] += scored
			mg.RemoveCurrentMarble()
			fmt.Println("Player ", player+1, "scored", scored, "for a total of", scores[player])
		} else if mg.Len() < 2 {
			mg.AddMarble(v)
			mg.NextIndex(1)
		} else {
			mg.NextIndex(STEP_FORWARD)
			mg.AddMarble(v)
		}
	}

	maxScore := 0
	for _, score := range scores {
		if score > maxScore {
			maxScore = score
		}
	}

	fmt.Println("Max score", maxScore)
	fmt.Println("Scores", scores)
}

func (mg MarbleGame) CurrentMarble() int {
	return mg.marbles[mg.index]
}

func (mg *MarbleGame) AddMarble(value int) {
	mg.marbles = append(mg.marbles, 0)
	copy(mg.marbles[mg.index+1:], mg.marbles[mg.index:])
	mg.marbles[mg.index] = value
}

func (mg *MarbleGame) RemoveCurrentMarble() {
	mg.marbles = append(mg.marbles[:mg.index], mg.marbles[mg.index+1:]...)
}

func newMarbleGame(numPlayers int, maxMarble int) MarbleGame {
	return MarbleGame{numPlayers, maxMarble, 0, []int{}}
}

func (mg MarbleGame) Len() int {
	return len(mg.marbles)
}

func main() {
	file, err := os.Open("input/d09-input.txt")
	check(err)
	defer file.Close()

	// parse parent/child steps
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		numPlayers, maxMarble := parseRow(row)
		mg := newMarbleGame(numPlayers, maxMarble)
		mg.Play()
	}

	fmt.Println("Done")
}
