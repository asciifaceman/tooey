//go:build ignore
// +build ignore

package main

import (
	"log"

	"github.com/asciifaceman/tooey"
	"github.com/asciifaceman/tooey/themes"
	"github.com/gdamore/tcell/v2"
)

func main() {
	if err := tooey.Init(); err != nil {
		log.Fatalf("failed to initialize tooey: %v", err)
	}
	defer tooey.Close()

	x, y := tooey.DrawableDimensions()
	hello := tooey.NewElement()
	hello.SetTitle("This is a really long title that hopefully should overflow on this screen size generally")
	hello.SetRect(0, 0, x/2-1, y/2-1)

	hello2 := tooey.NewElement()
	hello2.SetTitle("This is a really long title that hopefully should overflow on this screen size generally")
	hello2.SetRect(x/2+1, y/2+1, x, y)

	tooey.SetRootTheme(themes.ThemeRetroTerminalOrange)
	tooey.Render(hello, hello2)

	for {
		tooey.Render(hello, hello2)

		ev := tooey.PollEvents()
		switch ev.(type) {
		case *tcell.EventResize:
			tooey.Clear()
			x, y := tooey.DrawableDimensions()
			hello.SetRect(0, 0, x/2-1, y/2-1)
			hello2.SetRect(x/2+1, y/2+1, x, y)
		case *tcell.EventKey:
			tooey.Close()
			return

		}
	}
}
