package model_test

import (
	"testing"

	"github.com/MelleKoning/drieopeenrij/internal/model"
)

func TestBoardCreation_Empty(t *testing.T) {

	// Arrange
	b := model.NewBord()

	// Assert all fields are initialized as empty spaces
	for _, f := range b.Field {
		if f != model.EMPTY {
			t.Fail()
		}
	}
}

func TestAFewMovesOnTheBoard(t *testing.T) {
	// Arrange
	b := model.NewBord()

	b.Move(model.X, 5)
	b.Move(model.X, 1)
	b.Move(model.X, 1)
}
