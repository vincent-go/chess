package draw

import (
	"fmt"
	"sort"

	"github.com/vincent-go/chess/basis"
)

// Board will print the keys and its position to terminal
func Board() {
	fmt.Println("Current on chess board:")
	keys := []basis.Point{}
	for key := range basis.Board {
		keys = append(keys, key)
		// keys = quickSort(keys)
		sort.SliceStable(keys, func(i, j int) bool {
			if keys[i].Y == keys[j].Y && keys[i].X < keys[j].X {
				return true
			}
			if keys[i].Y < keys[j].Y {
				return true
			}
			return false
		})
	}
	for _, k := range keys {
		fmt.Printf("Point %v is occupied by %v\n", k, basis.Board[k].Color+basis.Board[k].Name)
	}
}
