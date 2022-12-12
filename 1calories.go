package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	// This program outputs the largest sum of integer
	// elements that are separated by an empty line.
	//
	// To do so it uses two variables:
	// - currentSum tracks the sum of the current group of elements
	// - previousSum tracks the highest sum of elements
	//
	// The main loop goes through every line of the file.
	// - When an empty line is reached, it compares the two sums and updates
	//   previousSum to the highest value of both, then resets currentSum to 0
	//   for the next iteration.
	// - When a number is in the next line, it gets added to currentSum

	f, err := os.Open("calories.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var previousSum int64 = 0
	var currentSum int64 = 0

	for scanner.Scan() {

		currentLine := scanner.Text()
		if currentLine == "" {
			if previousSum < currentSum {
				previousSum = currentSum
			}
			currentSum = 0
		} else {
			currentLineInt, _ := strconv.ParseInt(currentLine, 0, 64)
			currentSum = currentLineInt + currentSum
		}
	}

	println(previousSum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
