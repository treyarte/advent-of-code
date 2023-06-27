package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/**
Enemy
Rock = A
Paper = b
Scissors = c

Player
Rock = X
Paper = Y
Scissors = Z

Score
Rock = 1
Paper = 2
Scissors = 3
lost = 0
draw = 3
win = 6

RockPaperScissorGame
 - given to variables
 - determine if the variable is player or enemy
 - perform checks that they are both player and enemy
 - based on input determine who won
 - return the points if lost/won/draw

 StrategyReader
 - takes a strategy file
 - read the file
 - for each line take the input
 - separate the input by space
 - check to see if we have two variables
 - call RockPaperScissorGame with two variables
 - get output and calculate it to the total
**/

var pl = fmt.Println

var totalScore int

func main() {
	totalScore = 0
	
	ReadFile()

	pl(totalScore)
}

func OnError(err error) {
	if(err != nil) {
		panic(err)
	}
}

func ReadFile() {	 
	 file, err := os.Open("rock-paper-scissors-strategy.txt")
	 OnError(err)
	 defer file.Close()

	 scanner := bufio.NewScanner(file)

	 for scanner.Scan() {
		line := scanner.Text()
		
		arr := strings.Split(line, " ")
		
		enemyInput := arr[0]
		playerInput := arr[1]

		totalScore += RockPaperScissorGame(enemyInput, playerInput)
	 }
}

func RockPaperScissorGame(enemyInput string, playerInput string) (int) {
	const (
		Win int = 6
		Lose    = 0
		Draw    = 3
	)
	
	enemyOpts := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	if _, isEnemyMoveValid := enemyOpts[enemyInput]; isEnemyMoveValid == false {
		log.Fatal("Error enemy invalid move")
		return 0
	}

	playerPoints := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	points := playerPoints[playerInput]

	if points == 0 {
		log.Fatal("Error player invalid move")
		return 0
	}

	rockOutcome := map[string]int{
		"X": Draw,
		"Y": Win,
		"Z": Lose,
	}

	scissorsOutcome := map[string]int{
		"X": Win,
		"Y": Lose,
		"Z": Draw,
	}

	paperOutcome := map [string]int {
		"X": Lose,
		"Y": Draw,
		"Z": Win,
	}

	result := 0

	if enemyInput == "A" {
		result = rockOutcome[playerInput] 
	} else if enemyInput == "B" {
		result = paperOutcome[playerInput]
	} else {
		result = scissorsOutcome[playerInput] 		
	}

	return points + result
}