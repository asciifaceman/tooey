//go:build ignore
// +build ignore

package main

import (
	"log"

	"github.com/asciifaceman/tooey"
	"github.com/asciifaceman/tooey/themes"
	"github.com/asciifaceman/tooey/widgets"
)

/*
A simple demo application
*/

func main() {
	if err := tooey.Init(); err != nil {
		log.Fatalf("failed to initialize tooey: %v", err)
	}
	defer tooey.Close()

	x, y := tooey.DrawableDimensions()

	parent := tooey.NewContainer(themes.ThemeRetroTerminalOrange)
	parent.SetBorderCharacters(tooey.DoubleBarBorderChars)
	parent.Direction = tooey.FlexColumn
	parent.SetRect(0, 0, x, y)
	parent.SetTitle("Tooey Demo")

	headerContainer := tooey.NewContainer(themes.ThemeRetroTerminalGreen)
	headerContainer.SetBorderCharacters(tooey.RoundedBarBorderChars)
	headerContainer.Border.Enabled = false
	headerContainer.Padding = tooey.NewPadding()
	headerContainer.Direction = tooey.FlexRow

	headerContainerRight := tooey.NewContainer(themes.ThemeRetroTerminalGreen)
	headerContainerRight.Border.Enabled = false
	headerContainerRight.Direction = tooey.FlexColumn
	headerContainerRight.Padding = tooey.NewPadding()

	footerContainer := tooey.NewContainer(themes.ThemeRetroTerminalGreen)
	footerContainer.SetBorderCharacters(tooey.RoundedBarBorderChars)

	headerText := widgets.NewText(themes.ThemeRetroTerminalGreen)
	headerText.Content = "Press any key to exit"

	headerRight1 := widgets.NewText(themes.ThemeRetroTerminalGreen)
	headerRight1.Border.Enabled = false
	headerRight1.Content = "Resize to redraw"

	fullWrap := *themes.ThemeRetroTerminalGreen
	fullWrap.Text.Align = tooey.AlignFull
	headerRight2 := widgets.NewText(&fullWrap)
	headerRight2.Border.Enabled = false
	headerRight2.Content = "Tooey is a terminal UI (tui) library originally based off of gizak/termui and then rewritten under tcell by asciifaceman"

	headerContainerRight.Wrap(
		tooey.NewFlexChild(1, headerRight1),
		tooey.NewFlexChild(3, headerRight2),
	)

	headerContainer.Wrap(
		tooey.NewFlexChild(1, headerText),
		tooey.NewFlexChild(3, headerContainerRight),
	)

	parent.Wrap(
		tooey.NewFlexChild(1, headerContainer),
		tooey.NewFlexChild(5, footerContainer),
	)

	tooey.DefaultLoop(parent)

	/*
		for {
			tooey.Render(parent)

			ev := tooey.PollEvents()

			switch ev.(type) {
			case *tcell.EventResize:
				x, y := tooey.DrawableDimensions()
				parent.SetRect(0, 0, x, y)
				tooey.Render(parent)
			case *tcell.EventKey:
				tooey.Close()
				return
			}
		}
	*/
}
