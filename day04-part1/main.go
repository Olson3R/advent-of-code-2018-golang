package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Guard struct {
	id     int
	asleep []int
}

func AddSleepPeriod(g *Guard, start int, end int) {
	for i := start; i < end; i++ {
		fmt.Println("WWW", i, g.asleep[i]+1)
		g.asleep[i] = g.asleep[i] + 1
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseRow(row string) (string, string) {
	re := regexp.MustCompile("\\[(.+)\\] (.*)")
	parts := re.FindStringSubmatch(row)
	time := parts[1]
	action := strings.TrimSpace(parts[2])
	return time, action
}

func parseGuardId(action string) int {
	re := regexp.MustCompile("Guard #(\\d+) begins shift")
	parts := re.FindStringSubmatch(action)
	if len(parts) == 2 {
		id, err := strconv.Atoi(parts[1])
		check(err)
		return id
	}
	return 0
}

func parseMinutes(time string) int {
	re := regexp.MustCompile(".*:(\\d+)")
	parts := re.FindStringSubmatch(time)
	minutes, err := strconv.Atoi(parts[1])
	check(err)
	return minutes
}

func main() {
	file, err := os.Open("input/d04-input.txt")
	check(err)
	defer file.Close()

	var rows []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	sort.Strings(rows)

	guards := make(map[int]*Guard)
	var guard Guard
	var start int
	for _, row := range rows {
		time, action := parseRow(row)
		guardId := parseGuardId(action)
		if guardId > 0 {
			guardRef, ok := guards[guardId]
			if !ok {
				newGuard := Guard{guardId, make([]int, 60)}
				guards[guardId] = &newGuard
				guard = newGuard
			} else {
				guard = *guardRef
			}
		} else {
			if action == "falls asleep" {
				start = parseMinutes(time)
			} else if action == "wakes up" {
				AddSleepPeriod(&guard, start, parseMinutes(time))
			}
		}
	}

	for _, guardRef := range guards {
		fmt.Println(guardRef.asleep)
	}
}
