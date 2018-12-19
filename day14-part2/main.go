package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
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
	searchScores string
	scores       []int
	elves        []int
}

func (cc *ChocolateChart) Generations() int {
	match := -1
	for match == -1 {
		if len(cc.scores)%100000 == 0 {
			fmt.Println("LLLL", len(cc.scores))
		}
		cc.AddGeneration()
		match = cc.Match()
	}
	return match
}

func (cc ChocolateChart) Match() int {
	if len(cc.scores) < len(cc.searchScores)+1 {
		return -1
	}
	var scoreStr bytes.Buffer
	offset := len(cc.scores) - len(cc.searchScores) - 1
	for _, value := range cc.scores[offset:] {
		scoreStr.WriteString(strconv.Itoa(value))
	}
	idx := strings.Index(scoreStr.String(), cc.searchScores)
	if idx == -1 {
		return -1
	}
	return idx + offset
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
	cc.scores = append(cc.scores, score)
}

func (cc *ChocolateChart) MoveElves() {
	for i, elfIndex := range cc.elves {
		score := cc.scores[elfIndex]
		cc.elves[i] = (elfIndex + score + 1) % len(cc.scores)
	}
}

func (cc ChocolateChart) ScoreSum() int {
	sum := 0
	for _, scoreIndex := range cc.elves {
		sum += cc.scores[scoreIndex]
	}
	return sum
}

func NewChocolateChart(searchScores string, elfScores []int) ChocolateChart {
	scores := make([]int, len(elfScores))
	elves := make([]int, len(elfScores))
	for i, score := range elfScores {
		scores[i] = score
		elves[i] = i
	}
	return ChocolateChart{searchScores, scores, elves}
}

func main() {
	cc := NewChocolateChart("640441", []int{3, 7})
	generations := cc.Generations()
	fmt.Println("Generations", generations)

	fmt.Println("Done")
}
