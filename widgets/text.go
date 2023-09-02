// TCell implementation of paragraph
//
// Charles <asciifaceman> Corbett 2023
//
//

package widgets

import (
	"image"

	"github.com/asciifaceman/tooey"
	"github.com/gdamore/tcell/v2"
)

// NewText returns a basic empty *Text
func NewText() *Text {
	return &Text{
		Element: *tooey.NewElement(),
		Wrap:    true,
	}
}

// Text represents a simple styled block of text with wrapping
// paragraph capabilities
type Text struct {
	tooey.Element
	Content string
	Theme   *tooey.Theme
	Wrap    bool
}

func (t *Text) SetTheme(theme *tooey.Theme) {
	t.Theme = theme
	t.Element.SetTheme(theme)
}

// Draw ...
func (t *Text) Draw(s tcell.Screen) {
	t.Element.Draw(s)

	row := t.Rectangle.Min.Y + t.Padding.Top
	col := t.Rectangle.Min.X + t.Padding.Left

	for _, r := range t.Content {

		// TODO: Handle zero width characters

		s.SetContent(col, row, r, nil, t.Theme.Title.Style)
		col++

		if t.Wrap {
			if col > t.Rectangle.Max.X-t.Padding.Right {
				row++
				col = t.Rectangle.Min.X + t.Padding.Left
			}
			if row > t.Rectangle.Max.Y {
				// gobble the remainder
				break
			}
		}
	}

}

func (t *Text) DrawAsChild(s tcell.Screen, rect image.Rectangle) {
	t.Element.Draw(s)

	// maybe this makes more sense if we go back to the drawing
	// board with Drawable and Element and have
	// the lowest level "Element" naturally be a container
	// of other elements and Drawing is inherently aware of
	// the parent's inner rect
	// need grid and flex style layout
	//
}
