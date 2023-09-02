package tooey

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
)

// NewTitle returns a basic empty title
func NewTitle() *Title {
	return &Title{
		Padding: NewDefaultPadding(),
		Style:   StyleDefault,
	}
}

// Title represents a rendered title in the header of an Element
type Title struct {
	Content string
	Padding *Padding
	Style   Style
}

// Draw the title
func (t *Title) Draw(s tcell.Screen, rect *Rectangle) {
	if len(t.Content) == 0 {
		return
	}

	row := rect.Y1()
	col := rect.X1() + rect.Padding.Left

	leftPad := ""
	if t.Padding.Left > 0 {
		leftPad = strings.Repeat(" ", t.Padding.Left)
	}
	rightPad := ""
	if t.Padding.Right > 0 {
		rightPad = strings.Repeat(" ", t.Padding.Right)
	}

	t.Content = fmt.Sprintf("%s%s%s", leftPad, t.Content, rightPad)

	for _, r := range t.Content {
		s.SetContent(col, row, r, nil, t.Style.Style)
		col++

		if col > rect.X2()-rect.Padding.Right {
			// add ... at some point
			break
		}
	}
}
