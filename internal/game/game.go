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

// EndOfGame checks for game end and returns
// true if there was a winner or the game ended in a draw
func (g *Game) EndOfGame() bool {
	return g.board.Full() || g.board.Winner() != model.EMPTY
}

// MiniMax searches the best move for the
// player who's turn it currently is.
// When a winning move is found, BestMove will contain
// the spot where to place, if the game ends in a draw
// the spot will be 0
func (g *Game) MiniMax(depth int) model.BestMove {
	if g.EndOfGame() {
		// in case no winner, model.EMPTY is returned
		// meaning game will end in a draw
		return model.BestMove{
			Spot:   0,
			Winner: g.board.Winner(),
			Depth:  depth + 1,
		}
	}

	myTurn := g.Turn
	// initialize loss and search for win
	loss := model.X

	if myTurn == model.X {
		loss = model.O
	}

	best := model.BestMove{Spot: 0, Winner: loss, Depth: depth + 1}
	// search for a winning move
	for p := 1; p <= 9; p++ {
		// is this move possible?
		if g.board.Field[p] == model.EMPTY {
			// even in case of loss, we have to do a move
			// so if we have not found a best move yet
			// lets initialize best as this move
			if best.Spot == 0 {
				best.Spot = p
			}

			g.Move(p)
			try := g.MiniMax(depth)

			if try.Winner == myTurn {
				best.Spot = p
				best.Winner = myTurn
				best.Depth = try.Depth

				g.UndoMove(p)

				break // no need to search other moves
			} else if try.Winner == model.EMPTY &&
				try.Depth < best.Depth {
				best.Spot = p
				best.Winner = model.EMPTY
				best.Depth = try.Depth
			}

			g.UndoMove(p)
		}
	}
	// if winner found, return that, other
	return best
}
