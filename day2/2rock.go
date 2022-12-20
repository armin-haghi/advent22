package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func rpsScore(opponentMove string, playerMove string) int {
	var result int
	winScore := 6
	drawScore := 3
	lossScore := 0

	// Rock: A / X
	// Paper: B / Y
	// Scissor: C / Z

	winningCases := (opponentMove == "Rock" && playerMove == "Paper") ||
		(opponentMove == "Paper" && playerMove == "Scissor") ||
		(opponentMove == "Scissor" && playerMove == "Rock")

	if opponentMove == playerMove {
		result = drawScore
	} else if winningCases {
		result = winScore
	} else {
		result = lossScore
	}
	return result
}

func main() {

	f, err := os.Open("2rock.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var totalScore int
	opponentTranslation := make(map[string]string)
	opponentTranslation["A"] = "Rock"
	opponentTranslation["B"] = "Paper"
	opponentTranslation["C"] = "Scissor"

	playerTranslation := make(map[string]string)
	playerTranslation["X"] = "Rock"
	playerTranslation["Y"] = "Paper"
	playerTranslation["Z"] = "Scissor"

	for scanner.Scan() {

		currentLine := scanner.Text()
		lineElements := strings.Split(currentLine, " ")
		opponentMove := opponentTranslation[lineElements[0]]
		playerMove := playerTranslation[lineElements[1]]

		var shapeScore int

		switch playerMove {
		case "Rock":
			shapeScore = 1
		case "Paper":
			shapeScore = 2
		case "Scissor":
			shapeScore = 3
		}

		outcomeScore := rpsScore(opponentMove, playerMove)

		totalScore += (shapeScore + outcomeScore)
		fmt.Println(playerMove, opponentMove, shapeScore, outcomeScore, totalScore)

	}

	println(totalScore)

}
