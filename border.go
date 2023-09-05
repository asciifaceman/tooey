package tooey

import "github.com/gdamore/tcell/v2"

// BorderChars defines a characterset used for drawing borders
type BorderChars struct {
	H  rune
	V  rune
	UL rune
	UR rune
	LL rune
	LR rune
}

// DefaultBorderChars is the default character set for element borders
var DefaultBorderChars = BorderChars{
	H:  DefaultH,
	V:  DefaultV,
	UL: DefaultUL,
	UR: DefaultUR,
	LL: DefaultLL,
	LR: DefaultLR,
}

// RoundedBarBorderChars is like the Default border but the
// corners are rounded
var RoundedBarBorderChars = BorderChars{
	H:  DefaultH,
	V:  DefaultV,
	UL: '╭',
	UR: '╮',
	LL: '╰',
	LR: '╯',
}

// CornersOnlyBorderChars ...
var CornersOnlyBorderChars = BorderChars{
	H:  ' ',
	V:  ' ',
	UL: '◤',
	UR: '◥',
	LL: '◣',
	LR: '◢',
}

// NewBorder returns a border with the default character set
func NewBorder() *Border {
	return &Border{
		Enabled: true,
		Theme:   DefaultTheme,
		Chars:   DefaultBorderChars,
		Left:    true,
		Top:     true,
		Right:   true,
		Bottom:  true,
	}
}

// Border is a sub-element which contains the definition
// and drawing logic of an element border
type Border struct {
	Enabled bool
	Theme   Theme
	Chars   BorderChars

	Left   bool
	Top    bool
	Right  bool
	Bottom bool
}

func (b *Border) SetTheme(theme Theme) {
	b.Theme = theme
}

// SetChars sets the char map for borders
func (b *Border) SetChars(chars BorderChars) {
	b.Chars = chars
}

// Draw the borders for the given rect to the given tcell.Screen
func (b *Border) Draw(s tcell.Screen, rect *Rectangle) {

	if b.Enabled && !rect.ZeroSize() {
		for col := rect.X1(); col <= rect.X2(); col++ {

			if b.Top {
				s.SetContent(col, rect.Y1(), b.Chars.H, nil, GetCellStyle(b.Theme, b).Style)
			}

			if b.Bottom {
				s.SetContent(col, rect.Y2(), b.Chars.H, nil, GetCellStyle(b.Theme, b).Style)
			}

		}

		for row := rect.Y1(); row <= rect.Y2(); row++ {

			if b.Left {
				s.SetContent(rect.X1(), row, b.Chars.V, nil, GetCellStyle(b.Theme, b).Style)
			}

			if b.Right {
				s.SetContent(rect.X2(), row, b.Chars.V, nil, GetCellStyle(b.Theme, b).Style)
			}
		}

		// Patch corners as necessary
		if !rect.ZeroSize() {
			if b.Top && b.Left {
				s.SetContent(rect.X1(), rect.Y1(), b.Chars.UL, nil, GetCellStyle(b.Theme, b).Style)
			}
			if b.Top && b.Right {
				s.SetContent(rect.X2(), rect.Y1(), b.Chars.UR, nil, GetCellStyle(b.Theme, b).Style)
			}
			if b.Left && b.Bottom {
				s.SetContent(rect.X1(), rect.Y2(), b.Chars.LL, nil, GetCellStyle(b.Theme, b).Style)
			}
			if b.Bottom && b.Right {
				s.SetContent(rect.X2(), rect.Y2(), b.Chars.LR, nil, GetCellStyle(b.Theme, b).Style)
			}
		}

	}

}
