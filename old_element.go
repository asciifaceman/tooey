package tooey

import (
	"image"
	"sync"

	"github.com/gdamore/tcell/v2"
)

// NewOldElement returns a new stable OldElement that
// is drawable but can be modified
func NewOldElement() *OldElement {
	return &OldElement{
		Padding:         NewDefaultPadding(),
		TitlePadding:    NewTitlePadding(),
		Border:          true,
		BorderStyle:     StyleDefault,
		OldElementStyle: StyleDefault,
		BorderLeft:      true,
		BorderRight:     true,
		BorderTop:       true,
		BorderBottom:    true,
	}
}

// OldElement is the base struct inherited by most widgets
// OldElement manages size, position, inner drawable space, borders, and titles
// All other widgets and ui OldElements will oerride the Draw method
type OldElement struct {
	// Padding allows the drawable borders of an OldElement to be inset
	Padding *Padding

	// OldElementStyle represents the style of the space the OldElement covers
	OldElementStyle Style

	image.Rectangle
	// Inner represents the inner drawable space of an OldElement
	Inner image.Rectangle

	// Border whether or not to draw borders
	Border bool
	// BorderStyle represents the styling to apply to the borders
	BorderStyle Style

	BorderLeft   bool
	BorderRight  bool
	BorderTop    bool
	BorderBottom bool

	// The optional Title of this OldElement which will be rendered in the top row
	Title string
	// TitleStyle represents the style of the title characters - this will override the border it draws on
	TitleStyle Style
	// TitlePadding defines the left and right inset for title drawing
	// top and bottom are unused
	TitlePadding *Padding

	sync.Mutex
}

func (e *OldElement) drawBody(s tcell.Screen) {
	for row := e.Y1(); row <= e.Y2(); row++ {
		for col := e.X1(); col <= e.X2(); col++ {
			s.SetContent(col, row, ' ', nil, e.OldElementStyle.Style)
		}
	}
}

func (e *OldElement) drawBorder(s tcell.Screen) {
	for col := e.X1(); col <= e.X2(); col++ {
		s.SetContent(col, e.Y1(), tcell.RuneHLine, nil, e.BorderStyle.Style)
		s.SetContent(col, e.Y2(), tcell.RuneHLine, nil, e.BorderStyle.Style)
	}

	for row := e.Y1(); row <= e.Y2(); row++ {
		s.SetContent(e.X1(), row, tcell.RuneVLine, nil, e.BorderStyle.Style)
		s.SetContent(e.X2(), row, tcell.RuneVLine, nil, e.BorderStyle.Style)
	}

	if e.Y1() != e.Y2() && e.X1() != e.X2() {
		s.SetContent(e.X1(), e.Y1(), tcell.RuneULCorner, nil, e.BorderStyle.Style)
		s.SetContent(e.X2(), e.Y1(), tcell.RuneURCorner, nil, e.BorderStyle.Style)
		s.SetContent(e.X1(), e.Y2(), tcell.RuneLLCorner, nil, e.BorderStyle.Style)
		s.SetContent(e.X2(), e.Y2(), tcell.RuneLRCorner, nil, e.BorderStyle.Style)
	}
}

func (e *OldElement) drawTitle(s tcell.Screen) {
	row := e.Y1()
	col := e.X1() + e.TitlePadding.Left

	for _, r := range e.Title {
		s.SetContent(col, row, r, nil, e.TitleStyle.Style)
		col++

		if col > e.X2()-e.TitlePadding.Right {
			break
		}
	}
}

// Draw implements the Drawable interface
func (e *OldElement) Draw(s tcell.Screen) {
	e.drawBody(s)
	if e.Border {
		e.drawBorder(s)
	}
	e.drawTitle(s)
}

// SetRect implements the Drawable interface
func (e *OldElement) SetRect(x1 int, y1 int, x2 int, y2 int) {
	e.Rectangle = image.Rect(x1, y1, x2, y2)
	e.Inner = image.Rect(
		e.Min.X+e.Padding.Left,
		e.Min.Y+e.Padding.Top,
		e.Max.X-e.Padding.Right,
		e.Max.Y-e.Padding.Bottom,
	)
}

// X1 returns the rects Min.X
func (e *OldElement) X1() int {
	return e.Rectangle.Min.X
}

// X2 returns the rects Max.X
func (e *OldElement) X2() int {
	return e.Rectangle.Max.X
}

// Y1 returns the rects Min.X
func (e *OldElement) Y1() int {
	return e.Rectangle.Min.Y
}

// Y2 returns the rects Min.X
func (e *OldElement) Y2() int {
	return e.Rectangle.Max.Y
}

// Width ...
func (e *OldElement) DrawableWidth() int {
	return e.Inner.Max.X - e.Inner.Min.X
}

// DrawableHeight ...
func (e *OldElement) DrawableHeight() int {
	return e.Inner.Max.Y - e.Inner.Min.Y
}

// GetRect implements the Drawable interface
func (e *OldElement) GetRect() image.Rectangle {
	return e.Rectangle
}
