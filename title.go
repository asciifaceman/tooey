package tooey

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
)

// NewTitle returns a basic empty title
func NewTitle(theme *Theme) *Title {
	if theme == nil {
		theme = DefaultTheme
	}

	return &Title{
		Padding: NewTitlePadding(),
		Theme:   theme,
	}
}

// Title represents a rendered title in the header of an Element
type Title struct {
	Content string
	Padding *Padding
	Theme   *Theme
}

// Draw the title
func (t *Title) Draw(s tcell.Screen, rect *Rectangle) {
	if len(t.Content) == 0 {
		return
	}

	w := rect.DrawableWidth()
	//draw := TrimString(t.Content, w-1)

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

	draw := TrimString(t.Content, w-len(leftPad)-len(rightPad)-1)

	draw = fmt.Sprintf("%s%s%s", leftPad, draw, rightPad)

	for _, r := range draw {
		s.SetContent(col, row, r, nil, t.Theme.Title.Style)
		col++

		if col+1 > rect.X2()-rect.Padding.Right {
			// add ... at some point
			break
		}
	}
}
