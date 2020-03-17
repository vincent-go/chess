package draw

import "fmt"

// Guide will print the info about how to play the game and the board setup
func Guide() {
	guide := `************************************************************************
- 2 player's Chinese chess game
- Player 1 controls red piece, Player 2 controls black pieces
- Player 1 make the first move
- Player input the position on the board the piece is on to select
- Please then input another position on the board  to move the piece
- The board follow Cartesian coordinate system
- The left below corner is the set as point (0, 0)
- Player just need to input the digits like 00 to choose point on board
- Press Ctr + c to end the game when a player has no move available
************************************************************************`
	fmt.Println(guide)
}
