package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func digits(number int) (tens int, ones int) {
	numberStr := strconv.Itoa(number)
	tens, err := strconv.Atoi(string(numberStr[0]))
	check(err)
	ones, err = strconv.Atoi(string(numberStr[1]))
	check(err)
	return tens, ones
}

type ChocolateChart struct {
	generation  int
	strLength   int
	scoreLength int
	scores      []int
	elves       []int
}

func (cc *ChocolateChart) Scores() string {
	for cc.scoreLength < cc.generation+cc.strLength {
		cc.AddGeneration()
	}

	var scoreStr bytes.Buffer
	for _, value := range cc.scores[cc.generation:] {
		scoreStr.WriteString(strconv.Itoa(value))
	}

	return scoreStr.String()
}

func (cc *ChocolateChart) AddGeneration() {
	cc.AddScores()
	cc.MoveElves()
}

func (cc *ChocolateChart) AddScores() {
	sum := cc.ScoreSum()
	if sum >= 10 {
		tens, ones := digits(sum)
		cc.AddScore(tens)
		cc.AddScore(ones)
	} else {
		cc.AddScore(sum)
	}
}

func (cc *ChocolateChart) AddScore(score int) {
	if cc.scoreLength < len(cc.scores) {
		cc.scores[cc.scoreLength] = score
		cc.scoreLength++
	}
}

func (cc *ChocolateChart) MoveElves() {
	for i, elfIndex := range cc.elves {
		score := cc.scores[elfIndex]
		cc.elves[i] = (elfIndex + score + 1) % cc.scoreLength
	}
}

func (cc ChocolateChart) ScoreSum() int {
	sum := 0
	for _, scoreIndex := range cc.elves {
		sum += cc.scores[scoreIndex]
	}
	return sum
}

func NewChocolateChart(generation int, length int, elfScores []int) ChocolateChart {
	scores := make([]int, generation+length)
	elves := make([]int, len(elfScores))
	for i, score := range elfScores {
		scores[i] = score
		elves[i] = i
	}
	return ChocolateChart{generation, length, len(elfScores), scores, elves}
}

func main() {
	cc := NewChocolateChart(640441, 10, []int{3, 7})
	scores := cc.Scores()
	fmt.Println("Scores", scores, len(cc.scores))

	fmt.Println("Done")
}
