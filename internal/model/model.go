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

// BestMove stores the best move found for minimax
type BestMove struct {
	Spot   int           // represents the spot
	Winner FieldContents // represents winner for this move
	Depth  int           // represents at what depth this move is found
}

// Board basically consists of 9 places,
// to keep it human readable we start counting at 1 instead of 0
//     |     |
//  1  |  2  |  3
//_____|_____|_____
//     |     |
//  4  |  5  |  6
//_____|_____|_____
//     |     |
//  7  |  8  |  9
//     |     |
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

// Full checks if the board is fully occupied
func (b *Board) Full() bool {
	for p := 1; p <= 9; p++ {
		if b.Field[p] == EMPTY {
			return false
		}
	}

	return true
}

// Move puts a marker O or X on the provided place
func (b *Board) Move(piece FieldContents, place int) error {
	if piece == EMPTY {
		return fmt.Errorf("only X or O allowed for a move")
	}

	if place < 1 || place > 9 {
		return fmt.Errorf("we only support places 1..9")
	}

	if b.Field[place] != EMPTY {
		return fmt.Errorf("occupied space")
	}

	b.Field[place] = piece

	return nil
}

//UndoMove makes the place empty
func (b *Board) UndoMove(place int) {
	b.Field[place] = EMPTY
}

// PrintBoard to output the contents of the board to screen
func (b *Board) PrintBoard() {
	fmt.Println("-")
	fmt.Println("     |     |     ")
	fmt.Printf("  %s  |  %s  |  %s \n", b.printPiece(1), b.printPiece(2), b.printPiece(3))
	fmt.Println("_____|_____|_____")
	fmt.Println("     |     |     ")
	fmt.Printf("  %s  |  %s  |  %s \n", b.printPiece(4), b.printPiece(5), b.printPiece(6))
	fmt.Println("_____|_____|_____")
	fmt.Println("     |     |     ")
	fmt.Printf("  %s  |  %s  |  %s \n", b.printPiece(7), b.printPiece(8), b.printPiece(9))
	fmt.Println("     |     |     ")
}

func (b *Board) printPiece(piece int) string {
	switch b.Field[piece] {
	case EMPTY:
		return "."
	case X:
		return "X"
	case O:
		return "O"
	}

	return "."
}
