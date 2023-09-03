// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package tooey

import (
	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
)

var scrn tcell.Screen

// Init is refactor of init for tcell operation
func Init() error {
	// This says it is deprecated and you only need to import the package
	// but autoimport removes unused imports so...
	encoding.Register()

	s, err := tcell.NewScreen()
	if err != nil {
		return err
	}
	if err = s.Init(); err != nil {
		return err
	}

	defaultStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	s.SetStyle(defaultStyle)
	s.Clear()

	scrn = s

	return nil
}

// Close is refactor of close
func Close() {
	maybePanic := recover()
	scrn.Fini()
	if maybePanic != nil {
		panic(maybePanic)
	}
}

// GetRootScreen returns the root screen tooey writes to
func GetRootScreen() *tcell.Screen {
	return &scrn
}

// PollEvents returns a poll of events for the
// root screen
func PollEvents() tcell.Event {
	return scrn.PollEvent()
}

// DrawableDimensions is the same as TerminalDimensions -1 to represent visibly drawable space in
// most terminals
func DrawableDimensions() (int, int) {
	width, height := TerminalDimensions()
	return width - 1, height - 1
}

// Terminal dimensions returns an aggregate dimension for the terminal
// but it often is clipped on the right and buttom
// Use DrawableDimensions to get visible terminal dimensions
func TerminalDimensions() (int, int) {
	scrn.Sync()
	width, height := scrn.Size()
	return width, height
}

// Sync ...
func Sync() {
	scrn.Sync()
}

// Clear the global screen
func Clear() {
	scrn.Clear()
}
