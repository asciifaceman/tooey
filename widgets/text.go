// TCell implementation of paragraph
//
// Charles <asciifaceman> Corbett 2023
//
//

package widgets

import (
	"github.com/asciifaceman/tooey"
	"github.com/gdamore/tcell/v2"
)

// TextTheme extends RootTheme to add a Text compatible field
type TextTheme struct {
	tooey.RootTheme
	Text *tooey.Style
}

var DefaultTextTheme = &TextTheme{
	RootTheme: *tooey.DefaultTheme,
	Text:      tooey.StyleDefault,
}

// NewText returns a basic empty *Text with default theme and word wrapping
func NewText(theme tooey.Theme) *Text {
	if theme == nil {
		theme = tooey.DefaultTheme
	}

	return &Text{
		Element: *tooey.NewElement(),
		wrap:    true,
		Theme:   theme,
	}
}

// Text represents a simple styled block of text with wrapping
// paragraph capabilities
type Text struct {
	tooey.Element
	Content string
	Theme   tooey.Theme
	wrap    bool
}

// SetTheme sets the theme for the Text and its underlying Element
func (t *Text) SetTheme(theme *tooey.Theme) {
	t.Theme = theme
	t.Element.SetTheme(theme)
}

// SetWordWrap sets whether Text should wrap (true) or truncate (false)
func (t *Text) SetWordWrap(wrap bool) {
	t.wrap = wrap
}

// Draw ...
func (t *Text) Draw(s tcell.Screen) {
	t.Element.Draw(s)

	// produce [][]rune by processing newlines first, then word wrapping
	// then rslice.Left/Right or Normalize for alignment

}
