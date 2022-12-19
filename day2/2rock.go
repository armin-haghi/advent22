package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	f, err := os.Open("1calories.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		currentLine := scanner.Text()

	}

}
