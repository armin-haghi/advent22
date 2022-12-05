package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {

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
