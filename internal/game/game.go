package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/MelleKoning/drieopeenrij/internal/model"
)

// Game represents a tic-tac-toe game
type Game struct {
	board *model.Board        // keeps track of the board
	Turn  model.FieldContents // keeps track of who's turn it is
}

//NewGame initializes a new game with an empty board
func NewGame() *Game {
	return &Game{
		board: model.NewBord(),
	}
}

// PrintBoard delegates to board to print out
func (g *Game) PrintBoard() {
	g.board.PrintBoard()
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

// Move places a piece of the player who turns it is
// and switches the turn to next player
func (g *Game) Move(place int) {
	_ = g.board.Move(g.Turn, place)

	switch g.Turn {
	case model.X:
		g.Turn = model.O
	case model.O:
		g.Turn = model.X
	}
}

// UndoMove places an empty square and switches
// back the turn to the previous player
func (g *Game) UndoMove(place int) {
	g.board.Field[place] = model.EMPTY

	switch g.Turn {
	case model.X:
		g.Turn = model.O
	case model.O:
		g.Turn = model.X
	}
}

// MiniMax searches the best move for the
// player who's turn it currently is.
// When a winning move is found, BestMove will contain
// the spot where to place, if the game ends in a draw
// the spot will be 0
func (g *Game) MiniMax() model.BestMove {
	if g.board.Full() || g.board.Winner() != model.EMPTY {
		// return the winner, but we do not know
		// what move caused this here
		return model.BestMove{Winner: g.board.Winner()}
	}

	currentTurn := g.Turn
	// initialize loss and search for win
	loss := model.X

	if currentTurn == model.X {
		loss = model.O
	}

	best := model.BestMove{Spot: 0, Winner: loss}
	// search for a winning move
	for p := 1; p <= 9; p++ {
		// move possible?
		if g.board.Field[p] == model.EMPTY {
			g.Move(p)
			try := g.MiniMax()
			try.Spot = p

			if try.Winner == currentTurn {
				best.Spot = p
				best.Winner = currentTurn
			}

			g.UndoMove(p)
		}
	}
	// if winner found, return that, other
	return best
}
