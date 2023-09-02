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
