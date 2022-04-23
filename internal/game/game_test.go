package game_test

import (
	"testing"

	"github.com/MelleKoning/drieopeenrij/internal/game"
)

func TestPlayRandomGame(t *testing.T) {
	trialGame := game.NewGame()

	_ = trialGame.PlayRandomGame()
}
