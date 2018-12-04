package main

import (
	"fmt"
	"os"
	"bufio"
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

	var boxes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value := scanner.Text()
		boxes = append(boxes, value)
	}

	for boxIndex, box := range boxes {
		for idIndex, _ := range box {
			testId := box[:idIndex] + box[idIndex + 1:]
			for _, otherBox := range boxes[boxIndex + 1:] {
				if testId == otherBox[:idIndex] + otherBox[idIndex + 1:] {
					fmt.Println("ID: ", testId)
					fmt.Println("BOX ID: ", box)
					fmt.Println("OTHER ID: ", otherBox)
					os.Exit(0)
				}
			}
		}
	}
}
