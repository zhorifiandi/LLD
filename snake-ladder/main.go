package main

import (
	"log"

	snakeladder "github.com/zhorifiandi/LLD/snake-ladder/app"
)

type IApplication interface {
	PrintBoard()
	RollDice()
}

func main() {
	input := snakeladder.ApplicationInputs{
		PlayerIDs: []string{"player1", "player2"},
		BoardSize: 3,
	}
	var app IApplication = snakeladder.NewApplication(
		input,
	)

	log.Printf("App is running..... %+v\n", app)
	app.PrintBoard()
}
