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

func TestWinners(t *testing.T) {
	type move struct {
		pos int
		c   model.FieldContents
	}

	type test struct {
		name        string
		moves       []move
		expectedWin model.FieldContents
	}

	tests := []test{
		{
			name: "simple-win-x",
			moves: []move{
				{1, model.X},
				{2, model.X},
				{3, model.X}},
			expectedWin: model.X,
		},
		{
			name: "simple-win-O",
			moves: []move{
				{1, model.O},
				{5, model.O},
				{9, model.O}},
			expectedWin: model.O,
		},
		{
			name: "no-win",
			moves: []move{
				{1, model.X},
				{5, model.O},
				{9, model.O}},
			expectedWin: model.EMPTY,
		},
		{
			name: "win-bottom-row",
			moves: []move{
				{7, model.O},
				{8, model.O},
				{9, model.O}},
			expectedWin: model.O,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			brd := model.NewBord()
			// act
			for _, m := range tt.moves {
				_ = brd.Move(m.c, m.pos)
			}
			// assert
			if tt.expectedWin != brd.Winner() {
				t.Errorf("unexpected winner %v", brd.Winner())
				brd.PrintBoard()
			}
		})
	}
}
