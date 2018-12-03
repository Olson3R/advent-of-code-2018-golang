package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("d2-data.txt")
	check(err)
	defer file.Close()

	two_letter_count := 0
	three_letter_count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value := scanner.Text()
		chars := make(map[rune]struct{})
		two_letter_found := false
		three_letter_found := false

		for _, char_rune := range value {
			_, ok := chars[char_rune]
			if !ok {
				chars[char_rune] = struct{}{}
				count := strings.Count(value, string(char_rune))
				if count == 2 {
					two_letter_found = true
				}
				if count == 3 {
					three_letter_found = true
				}
			}
		}

		if two_letter_found {
			two_letter_count += 1
		}
		if three_letter_found {
			three_letter_count += 1
		}
	}

	fmt.Println("two letter: ", two_letter_count)
	fmt.Println("three letter: ", three_letter_count)
	fmt.Println("Checksum: ", two_letter_count * three_letter_count)
}
