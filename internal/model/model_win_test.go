package model_test

import (
	"testing"

	"github.com/MelleKoning/drieopeenrij/internal/model"
)

func TestWinner(t *testing.T) {
	brd := model.NewBord()

	// setup a win
	brd.Field[1] = model.X
	brd.Field[2] = model.X
	brd.Field[3] = model.X

	if brd.Winner() != model.X {
		t.Error()
	}
}
