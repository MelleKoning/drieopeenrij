package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/MelleKoning/drieopeenrij/internal/model"
)

// Game represents a tic-tac-toe game
type Game struct {
	board *model.Board
}

//NewGame initializes a new game with an empty board
func NewGame() *Game {
	return &Game{
		board: model.NewBord(),
	}
}

// PlayRandomGame plays a random game of tic-tac-toe
// and returns the winner
func (g *Game) PlayRandomGame() model.FieldContents {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	turn := model.X

	moves := 0
	for g.board.Winner() == model.EMPTY &&
		moves < 9 { // play random moves till there is a winner
		mv := r1.Intn(9) + 1

		// try move till empty space is found
		if err := g.board.Move(turn, mv); err != nil {
			continue
		}

		moves = moves + 1
		fmt.Println(moves)
		// print the new board to screen
		g.board.PrintBoard()

		switch turn {
		case model.X:
			turn = model.O
		case model.O:
			turn = model.X
		}
	}

	return g.board.Winner()
}
