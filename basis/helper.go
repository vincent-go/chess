package basis

// all points on the board
func allPoints() []Point {
	all := []Point{}
	for i := 0; i <= 8; i++ {
		for j := 0; j <= 9; j++ {
			all = append(all, Point{i, j})
		}
	}
	return all
}

// points at the lower part of the board
func redPoints() []Point {
	red := []Point{}
	for i := 0; i <= 8; i++ {
		for j := 0; j <= 4; j++ {
			red = append(red, Point{i, j})
		}
	}
	return red
}

// points at the higher part of the board
func blackPoints() []Point {
	black := []Point{}
	for i := 0; i <= 8; i++ {
		for j := 5; j <= 9; j++ {
			black = append(black, Point{i, j})
		}
	}
	return black
}
