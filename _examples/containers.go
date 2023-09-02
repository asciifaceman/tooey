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
	text1.Content = "Some text in text1"
	text1.SetTheme(themes.ThemeRetroTerminalGreen)

	text2 := widgets.NewText()
	text2.Content = "Some other text in text2"
	text2.SetTheme(themes.ThemeRetroTerminalGreen)

	text3 := widgets.NewText()
	text3.Content = "Some more yet other text in text3"
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
