//go:build ignore
// +build ignore

package main

import (
	"log"
	"time"

	"github.com/asciifaceman/tooey"
	"github.com/asciifaceman/tooey/themes"
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

	hello := tooey.NewElement()
	hello.SetTheme(themes.ThemeRetroTerminalOrange)

	hello.SetRect(0, 0, x, y)
	hello.Title.Content = "Example"

	tooey.Render(hello)

	time.Sleep(time.Duration(time.Second * 5))

}
