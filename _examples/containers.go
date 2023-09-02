//go:build ignore
// +build ignore

package main

import (
	"log"
	"time"

	"github.com/asciifaceman/tooey"
	"github.com/asciifaceman/tooey/themes"
	"github.com/asciifaceman/tooey/widgets"
	"github.com/gdamore/tcell/v2"
)

/*
A simple screen-filling element demonstrating the lowest level
renderable object

TODO: events to exit
*/

func main() {
	if err := tooey.Init(); err != nil {
		log.Fatalf("failed to initialize tooey: %v", err)
	}
	defer tooey.Close()

	x, y := tooey.DrawableDimensions()

	outerContainer := tooey.NewContainer(themes.ThemeRetroTerminalOrange)
	outerContainer.SetBorderCharacters(tooey.DoubleBarBorderChars)
	outerContainer.Direction = tooey.FlexRow
	outerContainer.SetRect(0, 0, x, y)
	outerContainer.Title.Content = "Containers Example"

	text1 := widgets.NewText(themes.ThemeRetroTerminalGreen)
	text1.SetBorderCharacters(tooey.CornersOnlyBorderChars)
	text1.Content = "Some text in text1 which will appear pretty scrambled"

	text2 := widgets.NewText(themes.ThemeRetroTerminalGreen)
	text2.Title.Content = "Left Justified"
	text2.Content = "Some other text in text2 which should attempt to wrap word-aware remove leading/trailing spaces depending on justification. Word-aware wrapping will only wrap a whole word if it could fit on a line by itself."
	text2.SetTheme(themes.ThemeRetroTerminalGreen)
	//text2.Theme.Chars = tooey.RoundedBarBorderChars

	text3 := widgets.NewText(themes.ThemeRetroTerminalGreen)
	text3.SetBorderCharacters(tooey.RoundedBarBorderChars)
	text3.Title.Content = "Right Justified"
	text3.Content = "This is some big text to fill some space. This text should be justifying right once RightJustify is implemented (hopefully soon) and otherwise behaving the same as LeftJustified just opposite."

	innerContainer := tooey.NewContainer(themes.ThemeRetroTerminalOrange)
	innerContainer.Direction = tooey.FlexColumn
	innerContainer.Title.Content = "Inner Container"

	innerContainer.Wrap(
		tooey.NewFlexChild(1, text3),
		tooey.NewFlexChild(1, text3),
	)

	outerContainer.Wrap(
		tooey.NewFlexChild(0.5, text1),
		tooey.NewFlexChild(1, text2),
		tooey.NewFlexChild(1.5, text2),
		tooey.NewFlexChild(5, innerContainer),
	)

	for {
		tooey.Render(outerContainer)

		ev := tooey.PollEvents()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			x, y := tooey.DrawableDimensions()
			outerContainer.SetRect(0, 0, x, y)
			tooey.Render(outerContainer)
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEnter {
				tooey.Render(outerContainer)
				continue
			}
			tooey.Close()
			return
		}
	}

	time.Sleep(time.Duration(time.Second * 5))

}
