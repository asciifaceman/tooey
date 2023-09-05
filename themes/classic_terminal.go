package themes

import (
	"github.com/asciifaceman/tooey"
	"github.com/gdamore/tcell/v2"
)

var StyleRetroTerminalOrange = tooey.Style{
	Style: tcell.StyleDefault.Foreground(tcell.ColorDarkOrange).Background(tcell.ColorBlack),
}

var ThemeRetroTerminalOrange = &tooey.RootTheme{
	Element: &StyleRetroTerminalOrange,
	Border:  &StyleRetroTerminalOrange,
	Title:   &StyleRetroTerminalOrange,
}

var StyleRetroTerminalGreen = tooey.Style{
	Style: tcell.StyleDefault.Foreground(tcell.ColorDarkGreen).Background(tcell.ColorBlack),
}

/*
var StyleRetroOrangeTerminal = tooey.Style{
	Style: tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorOrange),
}

var StyleRetroGreenTerminal = tooey.Style{
	Style: tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorGreen),
}

var ThemeRetroTerminalOrange = &tooey.Theme{
	Default: StyleRetroOrangeTerminal,
	Element: StyleRetroOrangeTerminal,
	Border:  StyleRetroOrangeTerminal,
	Title:   StyleRetroOrangeTerminal,
	Chars:   tooey.NewDefaultChars(),
}

var ThemeRetroTerminalGreen = &tooey.Theme{
	Default: StyleRetroGreenTerminal,
	Element: StyleRetroGreenTerminal,
	Border:  StyleRetroGreenTerminal,
	Title:   StyleRetroGreenTerminal,
	Chars:   tooey.NewDefaultChars(),
}
*/
