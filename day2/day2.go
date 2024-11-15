package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getGames() []string {
	games, err := os.ReadFile("./input")

	if err != nil {
		fmt.Println("couldn't open file ")
		panic(err)
	}

	return strings.Split(string(games), "\n")
}

func checkIfGameIsPossible(id int, game string) int {
	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	moves := strings.Split(game, ":")[1]
	for _, move := range strings.Split(moves, ";") {
		grabbedBalls := strings.Split(move, ",")

		for i := 0; i < len(grabbedBalls); i++ {
			ball := strings.Split(grabbedBalls[i], " ")

			ballColor := ball[2]
			numberTaken, _ := strconv.Atoi(ball[1])

			switch ballColor {
			case "blue":
				if numberTaken > maxBlue {
					return 0
				}
			case "red":
				if numberTaken > maxRed {
					return 0
				}
			case "green":
				if numberTaken > maxGreen {
					return 0
				}
			}

		}
	}

	return id
}

func main() {

	possibleGames := 0
	games := getGames()

	// remove newline from end of file
	games = games[:len(games)-1]

	for index, game := range games {
		id := index + 1
		possibleGames += checkIfGameIsPossible(id, game)
	}

	fmt.Println(possibleGames)

}
