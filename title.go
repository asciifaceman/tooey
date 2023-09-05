package tooey

import (
	"github.com/gdamore/tcell/v2"
)

// NewTitle returns a basic empty title
func NewTitle() *Title {
	return &Title{
		Padding: NewTitlePadding(),
		Theme:   DefaultTheme,
	}
}

// Title represents a rendered title in the header of an Element
// TODO: Alignment (L / R / C)
type Title struct {
	Content string
	Padding *Padding
	Theme   Theme
}

// SetTheme sets the title's theme
func (t *Title) SetTheme(theme Theme) {
	t.Theme = theme
}

// Set sets the title's content
func (t *Title) Set(content string) {
	t.Content = content
}

/*
Draw the title to the given rect's dimensions within the screen
*/
func (t *Title) Draw(s tcell.Screen, rect *Rectangle) {
	contentLength := len(t.Content)

	if contentLength == 0 {
		return
	}

	row := rect.Y1()      // draw on rect's top most line
	col := rect.InnerX1() // don't draw before rect's padded start point
	end := rect.InnerX2() // don't draw past rect's padded end point

	available := end - col

	// Pad and trim
	draw := TrimStringWithPadding(t.Content, available-t.Padding.TotalWidth(), t.Padding.Left, t.Padding.Right)

	for _, r := range draw {
		s.SetContent(col, row, r, nil, GetCellStyle(t.Theme, t).Style)
		col++
	}

}
