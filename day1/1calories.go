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

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// func printSliceElements

func main() {

	f, err := os.Open("1calories.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var topN int = 3
	topSums := make([]int64, topN)
	var currentSum int64 = 0
	var calorieTotals int64

	for scanner.Scan() {

		currentLine := scanner.Text()
		if currentLine == "" {
			for i := range topSums {
				if topSums[i] < currentSum {
					if i == 0 {
						topSums = append([]int64{currentSum}, topSums[:topN-1]...)
					} else if i < topN {
						topSums = append(topSums[:i], append([]int64{currentSum}, topSums[i:topN-1]...)...)
					} else {
						topSums = append(topSums[:topN-1], []int64{currentSum}...)
					}
					break
				}
			}

			currentSum = 0
		} else {
			currentLineInt, _ := strconv.ParseInt(currentLine, 0, 64)
			currentSum = currentLineInt + currentSum
		}
	}

	for j := range topSums {
		calorieTotals += topSums[j]
	}
	fmt.Println(calorieTotals)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
