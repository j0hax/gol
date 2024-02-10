package world

import (
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

type Cell struct {
	Color tcell.Color
}

// Random creates a new Cell with random color
func Random() *Cell {
	return &Cell{
		Color: tcell.Color(rand.Uint64()),
	}
}
