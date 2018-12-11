package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("input/d01-input.txt")
	check(err)
	defer file.Close()

	var values []int
	total := 0
	totals := make(map[int]struct{})

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		check(err)
		values = append(values, value)
	}

	for {
		for _, value := range values {
			total += value
			_, ok := totals[total]
			if ok {
				fmt.Println("First repeat total: ", total)
				os.Exit(0)
			} else {
				totals[total] = struct{}{}
			}
		}
	}
}
