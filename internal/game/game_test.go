package game_test

import (
	"fmt"
	"testing"

	"github.com/MelleKoning/drieopeenrij/internal/game"
	"github.com/MelleKoning/drieopeenrij/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestWinners(t *testing.T) {
	type test struct {
		name        string
		startTurn   model.FieldContents
		movelist    []int
		expectedWin model.FieldContents
	}

	tests := []test{
		{
			name:        "simple-win-x",
			startTurn:   model.X,
			movelist:    []int{1, 4, 2, 5, 3},
			expectedWin: model.X,
		},
		{
			name:        "simple-win-O",
			startTurn:   model.O,
			movelist:    []int{1, 2, 5, 8, 9},
			expectedWin: model.O,
		},
		{
			name:        "no-win",
			startTurn:   model.X,
			movelist:    []int{1, 5, 9, 2, 8, 7, 3, 6, 4}, // draw
			expectedWin: model.EMPTY,
		},
		{
			name:        "win-bottom-row",
			startTurn:   model.X,
			movelist:    []int{7, 1, 8, 5, 9},
			expectedWin: model.X,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			g := game.NewGame()
			g.Turn = tt.startTurn
			// act
			for _, m := range tt.movelist {
				g.Move(m)
			}
			// assert
			// assign to private var
			g.PrintBoard()
			if tt.expectedWin != g.Winner() {
				t.Errorf("unexpected winner %v", g.Winner())

			}
		})
	}
}

func TestPlayRandomGame(t *testing.T) {
	trialGame := game.NewGame()

	_ = trialGame.PlayRandomGame()
}

func TestGetBestMove(t *testing.T) {
	// given a boardstate, we assume minimax
	// can find the best move
	trialGame := game.NewGame()
	trialGame.Turn = model.X
	// do moves so that a win can be made
	trialGame.Move(1) // X
	trialGame.Move(2) // O
	trialGame.Move(7)
	trialGame.Move(4)
	trialGame.PrintBoard()
	// now X can win..
	for {
		getBestMove := trialGame.MiniMax(1)
		fmt.Printf("%v, %v", getBestMove, trialGame.Turn)
		trialGame.Move(getBestMove.Spot)
		trialGame.PrintBoard()

		if trialGame.EndOfGame() {
			break
		}
	}

	assert.Equal(t, model.X, trialGame.Winner())
}
