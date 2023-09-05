// TCell implementation of paragraph
//
// Charles <asciifaceman> Corbett 2023
//
//

package widgets

import (
	"fmt"
	"image"
	"unicode"

	"github.com/asciifaceman/tooey"
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

// NewText returns a basic empty *Text
func NewText(theme *tooey.Theme) *Text {
	if theme == nil {
		theme = tooey.DefaultTheme
	}

	return &Text{
		Element: *tooey.NewElement(theme),
		Wrap:    true,
		Theme:   theme,
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

// SetTheme sets the theme for the Text and it's underlying Element
func (t *Text) SetTheme(theme *tooey.Theme) {
	t.Theme = theme
	t.Element.SetTheme(theme)
}

// DrawableRect returns a set of ints representing
// the calculated drawable space based on all padding
// and positioning
//
// Returns image.Rectangle
func (t *Text) DrawableRect() image.Rectangle {
	// Calculate rect for drawable text
	minX1 := t.InnerX1() + t.Padding.Left
	minX2 := t.InnerX2() // - t.Padding.Right
	minY1 := t.InnerY1() + t.Padding.Top
	minY2 := t.InnerY2() // - t.Padding.Bottom

	contentRect := image.Rect(minX1, minY1, minX2, minY2)
	return contentRect
}

// Draw draws the text to its given Rect taking into account
// left and right padding, wrapping, etc
func (t *Text) Draw(s tcell.Screen) {
	t.Element.Draw(s)

	sw, sh := tooey.DrawableDimensions()

	contentCellLength := len([]rune(t.Content))
	if contentCellLength == 0 {
		// If the string is empty let's exit
		// early and save cycles
		return
	}

	contentRect := t.DrawableRect()

	if t.Wrap {
		alignedContent := make([][]rune, sh)

		for y := 0; y < sh; y++ {
			alignedContent[y] = make([]rune, sw)
		}

		switch t.Theme.Text.Align {
		case tooey.AlignLeft:
			alignedContent = t.ProcessLeftAlignment(contentRect)
		case tooey.AlignCenter:
			alignedContent = t.ProcessCenterAlignment(contentRect)
		case tooey.AlignRight:
			alignedContent = t.ProcessRightAlignment(contentRect)
		case tooey.AlignFull:
			alignedContent = t.ProcessFullAlignment(contentRect)
		}

		// aligned content is a precalculated grid so
		// we don't need to worry about placement logic here
		for y := 0; y < contentRect.Dy(); y++ {
			for x := 0; x < contentRect.Dx(); x++ {
				draw := alignedContent[y][x]

				// handle zero width chars
				var comb []rune
				w := runewidth.RuneWidth(draw)
				if w == 0 {
					comb = []rune{draw}
					draw = ' '
					w = 1
				}
				s.SetContent(x+contentRect.Min.X, y+contentRect.Min.Y, draw, comb, t.Theme.Text.Style)
			}
		}
	} else {

		processed := t.ProcessUnwrapped(contentRect.Min.X, sw-contentRect.Min.X)
		y := contentRect.Min.Y
		for x := contentRect.Min.X; x < sw; x++ {
			draw := processed[x]
			// handle zero width chars
			var comb []rune
			w := runewidth.RuneWidth(draw)
			if w == 0 {
				comb = []rune{draw}
				draw = ' '
				w = 1
			}
			s.SetContent(x, y, draw, comb, t.Theme.Text.Style)
		}

	}

}

// FillRuneBuffer prefills a buffer of size image.Rectangle for
// filling drawable text areas
//
// This is unfortunately destructive and prevents things like overlapping
// which may or may not be desirable but it was the way I could think of
// to prepare a buffer for managing alignment of text within a text space
func (t *Text) FillRuneBuffer(r image.Rectangle) [][]rune {
	buf2 := make([][]rune, r.Dy())

	for y := 0; y < r.Dy(); y++ {
		buf2[y] = make([]rune, r.Dx())
	}

	for y := 0; y < r.Dy(); y++ {
		for x := 0; x < r.Dx(); x++ {
			buf2[y][x] = rune(' ')
		}
	}

	return buf2
}

// ProcessLeftAlignment accepts a space of image.Rectangle and processes
// Text.Content within that space to be left justified, and attempt to
// wrap word-aware when possible. If the word will fit on a line by itself
// it will wrap, otherwise it will wrap mid-word.
//
// returns [y][x]rune
func (t *Text) ProcessLeftAlignment(r image.Rectangle) [][]rune {

	content := []rune(t.Content)
	contentLength := len(content)
	processed := t.FillRuneBuffer(r)

	// if the entire content's length is less than the width of
	// the image.Rectangle just spew it really quick and return early
	if contentLength < r.Dx() {
		for x := 0; x < contentLength; x++ {
			processed[0][x] = content[x]
		}
		return processed
	}

	offset := 0
	var previousRune rune

	for y := 0; y < r.Dy(); y++ {

		for x := 0; x < r.Dx(); x++ {

			// Attempt to wrap early if a word won't fit
			// but only if it fits in the drawable width to start with
			if !unicode.IsSpace(content[offset]) {
				if unicode.IsSpace(previousRune) {
					var word []rune
					word = append(word, content[offset])
					for i := offset + 1; i < contentLength; i++ {
						if unicode.IsSpace(content[i]) {
							break
						}
						word = append(word, content[i])
					}
					if len(word) > (r.Dx()-x) && len(word) < (r.Dx()) {
						break
					}
				}
			}

			processed[y][x] = content[offset]
			previousRune = content[offset]
			offset++

			if offset == contentLength {
				return processed
			}

		}
		block := processed[y]
		block = tooey.ShiftRuneWhitespaceToRight(block)
		processed[y] = block
	}

	return processed
}

// ProcessRightAlignment accepts a space of image.Rectangle and processes
// Text.Content within that space to be right justified and attempt to
// wrap word-aware when possible. If the word will fit on a line by itself
// it will wrap, otherwise it will wrap mid-word
//
// returns [y][x]rune
func (t *Text) ProcessRightAlignment(r image.Rectangle) [][]rune {
	leftAligned := t.ProcessLeftAlignment(r)

	for y := 0; y < r.Dy(); y++ {
		block := leftAligned[y]
		right := tooey.ShiftRuneWhitespaceToLeft(block)
		leftAligned[y] = right
	}
	return leftAligned
}

// ProcessCenterAlignment ...
func (t *Text) ProcessCenterAlignment(r image.Rectangle) [][]rune {
	leftAligned := t.ProcessLeftAlignment(r)

	for y := 0; y < r.Dy(); y++ {
		block := leftAligned[y]
		full := tooey.SpreadWhitespaceAcrossSliceInterior(block)
		leftAligned[y] = full
	}

	return leftAligned
}

// ProcessFullAlignment ...
func (t *Text) ProcessFullAlignment(r image.Rectangle) [][]rune {

	leftAligned := t.ProcessLeftAlignment(r)

	for y := 0; y < r.Dy(); y++ {
		block := leftAligned[y]
		full := tooey.SpreadWhitespaceAcrossSliceInterior(block)
		leftAligned[y] = full
	}

	return leftAligned
}

// Draw ...
func (t *Text) Draw2(s tcell.Screen) {
	t.Element.Draw(s)

	row := t.Rectangle.Min.Y + t.Padding.Top
	col := t.Rectangle.Min.X + t.Padding.Left

	wrapped := "&"
	previousRune := rune(wrapped[0])
	var currentWord string

	for i, r := range t.Content {

		// Lookahead
		if !unicode.IsSpace(r) {
			if unicode.IsSpace(previousRune) {
			INNER:
				for ix := i; i < len(t.Content); i++ {
					runeIter := rune(t.Content[ix])
					if unicode.IsSpace(runeIter) {
						break INNER
					}
					currentWord = fmt.Sprintf("%s%v", currentWord, runeIter)
				}

				// look ahead to find word
				if len([]rune(currentWord)) > (t.Rectangle.X2()-t.Padding.Right)-col {
					wrapped := "&"
					previousRune = rune(wrapped[0])
					row++
					col = t.Rectangle.X1() + t.Padding.Left
					continue
				}

			}
		} else {
			// if alignment is full make sure we aren't starting a new line on a space
			if t.Theme.Text.Align == tooey.AlignFull {
				previousRune = r
				continue
			}
			// see if I am at the beginning of the width and
			// skip spaces
		}

		// if character is a newline advance row and continue
		if fmt.Sprintf("%v", r) == "\n" {
			previousRune = r
			row++
			col = t.Rectangle.Min.X + t.Padding.Left
			continue
		}

		// Write the cell
		s.SetContent(col, row, r, nil, t.Theme.Title.Style)
		col++

		previousRune = r

		if t.Wrap {
			if col > t.Rectangle.Max.X-t.Padding.Right {
				// Wordwrap
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

// ProcessUnwrapped needs rewritten...
//
// returns [y][x]rune
func (t *Text) ProcessUnwrapped(x int, maxWidth int) map[int]rune {
	draw := t.Content
	if runewidth.StringWidth(t.Content) > maxWidth {
		draw = runewidth.Truncate(t.Content, maxWidth, string(tooey.ELLIPSES))

	}

	processed := map[int]rune{}

	col := x
	for _, r := range draw {

		processed[x] = r
		col++
	}

	return processed
}
