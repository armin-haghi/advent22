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

func findPlayerMove(opponentMove string, playerStrategy string) string {
	var playerMove string

	if playerStrategy == "Draw" {
		playerMove = opponentMove
	} else if playerStrategy == "Win" {
		switch opponentMove {
		case "Rock":
			playerMove = "Paper"
		case "Paper":
			playerMove = "Scissor"
		case "Scissor":
			playerMove = "Rock"
		}
	} else {
		switch opponentMove {
		case "Rock":
			playerMove = "Scissor"
		case "Paper":
			playerMove = "Rock"
		case "Scissor":
			playerMove = "Paper"
		}
	}

	return playerMove
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
	playerTranslation["X"] = "Lose"
	playerTranslation["Y"] = "Draw"
	playerTranslation["Z"] = "Win"

	for scanner.Scan() {

		currentLine := scanner.Text()
		lineElements := strings.Split(currentLine, " ")
		opponentMove := opponentTranslation[lineElements[0]]
		playerStrategy := playerTranslation[lineElements[1]]

		var shapeScore int

		playerMove := findPlayerMove(opponentMove, playerStrategy)

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
