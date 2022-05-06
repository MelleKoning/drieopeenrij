package model_test

import (
	"testing"

	"github.com/MelleKoning/drieopeenrij/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestBoardFullNopes(t *testing.T) {
	brd := model.NewBord()
	assert.False(t, brd.Full())
}
