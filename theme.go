package tooey

import (
	"github.com/gdamore/tcell/v2"
)

const (
	DefaultULCorner = tcell.RuneULCorner
	DefaultURCorner = tcell.RuneURCorner
	DefaultLLCorner = tcell.RuneLLCorner
	DefaultLRCorner = tcell.RuneLRCorner
	DefaultHLine    = tcell.RuneHLine
	DefaultVLine    = tcell.RuneVLine
)

// Chars is to enable theming border characters
type Chars struct {
	HLine    rune
	VLine    rune
	ULCorner rune
	URCorner rune
	LLCorner rune
	LRCorner rune
}

// NewDefaultChars returns the default character set
// for borders
func NewDefaultChars() *Chars {
	return &Chars{
		HLine:    DefaultHLine,
		VLine:    DefaultVLine,
		ULCorner: DefaultULCorner,
		URCorner: DefaultURCorner,
		LLCorner: DefaultLLCorner,
		LRCorner: DefaultLRCorner,
	}
}

var DefaultStylized = &Chars{
	HLine:    DefaultHLine,
	VLine:    DefaultVLine,
	ULCorner: '╒',
	URCorner: DefaultURCorner,
	LLCorner: DefaultLLCorner,
	LRCorner: DefaultLRCorner,
}

// RoundedBarBorderChars is like the Default border but the
// corners are rounded
var RoundedBarBorderChars = &Chars{
	HLine:    DefaultHLine,
	VLine:    DefaultVLine,
	ULCorner: '╭',
	URCorner: '╮',
	LLCorner: '╰',
	LRCorner: '╯',
}

// CornersOnlyBorderChars will only render corners
// with triangle ascii
var CornersOnlyBorderChars = &Chars{
	HLine:    ' ',
	VLine:    ' ',
	ULCorner: '◤',
	URCorner: '◥',
	LLCorner: '◣',
	LRCorner: '◢',
}

// DoubleBarBorder is like the Default border but with
// double bars for more dramatic effect // rune(2550)
var DoubleBarBorderChars = &Chars{
	HLine:    '═',
	VLine:    '║',
	ULCorner: '╔',
	URCorner: '╗',
	LLCorner: '╚',
	LRCorner: '╝',
}

/*
Theme is a bundle of Styles for different elements, subelements, and widgets
If Inherit is true in the theme, then when a theme is set it will propagate the Default
down to the others
*/
type Theme struct {
	Default Style
	Element Style
	Border  Style
	Title   Style
	Text    Style
	Chars   *Chars
}

// DefaultTheme is a basic white foreground and black background for all elements
var DefaultTheme = &Theme{
	Default: StyleDefault,
	Element: StyleClear,
	Border:  StyleDefault,
	Title:   StyleDefault,
	Text:    StyleDefault,
	Chars:   NewDefaultChars(),
}
