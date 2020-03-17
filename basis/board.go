package basis

// Board is the the current status of the game
var Board ChessBoard

// AllPieces contains all available pieces for the game
var AllPieces = []Piece{
	RedShuai, BlackJiang, RedShi1, RedShi2, BlackShi1, BlackShi2, RedXiang1, RedXiang2, BlackXiang1, BlackXiang2,
	RedMa1, RedMa2, BlackMa1, BlackMa2, RedJu1, RedJu2, BlackJu1, BlackJu2, RedBing1, RedBing2, RedBing3, RedBing4,
	RedBing5, BlackZu1, BlackZu2, BlackZu3, BlackZu4, BlackZu5, RedPao1, RedPao2, BlackPao1, BlackPao2,
}

var (
	RedShuai = Piece{
		Color:   "red",
		Name:    "Shuai",
		Current: Point{4, 0},
		Bound:   []Point{{3, 0}, {3, 1}, {3, 2}, {4, 0}, {4, 1}, {4, 2}, {5, 0}, {5, 1}, {5, 2}},
	}
	BlackJiang = Piece{
		Color:   "black",
		Name:    "Jiang",
		Current: Point{4, 9},
		Bound:   []Point{{3, 7}, {3, 8}, {3, 9}, {4, 7}, {4, 8}, {4, 9}, {5, 7}, {5, 8}, {5, 9}},
	}
	RedShi1 = Piece{
		Color:   "red",
		Name:    "Shi",
		Current: Point{3, 0},
		Bound:   []Point{{3, 0}, {3, 2}, {4, 1}, {5, 0}, {5, 2}},
	}
	RedShi2 = Piece{
		Color:   "red",
		Name:    "Shi",
		Current: Point{5, 0},
		Bound:   []Point{{3, 0}, {3, 2}, {4, 1}, {5, 0}, {5, 2}},
	}
	BlackShi1 = Piece{
		Color:   "black",
		Name:    "Shi",
		Current: Point{3, 9},
		Bound:   []Point{{3, 9}, {3, 7}, {4, 8}, {5, 9}, {5, 7}},
	}
	BlackShi2 = Piece{
		Color:   "black",
		Name:    "Shi",
		Current: Point{5, 9},
		Bound:   []Point{{3, 9}, {3, 7}, {4, 8}, {5, 9}, {5, 7}},
	}
	RedXiang1 = Piece{
		Color:   "red",
		Name:    "Xiang",
		Current: Point{2, 0},
		Bound:   redBoundary,
	}
	RedXiang2 = Piece{
		Color:   "red",
		Name:    "Xiang",
		Current: Point{6, 0},
		Bound:   redBoundary,
	}
	BlackXiang1 = Piece{
		Color:   "black",
		Name:    "Xiang",
		Current: Point{2, 9},
		Bound:   blackBoundary,
	}
	BlackXiang2 = Piece{
		Color:   "black",
		Name:    "Xiang",
		Current: Point{6, 9},
		Bound:   blackBoundary,
	}
	RedMa1 = Piece{
		Color:   "red",
		Name:    "Ma",
		Current: Point{1, 0},
		Bound:   noBoundary,
	}
	RedMa2 = Piece{
		Color:   "red",
		Name:    "Ma",
		Current: Point{7, 0},
		Bound:   noBoundary,
	}
	BlackMa1 = Piece{
		Color:   "black",
		Name:    "Ma",
		Current: Point{1, 9},
		Bound:   noBoundary,
	}
	BlackMa2 = Piece{
		Color:   "black",
		Name:    "Ma",
		Current: Point{7, 9},
		Bound:   noBoundary,
	}
	RedJu1 = Piece{
		Color:   "red",
		Name:    "Ju",
		Current: Point{0, 0},
		Bound:   noBoundary,
	}
	RedJu2 = Piece{
		Color:   "red",
		Name:    "Ju",
		Current: Point{8, 0},
		Bound:   noBoundary,
	}
	BlackJu1 = Piece{
		Color:   "black",
		Name:    "Ju",
		Current: Point{0, 9},
		Bound:   noBoundary,
	}
	BlackJu2 = Piece{
		Color:   "black",
		Name:    "Ju",
		Current: Point{8, 9},
		Bound:   noBoundary,
	}
	RedBing1 = Piece{
		Color:   "red",
		Name:    "Bing",
		Current: Point{0, 3},
		Bound:   append(blackBoundary, Point{0, 3}, Point{0, 4}),
	}
	RedBing2 = Piece{
		Color:   "red",
		Name:    "Bing",
		Current: Point{2, 3},
		Bound:   append(blackBoundary, Point{2, 3}, Point{2, 4}),
	}
	RedBing3 = Piece{
		Color:   "red",
		Name:    "Bing",
		Current: Point{4, 3},
		Bound:   append(blackBoundary, Point{4, 3}, Point{4, 4}),
	}
	RedBing4 = Piece{
		Color:   "red",
		Name:    "Bing",
		Current: Point{6, 3},
		Bound:   append(blackBoundary, Point{6, 3}, Point{6, 4}),
	}
	RedBing5 = Piece{
		Color:   "red",
		Name:    "Bing",
		Current: Point{8, 3},
		Bound:   append(blackBoundary, Point{8, 3}, Point{8, 4}),
	}
	BlackZu1 = Piece{
		Color:   "black",
		Name:    "Zu",
		Current: Point{0, 6},
		Bound:   append(blackBoundary, Point{0, 6}, Point{0, 5}),
	}
	BlackZu2 = Piece{
		Color:   "black",
		Name:    "Zu",
		Current: Point{2, 6},
		Bound:   append(blackBoundary, Point{2, 6}, Point{2, 5}),
	}
	BlackZu3 = Piece{
		Color:   "black",
		Name:    "Zu",
		Current: Point{4, 6},
		Bound:   append(blackBoundary, Point{4, 6}, Point{4, 5}),
	}
	BlackZu4 = Piece{
		Color:   "black",
		Name:    "Zu",
		Current: Point{6, 6},
		Bound:   append(blackBoundary, Point{6, 6}, Point{6, 5}),
	}
	BlackZu5 = Piece{
		Color:   "black",
		Name:    "Zu",
		Current: Point{8, 6},
		Bound:   append(blackBoundary, Point{8, 6}, Point{8, 5}),
	}
	RedPao1 = Piece{
		Color:   "red",
		Name:    "Pao",
		Current: Point{1, 2},
		Bound:   noBoundary,
	}
	RedPao2 = Piece{
		Color:   "red",
		Name:    "Pao",
		Current: Point{7, 2},
		Bound:   noBoundary,
	}
	BlackPao1 = Piece{
		Color:   "black",
		Name:    "Pao",
		Current: Point{1, 7},
		Bound:   noBoundary,
	}
	BlackPao2 = Piece{
		Color:   "black",
		Name:    "Pao",
		Current: Point{7, 7},
		Bound:   noBoundary,
	}
)

var (
	noBoundary    = allPoints()
	redBoundary   = redPoints()
	blackBoundary = blackPoints()
)
