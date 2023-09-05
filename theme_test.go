package tooey

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

type CustomObject struct{}

type NewThemeWithCustomObject struct {
	RootTheme
	CustomObject *Style
}

func TestGetCellStyle(t *testing.T) {

	test := GetCellStyle(DefaultTheme, &Element{})
	if test.Style != tcell.StyleDefault {
		t.Fatal("Expected default theme to utilize StyleDefault for Element")
	}

	testStyle := WrapStyle(tcell.StyleDefault.Foreground(tcell.ColorRed))

	th := &NewThemeWithCustomObject{
		RootTheme:    *DefaultTheme,
		CustomObject: testStyle,
	}

	test2 := GetCellStyle(th, &CustomObject{})

	if test2 != *testStyle {
		t.Fatal("Expected lookup on custom theme and custom object to return expected style")
	}
}
