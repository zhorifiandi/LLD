package snakeladder

import (
	"fmt"
	"math/rand/v2"
)

// This is to represent the position of a cell in the board
type Position struct {
	RowIndex int
	ColIndex int
}

type BoardCell struct {
	ID       int
	Position Position
	Type     string
}

type Board struct {
	Cells  map[int]BoardCell
	Matrix [][]*BoardCell
}

type Player struct {
	ID          string
	BoardCellID int
}

type Application struct {
	Board           Board
	Players         map[string]Player
	CurrentPlayerID string
}

func (app *Application) PrintBoard() {
	for i := 0; i < len(app.Board.Matrix); i++ {
		for j := 0; j < len(app.Board.Matrix[i]); j++ {
			fmt.Printf("%+v ", *(app.Board.Matrix[i][j]))
		}
		fmt.Println()
	}
}

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func (app *Application) RollDice() {
	player := app.Players[app.CurrentPlayerID]
	fmt.Printf("Player %s is rolling dice\n", player.ID)
	// Randomize dice value
	diceValue := randRange(1, 6)
	fmt.Printf("... Dice value is %d\n", diceValue)

	// Move player to the next cell
	nextCellID := player.BoardCellID + diceValue
	player.BoardCellID = nextCellID
}

type ApplicationInputs struct {
	PlayerIDs []string
	BoardSize int
}

func NewApplication(inputs ApplicationInputs) *Application {
	players := map[string]Player{}
	for _, playerID := range inputs.PlayerIDs {
		players[playerID] = Player{
			ID: playerID,
			// All players start in the beginning cell
			BoardCellID: 0,
		}
	}

	cells := map[int]BoardCell{}
	matrix := make([][]*BoardCell, inputs.BoardSize)

	for i := inputs.BoardSize - 1; i >= 0; i-- {
		matrix[i] = []*BoardCell{}
		for j := 0; j < inputs.BoardSize; j++ {
			actualID := (inputs.BoardSize-i-1)*inputs.BoardSize + j + 1
			if i%2 == 1 {
				actualID = (inputs.BoardSize-i)*inputs.BoardSize - j
			}

			boardCell := BoardCell{
				ID:   actualID,
				Type: "normal",
				Position: Position{
					RowIndex: i / 10,
					ColIndex: j,
				},
			}
			cells[actualID] = boardCell
			matrix[i] = append(matrix[i], &boardCell)
		}
	}

	return &Application{
		Board:           Board{Cells: cells, Matrix: matrix},
		Players:         players,
		CurrentPlayerID: players[inputs.PlayerIDs[0]].ID,
	}
}
