package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("d1-data.txt")
	check(err)
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		value, err := strconv.Atoi(scanner.Text())
		check(err)
		total += value
	}
	fmt.Println("total: ", total)
}
