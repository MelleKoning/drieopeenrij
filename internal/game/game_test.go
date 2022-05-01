package game_test

import (
	"fmt"
	"testing"

	"github.com/MelleKoning/drieopeenrij/internal/game"
	"github.com/MelleKoning/drieopeenrij/internal/model"
)

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
	getBestMove := trialGame.MiniMax()
	fmt.Printf("%v", getBestMove)
	trialGame.PrintBoard()
	trialGame.Move(getBestMove.Spot)
	trialGame.PrintBoard()
}
