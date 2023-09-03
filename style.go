package tooey

import "github.com/gdamore/tcell/v2"

// Color is an integer from -1 to 255
// -1 = ColorClear
// 0-255 = Xterm colors
type Color tcell.Color

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
)

// Style represents the style of one terminal cell
type Style struct {
	tcell.Style
	Align Alignment
}

// StyleClear represents an empty Style, with no colors or modifiers
var StyleClear = Style{
	Style: tcell.StyleDefault,
}

// StyleDefault represents a simple white on black default
var StyleDefault = Style{
	Style: tcell.StyleDefault.Foreground(ColorWhite).Background(ColorBlack),
	Align: AlignFull,
}

// StyleClassicTerminal is a classic green-on-black terminal styling
var StyleClassicTerminal = Style{
	Style: tcell.StyleDefault.Foreground(tcell.ColorLightGreen).Background(tcell.ColorBlack),
	Align: AlignFull,
}
