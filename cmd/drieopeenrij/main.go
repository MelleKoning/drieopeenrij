package main

import (
	"fmt"

	"github.com/MelleKoning/drieopeenrij/internal/game"
)

func main() {
	fmt.Println("Boter, Kaas en Eieren")

	g := game.NewGame()
	moveReader := game.NewConsolePlayer()

	g.PrintBoard()

	for !g.EndOfGame() {
		bestmove := g.MiniMax(100)
		g.Move(bestmove.Spot)
		g.PrintBoard()

		if g.EndOfGame() {
			break
		}

		moveRead := moveReader.GetPlayerMove(nil)
		g.Move(moveRead.Spot)
		g.PrintBoard()
	}
}
