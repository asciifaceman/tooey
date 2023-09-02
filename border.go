package tooey

import "github.com/gdamore/tcell/v2"

// NewDefaultBorder returns a border with the default character set
func NewDefaultBorder(theme *Theme) *Border {
	if theme == nil {
		theme = DefaultTheme
	}

	return &Border{
		Enabled: true,
		Theme:   theme,
		Chars:   theme.Chars,
		Left:    true,
		Top:     true,
		Right:   true,
		Bottom:  true,
	}
}

// Border contains the definition and drawing logic
// of an element border
type Border struct {
	Enabled bool
	Theme   *Theme
	Chars   *Chars

	Left   bool
	Top    bool
	Right  bool
	Bottom bool
}

// SetChars sets the char map for borders
func (b *Border) SetChars(chars *Chars) {
	b.Chars = chars
}

// Draw the borders for the given rect to the given tcell.Screen
func (b *Border) Draw(s tcell.Screen, rect *Rectangle) {

	if b.Enabled {
		for col := rect.X1(); col <= rect.X2(); col++ {

			if b.Top {
				s.SetContent(col, rect.Y1(), b.Chars.HLine, nil, b.Theme.Border.Style)
			}

			if b.Bottom {
				s.SetContent(col, rect.Y2(), b.Chars.HLine, nil, b.Theme.Border.Style)
			}

		}

		for row := rect.Y1(); row <= rect.Y2(); row++ {

			if b.Left {
				s.SetContent(rect.X1(), row, b.Chars.VLine, nil, b.Theme.Border.Style)
			}

			if b.Right {
				s.SetContent(rect.X2(), row, b.Chars.VLine, nil, b.Theme.Border.Style)
			}
		}

		// Patch corners as necessary
		if !rect.ZeroSize() {
			if b.Top && b.Left {
				s.SetContent(rect.X1(), rect.Y1(), b.Chars.ULCorner, nil, b.Theme.Border.Style)
			}
			if b.Top && b.Right {
				s.SetContent(rect.X2(), rect.Y1(), b.Chars.URCorner, nil, b.Theme.Border.Style)
			}
			if b.Left && b.Bottom {
				s.SetContent(rect.X1(), rect.Y2(), b.Chars.LLCorner, nil, b.Theme.Border.Style)
			}
			if b.Bottom && b.Right {
				s.SetContent(rect.X2(), rect.Y2(), b.Chars.LRCorner, nil, b.Theme.Border.Style)
			}
		}

	}

}
