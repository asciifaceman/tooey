package tooey

import (
	"github.com/gdamore/tcell/v2"
)

// NewElement returns a stable empty Element ready to be modified
func NewElement() *Element {
	return &Element{
		Rectangle: NewRectangle(),
		Theme:     DefaultTheme,
		Border:    NewBorder(),
		Title:     NewTitle(),
	}
}

// Element is the base drawable struct inherited by most widgets
// which extends Rectangle
//
// Element manages size, position, inner drawable space,
// and each one comes with the ability to draw a Border and Title
//
// An Element bears with it a Theme.
// All other ui elements will inherit
type Element struct {
	Rectangle
	Theme  Theme
	Border *Border
	Title  *Title
}

// SetTheme will set the theme of the element
func (e *Element) SetTheme(theme Theme) {
	e.Theme = theme
	e.Border.Theme = theme
	e.Title.Theme = theme
}

// SetTitle will set the underlying title content
func (e *Element) SetTitle(title string) {
	e.Title.Set(title)
}

// SetBorderCharacters allows you to set the border characters
// for an element without touching the theme
func (e *Element) SetBorderCharacters(chars BorderChars) {
	e.Border.SetChars(chars)
}

// Draw call on the element to write to the tcell.Screen
func (e *Element) Draw(s tcell.Screen) {

	// Draw body of the element
	for y := e.Y1(); y <= e.Y2(); y++ {
		for x := e.X1(); x <= e.X2(); x++ {
			s.SetContent(x, y, ' ', nil, GetCellStyle(e.Theme, e).Style)
		}
	}

	// Draw the border
	e.Border.Draw(s, &e.Rectangle)

	// Draw the title
	e.Title.Draw(s, &e.Rectangle)

}
