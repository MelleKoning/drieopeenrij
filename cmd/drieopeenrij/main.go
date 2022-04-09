package main

import (
	"fmt"

	"github.com/MelleKoning/drieopeenrij/internal/model"
)

func main() {
	fmt.Println("drie op een rij")

	b := model.NewBord()

	b.PrintBoard()
}
