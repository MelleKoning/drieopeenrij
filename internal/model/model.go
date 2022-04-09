package model

import "fmt"

// package to represents the board
// and actions that can be done

//FieldContents represents the contents of one space on the board
type FieldContents int

// FieldContents supported
const (
	EMPTY FieldContents = iota
	X
	O
)

// Board basically consists of 9 places,
// to keep it human readable we start counting at 1 instead of 0
// 1 2 3
// 4 5 6
// 7 8 9
type Board struct {
	Field map[int]FieldContents
}

//NewBord returns a new empty board
func NewBord() *Board {
	bord := &Board{
		Field: make(map[int]FieldContents),
	}
	for i := 1; i <= 9; i++ {
		bord.Field[i] = EMPTY
	}
	return bord
}

// Move puts a marker O or X on the provided place
func (b *Board) Move(piece FieldContents, place int) error {
	if piece == EMPTY {
		return fmt.Errorf("only X or O allowed for a move")
	}
	if place < 1 || place > 9 {
		return fmt.Errorf("we only support places 1..9")
	}
	b.Field[place] = piece
	return nil
}
