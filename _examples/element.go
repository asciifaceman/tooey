//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"

	"github.com/asciifaceman/tooey"
	"github.com/gdamore/tcell/v2"
)

/*
A simple screen-filling element demonstrating the lowest level
renderable object

TODO: events to exit
*/

var themes = []*tooey.RootTheme{
	tooey.DefaultTheme,
	tooey.ClassicTheme,
	&tooey.RootTheme{
		Element: tooey.WrapStyle(tcell.StyleDefault.Background(tcell.ColorRed)),
		Border:  tooey.WrapStyle(tcell.StyleDefault.Background(tcell.ColorDarkTurquoise).Foreground(tcell.ColorDarkBlue)),
		Title:   tooey.WrapStyle(tcell.StyleDefault.Background(tcell.ColorDarkTurquoise).Foreground(tcell.ColorDarkBlue)),
	},
	&tooey.RootTheme{
		Element: tooey.WrapStyle(tcell.StyleDefault.Background(tcell.ColorOrange)),
		Border:  tooey.WrapStyle(tcell.StyleDefault.Foreground(tcell.ColorAliceBlue).Background(tcell.ColorDarkMagenta)),
		Title:   tooey.ClassicTheme.Title,
	},
}

func main() {
	if err := tooey.Init(); err != nil {
		log.Fatalf("failed to initialize tooey: %v", err)
	}
	defer tooey.Close()

	x, y := tooey.DrawableDimensions()
	hello := tooey.NewElement()
	hello.SetRect(0, 0, x, y)

	iter := 0

	for {
		th := themes[iter]

		hello.SetTheme(th)
		hello.SetTitle(fmt.Sprintf("Element: %d", iter))
		tooey.Render(hello)

		ev := tooey.PollEvents()
		switch ev := ev.(type) {
		case *tcell.EventResize:
			x, y := tooey.DrawableDimensions()
			hello.SetRect(0, 0, x, y)
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEnter {
				iter++
				if iter >= len(themes) {
					iter = 0
				}
			} else {
				tooey.Close()
				return
			}

		}
	}

}
