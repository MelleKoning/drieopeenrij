package humanplayer

import (
	"bufio"
	"fmt"
	"os"

	"github.com/MelleKoning/drieopeenrij/internal/game"
	"github.com/MelleKoning/drieopeenrij/internal/model"
)

// MoveHumanReader is an interface for implementing
// a way for the user to play a move
type MoveHumanReader interface {
	GetPlayerMove(gamestate *game.Game) model.PlayerMove
}

// ConsolePlayer is a struct for implementing a
// simple console implementation of MoveHumanReader
type ConsolePlayer struct{}

// NewConsolePlayer creates a Console version of MoveHumanReader
func NewConsolePlayer() MoveHumanReader {
	return &ConsolePlayer{}
}

// GetPlayerMove reads from console a move or returns 0 if player wants to quit
func (c *ConsolePlayer) GetPlayerMove(gamestate *game.Game) model.PlayerMove {
	fmt.Print("Your move (1-9) 0 to quit:")

	scanner := bufio.NewScanner(os.Stdin)
	spot := -1

	for spot < 0 || spot > 9 {
		var char string

		for scanner.Scan() {
			char = scanner.Text()
			break
		}

		spot = int(char[0] - '0')
	}

	return model.PlayerMove{Spot: spot, Field: model.EMPTY}
}
