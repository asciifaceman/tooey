package tooey

import "github.com/gdamore/tcell/v2"

/*
SubElement represents an element that is a part of a parent Element and does
not contain its own *Rectangle but instead draws relative to its parent
and accept a Theme

NOT YET IMPLEMENTED
*/
type SubElement interface {
	SetTheme(theme Theme)
	Draw(tcell.Screen, *Rectangle)
}
