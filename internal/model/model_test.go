package model_test

import (
	"errors"
	"testing"

	"github.com/MelleKoning/drieopeenrij/internal/model"
	"github.com/stretchr/testify/assert"
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

	_ = b.Move(model.X, 5)
	_ = b.Move(model.O, 1)
	_ = b.Move(model.X, 9)

	b.PrintBoard()
}

func TestWhenMoveOnOccupiedErrors(t *testing.T) {
	// Arrange
	b := model.NewBord()
	expectedError := errors.New("occupied space")
	err := b.Move(model.X, 5)
	assert.Nil(t, err)

	err = b.Move(model.X, 5)
	assert.Equal(t, expectedError, err)
}
