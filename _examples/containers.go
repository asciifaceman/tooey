//go:build ignore
// +build ignore

package main

import (
	"log"
	"time"

	"github.com/asciifaceman/tooey"
	"github.com/asciifaceman/tooey/themes"
	"github.com/asciifaceman/tooey/widgets"
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

	outerContainer := tooey.NewContainer()
	outerContainer.Direction = tooey.FlexRow
	outerContainer.SetTheme(themes.ThemeRetroTerminalOrange)
	outerContainer.SetRect(0, 0, x, y)
	outerContainer.Title.Content = "Containers Example"

	text1 := widgets.NewText()
	text1.Content = "Some text in text1 which will appear pretty scrambled"
	text1.SetTheme(themes.ThemeRetroTerminalGreen)

	text2 := widgets.NewText()
	text2.Title.Content = "Left Justified"
	text2.Content = "Some other text in text2 which should attempt to wrap word-aware remove leading/trailing spaces depending on justification. Word-aware wrapping will only wrap a whole word if it could fit on a line by itself."
	text2.SetTheme(themes.ThemeRetroTerminalGreen)

	text3 := widgets.NewText()
	text3.Title.Content = "Right Justified"
	text3.Content = "This is some big text to fill some space. This text should be justifying right once RightJustify is implemented (hopefully soon) and otherwise behaving the same as LeftJustified just opposite."
	text3.SetTheme(themes.ThemeRetroTerminalGreen)

	innerContainer := tooey.NewContainer()
	innerContainer.Direction = tooey.FlexColumn
	innerContainer.SetTheme(themes.ThemeRetroTerminalOrange)
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

	tooey.Render(outerContainer)

	time.Sleep(time.Duration(time.Second * 25))

}
