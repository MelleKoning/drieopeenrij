package main

import (
	"fmt"

	"github.com/MelleKoning/drieopeenrij/internal/game"
	"github.com/MelleKoning/drieopeenrij/internal/humanplayer"
	"github.com/MelleKoning/drieopeenrij/internal/model"
)

func main() {
	fmt.Println("Boter, Kaas en Eieren")

	g := game.NewGame()
	startPlayer := model.X
	moveReader := humanplayer.NewConsolePlayer()

	g.PrintBoard()

	for {
		g := game.NewGame()
		g.PrintBoard()
		g.Turn = startPlayer

		for !g.EndOfGame() {
			switch g.Turn {
			case model.X: // machine
				bestmove := g.MiniMax(100)
				g.Move(bestmove.Spot)
			case model.O: // console player
				moveRead := moveReader.GetPlayerMove(nil)
				if moveRead.Spot == 0 {
					return
				}

				g.Move(moveRead.Spot)
			}

			g.PrintBoard()

			if g.EndOfGame() {
				// next game this player to start:
				startPlayer = g.Turn
				break
			}
		}
	}
}
