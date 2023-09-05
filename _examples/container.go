//go:build ignore
// +build ignore

package main

import (
	"log"

	"github.com/asciifaceman/tooey"
)

/*
A demonstration of the flexible containers
*/

func main() {
	if err := tooey.Init(); err != nil {
		log.Fatalf("failed to initialize tooey: %v", err)
	}
	defer tooey.Close()

	x, y := tooey.DrawableDimensions()

	parent := tooey.NewContainer()
	parent.SetBorderCharacters(tooey.DoubleBarBorderChars)
	parent.Direction = tooey.FlexRow
	parent.SetRect(0, 0, x, y)
	parent.SetTitle("Parent Container")

	firstRow := tooey.NewContainer()
	firstRow.SetTitle("Flexible Row")
	firstRow.SetBorderCharacters(tooey.RoundedBarBorderChars)
	firstRow.Direction = tooey.FlexRow

	firstRow.Wrap(
		tooey.NewFlexChild(1, tooey.NewContainer()),
		tooey.NewFlexChild(2, tooey.NewContainer()),
		tooey.NewFlexChild(3, tooey.NewContainer()),
		tooey.NewFlexChild(4, tooey.NewContainer()),
		tooey.NewFlexChild(5, tooey.NewContainer()),
	)

	secondRow := tooey.NewContainer()
	secondRow.SetTitle("Flexible Column")
	secondRow.SetBorderCharacters(tooey.RoundedBarBorderChars)
	secondRow.Direction = tooey.FlexColumn

	secondRow.Wrap(
		tooey.NewFlexChild(1, tooey.NewContainer()),
		tooey.NewFlexChild(2, tooey.NewContainer()),
		tooey.NewFlexChild(3, tooey.NewContainer()),
		tooey.NewFlexChild(4, tooey.NewContainer()),
		tooey.NewFlexChild(5, tooey.NewContainer()),
	)

	parent.Wrap(
		tooey.NewFlexChild(1, firstRow),
		tooey.NewFlexChild(1, secondRow),
	)

	tooey.DefaultLoop(parent)
}
