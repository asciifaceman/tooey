package tooey

import "github.com/gdamore/tcell/v2"

// NewStyle returns a blank Style
func NewStyle() *Style {
	return WrapStyle(tcell.StyleDefault)
}

// WrapStyle returns a Tooey Style with the given tcell.Style
func WrapStyle(s tcell.Style) *Style {
	return &Style{
		Style: s,
	}
}

// AppendStyle accepts a

// Basic terminal colors
const (
	ColorClear   tcell.Color = tcell.ColorDefault
	ColorBlack   tcell.Color = tcell.ColorBlack
	ColorRed     tcell.Color = tcell.ColorRed
	ColorGreen   tcell.Color = tcell.ColorGreen
	ColorYellow  tcell.Color = tcell.ColorYellow
	ColorBlue    tcell.Color = tcell.ColorBlue
	ColorMagenta tcell.Color = tcell.ColorDarkMagenta
	ColorCyan    tcell.Color = tcell.ColorLightCyan
	ColorWhite   tcell.Color = tcell.ColorWhite
	ColorOrange  tcell.Color = tcell.ColorOrange
)

// Style represents the style of one terminal cell
// and contains a tcell.Style
type Style struct {
	tcell.Style
}

// StyleDefault represents a simple white on black default
var StyleDefault = NewStyle()

// StyleClassicTerminal is a classic green-on-black terminal styling
var StyleClassicTerminal = WrapStyle(NewStyle().Foreground(tcell.ColorLimeGreen).Background(tcell.ColorBlack))

// StyleSoftClassicTerminal is a classic green-on-black with a softer green
var StyleSoftClassicTerminal = WrapStyle(NewStyle().Foreground(tcell.ColorLightGreen).Background(tcell.ColorBlack))
