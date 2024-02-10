package main

import (
	"log"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/j0hax/gol/world"
)

func main() {
	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(defStyle)
	s.Clear()

	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	ev := make(chan tcell.Event)
	qt := make(chan struct{})

	go s.ChannelEvents(ev, qt)

	w, h := s.Size()

	world := world.New(w, h)
	world.Randomize()

	for {
		select {
		case msg := <-ev:
			switch msg := msg.(type) {
			case *tcell.EventResize:
				s.Sync()
			case *tcell.EventKey:
				if msg.Key() == tcell.KeyEscape || msg.Key() == tcell.KeyCtrlC {
					return
				} else if msg.Rune() == 'R' || msg.Rune() == 'r' {
					world.Randomize()
				}
			}
		default:
		}

		world.Draw(s)
		s.Show()
		time.Sleep(time.Second / 30)
		world = world.Next()
	}
}
