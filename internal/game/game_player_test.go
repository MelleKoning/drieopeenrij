package game_test

import (
	"testing"

	"github.com/MelleKoning/drieopeenrij/internal/game"
	"github.com/stretchr/testify/assert"
)

func TestGetConsoleDigit(t *testing.T) {
	moveReader := game.NewConsolePlayer()
	moveRead := moveReader.GetPlayerMove(nil)

	assert.Greater(t, moveRead.Spot, 0)
	assert.Less(t, moveRead.Spot, 10)
}
