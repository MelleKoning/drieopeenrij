package humanplayer_test

import (
	"testing"

	"github.com/MelleKoning/drieopeenrij/internal/humanplayer"
	"github.com/stretchr/testify/assert"
)

func TestGetConsoleDigit(t *testing.T) {
	moveReader := humanplayer.NewConsolePlayer()
	moveRead := moveReader.GetPlayerMove(nil)

	assert.Greater(t, moveRead.Spot, 0)
	assert.Less(t, moveRead.Spot, 10)
}
