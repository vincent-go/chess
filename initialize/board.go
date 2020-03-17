package initialize

import "github.com/vincent-go/chess/basis"

// SetupBoard will set initialize the board
func SetupBoard() basis.ChessBoard {
	basis.Board = make(map[basis.Point]basis.Piece)
	for _, p := range basis.AllPieces {
		basis.Board[p.Current] = p
	}
	return basis.Board
}
