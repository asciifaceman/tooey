package tooey

import (
	"github.com/gdamore/tcell/v2"
)

// NewElement returns a stable empty Element
func NewElement() *Element {
	e := &Element{
		Rectangle: NewRectangle(nil),
		Border:    NewDefaultBorder(nil),
		Title:     NewTitle(),
	}

	e.Rectangle.Padding = NewDefaultPadding()
	e.SetTheme(DefaultTheme)

	return e
}

// Element is the base drawable struct inherited by most widgets
// Element manages size, positin, inner drawable space
// All other ui elements will inherit
type Element struct {
	Rectangle
	Theme  *Theme
	Border *Border
	Title  *Title
}

func (e *Element) SetTheme(theme *Theme) {
	e.Theme = theme
	e.Border.Style = theme.Border
	e.Title.Style = theme.Title
}

func (e *Element) Draw(s tcell.Screen) {

	// Draw body of the element
	for row := e.Y1(); row <= e.Y2(); row++ {
		for col := e.X1(); col <= e.X2(); col++ {
			s.SetContent(col, row, ' ', nil, e.Theme.Element.Style)
		}
	}

	// Draw the border
	e.Border.Draw(s, &e.Rectangle)

	// Draw the title
	e.Title.Draw(s, &e.Rectangle)

}
