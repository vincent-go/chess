package input

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/vincent-go/chess/basis"
	"github.com/vincent-go/chess/draw"
)

// Player is the Player in the game
type Player struct {
	// ID is the palyer name
	ID string
	// Color of piece the player own
	Color string
}

// Move will move the chosen piece at location p1 to p2
func (player Player) Move(p1, p2 basis.Point) error {
	chosen := basis.Board[p1]
	if chosen.Color != player.Color {
		return errors.New("WARNING: you cannot move your opponent's piece")
	}
	moves := PointsCanMoveTo(chosen)
	if p2.IsInbound(moves) {
		delete(basis.Board, p1)
		chosen.Current = p2
		basis.Board[p2] = chosen
		return nil
	}
	return errors.New("the chosen piece can not move to this location")
}

// PointsCanMoveTo calculate the possible moves a piece can do at the current position
func PointsCanMoveTo(p basis.Piece) []basis.Point {
	ps := []basis.Point{}
	switch p.Name {
	case "Shuai":
		ps = basis.PieceKing(p).DoableMoves()
	case "Jiang":
		ps = basis.PieceKing(p).DoableMoves()
	case "Shi":
		ps = basis.PieceShi(p).DoableMoves()
	case "Xiang":
		ps = basis.PieceXiang(p).DoableMoves()
	case "Ma":
		ps = basis.PieceMa(p).DoableMoves()
	case "Ju":
		ps = basis.PieceJu(p).DoableMoves()
	case "Pao":
		ps = basis.PiecePao(p).DoableMoves()
	case "Bing":
		ps = basis.PieceBingZu(p).DoableMoves()
	case "Zu":
		ps = basis.PieceBingZu(p).DoableMoves()
	}
	ps = removeDuplicate(ps)
	return ps
}

// PlayerMove Player input the moves he wants to excecute
func PlayerMove(player Player) {
	reader := bufio.NewReader(os.Stdin)
	var p1 basis.Point
	var p2 basis.Point
	fmt.Printf("%v's turn, Please choose a piece' current location to move the piece\n", player.ID)
	for {
		p, err := strToPoint(read(reader))
		if err != nil {
			fmt.Println(err, "\nPlease try again: ")
		} else if basis.Board[p].Color != player.Color {
			fmt.Println("WARNING: Can't choose a piece that does not exist or does not belong to you \nPlease try again: ")
		} else {
			p1 = p
			break
		}
	}
	fmt.Println("Please choose a point to move the chosen piece to: ")
	for {
		p, err := strToPoint(read(reader))
		if err != nil {
			fmt.Print(err, "\nPlease try again: ")
		} else {
			p2 = p
			break
		}
	}
	// Move function will update Board when successful
	tempPiece := basis.Board[p2]
	err := player.Move(p1, p2)
	if err != nil {
		fmt.Println(err, "\nPlease try again: ")
		PlayerMove(player)
	} else if isKingInDanger(player) {
		player.Move(p2, p1)
		basis.Board[p2] = tempPiece
		fmt.Println("WARNING: This move will put king in danger \nPlease try again: ")
		draw.Board()
		PlayerMove(player)
	}
}
