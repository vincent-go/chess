package main

import (
	"fmt"

	"github.com/vincent-go/chess/basis"
	"github.com/vincent-go/chess/draw"
	"github.com/vincent-go/chess/initialize"
	"github.com/vincent-go/chess/input"
)

func init() {
	basis.Board = initialize.SetupBoard()
}

func main() {
	var player1 = input.Player{
		ID:    "Player1",
		Color: "red",
	}
	var player2 = input.Player{
		ID:    "Player2",
		Color: "black",
	}
	draw.Guide()
	fmt.Println("Game started.")
	draw.Board()
	for {
		for _, player := range []input.Player{player1, player2} {
			input.PlayerMove(player)
			draw.Board()
		}
	}
}
