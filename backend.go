// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package tooey

import (
	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
)

var scrn tcell.Screen

// Init initializes the screen and sets default styling
func Init() error {
	// This says it is deprecated and you only need to import the package
	// but autoimport removes unused imports so not sure if they meant
	// importing tcell or encoding...
	encoding.Register()

	s, err := tcell.NewScreen()
	if err != nil {
		return err
	}
	if err = s.Init(); err != nil {
		return err
	}

	s.SetStyle(StyleDefault.Style)
	s.Clear()
	s.Sync()

	scrn = s

	return nil
}

// InitSim is just to support testing
func InitSim() error {
	s := tcell.NewSimulationScreen("")
	if err := s.Init(); err != nil {
		return err
	}

	defaultStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	s.SetStyle(defaultStyle)
	s.Clear()
	s.Sync()
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

// SetRootTheme sets the global theme to the Element style of the passed Theme
func SetRootTheme(theme Theme) {
	scrn.SetStyle(GetCellStyle(theme, &Element{}).Style)
}

// GetRootScreen returns the root screen tooey writes to
func GetRootScreen() tcell.Screen {
	return scrn
}

// PollEvents returns a poll of events for the
// root screen
func PollEvents() tcell.Event {
	return scrn.PollEvent()
}

// Resize handles resize events by syncing the screen and returning the new
// drawable dimensions (gently padded to prevent off-screen draws)
//
// returns x, y
func Resize() (int, int) {
	Sync()
	return DrawableDimensions()
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
	width, height := scrn.Size()
	return width, height
}

// Sync should only be called in limited circumstances
func Sync() {
	scrn.Sync()
}

// Clear the global screen
func Clear() {
	scrn.Clear()
}

// DefaultLoop is a very simple loop mostly used in the examples
// but is viable if you want a very simple runtime loop that renders
// a parent container and waits for keypress to exit while responding
// to resize events
func DefaultLoop(parent *Container) {
	for {
		Render(parent)

		ev := PollEvents()

		switch ev.(type) {
		case *tcell.EventResize:
			x, y := DrawableDimensions()
			parent.SetRect(0, 0, x, y)
			Render(parent)
		case *tcell.EventKey:
			Close()
			return
		}
	}
}
