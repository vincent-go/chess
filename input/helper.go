package input

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/vincent-go/chess/basis"
)

func strToPoint(s string) (basis.Point, error) {
	s = strings.TrimSpace(s)
	p := basis.Point{}

	x, err := strconv.Atoi(string(s[0]))
	if err != nil {
		return p, errors.New("input not valid")
	}
	p.X = x

	y, err := strconv.Atoi(string(s[len(s)-1]))
	if err != nil {
		return p, errors.New("input not valid")
	}
	p.Y = y
	return p, nil
}

func read(r *bufio.Reader) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Panic("User input failed, game exit.")
	}
	text = strings.TrimSpace(text)
	return text
}

func isKingInDanger(player Player) bool {
	var king basis.Piece
	enemyMoves := []basis.Point{}
	for _, piece := range basis.Board {
		if piece.Color == player.Color {
			continue
		}
		enemyMoves = append(enemyMoves, PointsCanMoveTo(piece)...)
	}
	enemyMoves = removeDuplicate(enemyMoves)
	if player.Color == "red" {
		king = basis.RedShuai
	} else {
		king = basis.BlackJiang
	}
	if king.Current.IsInbound(enemyMoves) {
		fmt.Println(king.Color, king.Current)
		fmt.Println("Enemy moves:", enemyMoves)
		return true
	}
	return false
}

func removeDuplicate(ps []basis.Point) []basis.Point {
	noDuplicate := []basis.Point{}
	for _, p := range ps {
		if !p.IsInbound(noDuplicate) {
			noDuplicate = append(noDuplicate, p)
		}
	}
	return noDuplicate
}
