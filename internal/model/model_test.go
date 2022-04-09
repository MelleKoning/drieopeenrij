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
