package world

import (
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

// World is a two-dimensional array of cells
type World [][]*Cell

// New allocates a World with size specified by rows and columns
func New(rows, cols int) World {
	g := make([][]*Cell, rows)
	for i := range g {
		g[i] = make([]*Cell, cols)
	}

	return g
}

// Size returns the rows and columns of the world
func (w World) Size() (rows int, cols int) {
	return len(w), len((w)[0])
}

// Next computes the next generation according to
// the rules of the Game of Life according to John Conway
func (w World) Next() World {
	r, c := w.Size()
	n := New(r, c)

	for r := range w {
		for c := range w[r] {
			cnt, avg := w.Neighbors(r, c)
			if w[r][c] == nil {
				// Dead cells with 3 neighbors "Spawn"
				if cnt == 3 {
					n[r][c] = &Cell{Color: avg}
				}
			} else {
				if cnt == 2 || cnt == 3 {
					// Live on
					n[r][c] = w[r][c]
				} else {
					// Over or under populated.
					n[r][c] = nil
				}
			}
		}
	}

	return n
}

// Get returns the cell at its Row and Column.
//
// Out-of-bounds coordinates are wrapped.
//
// If the cell is dead, nil is returned.
func (w World) Get(x, y int) *Cell {
	rows, cols := w.Size()
	r, c := (x+rows)%rows, (y+cols)%cols
	return w[r][c]
}

// Neighbors counts the live cells in the Moore neighborhood of the specified coordinate.
// The average color of surrounding live cells is also computed.
//
// If there are no live cells, tcell.ColorDefault is used as the average color.
func (w World) Neighbors(x, y int) (count int, average tcell.Color) {
	var colorAvg uint64
	for r := -1; r <= 1; r++ {
		for c := -1; c <= 1; c++ {
			if !(r == 0 && c == 0) {
				cell := w.Get(x+r, y+c)
				if cell != nil {
					count++
					colorAvg += uint64(cell.Color)
				}
			}
		}
	}
	if count > 0 {
		return count, tcell.Color(colorAvg / uint64(count))
	}

	return count, tcell.ColorDefault
}

// Randomize resets the world with random cells of random colors.
func (w World) Randomize() {
	for r := range w {
		for c := range w[r] {
			if rand.Intn(2) == 1 {
				w[r][c] = Random()
			} else {
				w[r][c] = nil
			}
		}
	}
}

// Draw draws the state of the board to the screen.
func (w World) Draw(s tcell.Screen) {
	for r := range w {
		for c := range w[r] {
			cell := w.Get(r, c)
			if cell != nil {
				style := tcell.StyleDefault.Foreground(cell.Color)
				s.SetContent(r, c, tcell.RuneBlock, nil, style)
			} else {
				s.SetContent(r, c, ' ', nil, tcell.StyleDefault)
			}
		}
	}
}
