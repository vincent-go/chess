package basis

// Point is the Point of the chess board
type Point struct {
	X, Y int
}

// ChessBoard is the a map of all the pieces on the board
type ChessBoard map[Point]Piece

// IsInbound checks if the point is in the slice of points
func (p Point) IsInbound(b []Point) bool {
	for _, v := range b {
		if p == v {
			return true
		}
	}
	return false
}

func (p Point) isEmpty() bool {
	if _, occupied := Board[p]; occupied {
		return false
	}
	return true
}

// Piece is the base struct foor any specific type of pieces
type Piece struct {
	Color   string
	Name    string
	Current Point
	Bound   []Point
}

// DoableMoves find moves a piece could do in the current game status, implenment here to access the field only
func (s Piece) DoableMoves() []Point {
	return []Point{}
}

// Movable is an interface for all pieces
type Movable interface {
	DoableMoves() []Point
}

// PieceKing is the king with red Color
type PieceKing Piece

// DoableMoves find moves a piece could do in the current game status.
func (s PieceKing) DoableMoves() []Point {
	ps := []Point{
		Point{s.Current.X, s.Current.Y + 1},
		Point{s.Current.X, s.Current.Y - 1},
		Point{s.Current.X + 1, s.Current.Y},
		Point{s.Current.X - 1, s.Current.Y},
	}
	filtered := ps[:0]
	for _, p := range ps {
		if Board[p].Color != s.Color && p.IsInbound(s.Bound) {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

// PieceShi is the shi with red Color
type PieceShi Piece

// DoableMoves find moves a piece could do in the current game status.
func (s PieceShi) DoableMoves() []Point {
	ps := []Point{
		Point{s.Current.X + 1, s.Current.Y + 1},
		Point{s.Current.X - 1, s.Current.Y - 1},
		Point{s.Current.X + 1, s.Current.Y - 1},
		Point{s.Current.X - 1, s.Current.Y + 1},
	}
	filtered := ps[:0]
	for _, p := range ps {
		if Board[p].Color != s.Color && p.IsInbound(s.Bound) {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

// PieceXiang is the xiang with red or black Color
type PieceXiang Piece

func xiangMoveBlocked(src, dst Point) bool {
	interferePoint := Point{dst.X - src.X, dst.Y - src.Y}
	if interferePoint.isEmpty() {
		return false
	}
	return true
}

// DoableMoves find moves a piece could do in the current game status.
func (s PieceXiang) DoableMoves() []Point {
	ps := []Point{
		Point{s.Current.X + 2, s.Current.Y + 2},
		Point{s.Current.X - 2, s.Current.Y - 2},
		Point{s.Current.X + 2, s.Current.Y - 2},
		Point{s.Current.X - 2, s.Current.Y + 2},
	}
	filtered := ps[:0]
	for _, p := range ps {
		if Board[p].Color != s.Color && p.IsInbound(s.Bound) && !xiangMoveBlocked(s.Current, p) {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

// PieceMa is the ma with red or black Color
type PieceMa Piece

func maMoveBlocked(src, dst Point) bool {
	var interferePoint Point
	switch h, v := dst.X-src.X, dst.Y-src.Y; {
	case v == 2:
		interferePoint = Point{src.X, src.Y + 1}
	case v == -2:
		interferePoint = Point{src.X, src.Y - 1}
	case h == 2:
		interferePoint = Point{src.X + 1, src.Y}
	case h == -2:
		interferePoint = Point{src.X - 1, src.Y}
	}
	if interferePoint.isEmpty() {
		return false
	}
	return true
}

// DoableMoves find moves a piece could do in the current game status.
func (s PieceMa) DoableMoves() []Point {
	ps := []Point{
		Point{s.Current.X + 1, s.Current.Y + 2},
		Point{s.Current.X + 2, s.Current.Y + 1},
		Point{s.Current.X + 2, s.Current.Y - 1},
		Point{s.Current.X + 1, s.Current.Y - 2},
		Point{s.Current.X - 1, s.Current.Y + 2},
		Point{s.Current.X - 2, s.Current.Y + 1},
		Point{s.Current.X - 2, s.Current.Y - 1},
		Point{s.Current.X - 1, s.Current.Y - 2},
	}
	filtered := ps[:0]
	for _, p := range ps {
		if Board[p].Color != s.Color && p.IsInbound(s.Bound) && !maMoveBlocked(s.Current, p) {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

// PieceJu is the Ju with red or black Color
type PieceJu Piece

// DoableMoves find moves a piece could do in the current game status.
func (s PieceJu) DoableMoves() []Point {
	ps := []Point{}
	// explore left side
	for i := s.Current.X - 1; i >= 0; i-- {
		p := Point{i, s.Current.Y}
		if p.isEmpty() {
			ps = append(ps, p)
		} else if Board[p].Color != s.Color {
			ps = append(ps, p)
			break
		} else {
			break
		}
	}
	// explore right side
	for i := s.Current.X + 1; i <= 8; i++ {
		p := Point{i, s.Current.Y}
		if p.isEmpty() {
			ps = append(ps, p)
		} else if Board[p].Color != s.Color {
			ps = append(ps, p)
			break
		} else {
			break
		}
	}
	// explore lower side
	for i := s.Current.Y - 1; i >= 0; i-- {
		p := Point{s.Current.X, i}
		if p.isEmpty() {
			ps = append(ps, p)
		} else if Board[p].Color != s.Color {
			ps = append(ps, p)
			break
		} else {
			break
		}
	}
	// explore upper side
	for i := s.Current.Y + 1; i <= 9; i++ {
		p := Point{s.Current.X, i}
		if p.isEmpty() {
			ps = append(ps, p)
		} else if Board[p].Color != s.Color {
			ps = append(ps, p)
			break
		} else {
			break
		}
	}
	return ps
}

// PieceBingZu is the bing/zu with red or black Color
type PieceBingZu Piece

// DoableMoves find moves a piece could do in the current game status.
func (s PieceBingZu) DoableMoves() []Point {
	ps := []Point{}
	if s.Color == "red" {
		if s.Current.Y < 5 {
			p := Point{s.Current.X, s.Current.Y + 1}
			if p.isEmpty() || Board[p].Color != s.Color {
				ps = append(ps, p)
			}
		} else {
			temp := []Point{
				Point{s.Current.X, s.Current.Y + 1},
				Point{s.Current.X + 1, s.Current.Y},
				Point{s.Current.X - 1, s.Current.Y},
			}
			filtered := temp[:0]
			for _, p := range temp {
				if Board[p].Color != s.Color && p.IsInbound(s.Bound) {
					filtered = append(filtered, p)
				}
			}
			ps = append(ps, filtered...)
		}
	}
	if s.Color == "black" {
		if s.Current.Y > 4 {
			p := Point{s.Current.X, s.Current.Y - 1}
			if p.isEmpty() || Board[p].Color != s.Color {
				ps = append(ps, p)
			}
		} else {
			temp := []Point{
				Point{s.Current.X, s.Current.Y - 1},
				Point{s.Current.X + 1, s.Current.Y},
				Point{s.Current.X - 1, s.Current.Y},
			}
			filtered := temp[:0]
			for _, p := range temp {
				if Board[p].Color != s.Color && p.IsInbound(s.Bound) {
					filtered = append(filtered, p)
				}
			}
			ps = append(ps, filtered...)
		}
	}
	return ps
}

// PiecePao is the pao with red or black Color
type PiecePao Piece

// DoableMoves find moves a piece could do in the current game status.
func (s PiecePao) DoableMoves() []Point {
	ps := []Point{}
	// explore left side
	counter := 0
	for i := s.Current.X - 1; i >= 0; i-- {
		p := Point{i, s.Current.Y}
		if p.isEmpty() {
			ps = append(ps, p)
		} else {
			counter++
		}
		if Board[p].Color != s.Color && counter == 2 {
			ps = append(ps, p)
			break
		}
	}
	// explore right side
	counter = 0
	for i := s.Current.X + 1; i <= 8; i++ {
		p := Point{i, s.Current.Y}
		if p.isEmpty() {
			ps = append(ps, p)
		} else {
			counter++
		}
		if Board[p].Color != s.Color && counter == 2 {
			ps = append(ps, p)
			break
		}
	}
	// explore lower side
	counter = 0
	for i := s.Current.Y - 1; i >= 0; i-- {
		p := Point{s.Current.X, i}
		if p.isEmpty() {
			ps = append(ps, p)
		} else {
			counter++
		}
		if Board[p].Color != s.Color && counter == 2 {
			ps = append(ps, p)
			break
		}
	}
	// explore upper side
	counter = 0
	for i := s.Current.Y + 1; i <= 9; i++ {
		p := Point{s.Current.X, i}
		if p.isEmpty() {
			ps = append(ps, p)
		} else {
			counter++
		}
		if Board[p].Color != s.Color && counter == 2 {
			ps = append(ps, p)
			break
		}
	}
	return ps
}
