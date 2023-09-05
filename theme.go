package tooey

import (
	"fmt"
	"reflect"

	"github.com/davecgh/go-spew/spew"
)

/*
Theme is an interface that allows for an extensible theming structure wherein
the RootTheme can be extended with new SubElements or Widgets and allow a user
to still acquire styles at Draw time
*/
type Theme interface {
}

// RootTheme is the parent theme all other themes extend
type RootTheme struct {
	Element *Style
	Border  *Style
	Title   *Style
}

// GetCellStyle returns a *Style that matches sub's tyle if it exists within
// the given Theme else it panics
func GetCellStyle(t Theme, sub interface{}) Style {
	ts := reflect.TypeOf(sub)
	tv := reflect.ValueOf(t)

	if ts.Kind() == reflect.Ptr {
		ts = ts.Elem()
	}

	if tv.Kind() == reflect.Ptr {
		tv = tv.Elem()
	}

	field := tv.FieldByName(ts.Name())
	if field.Kind() == reflect.Ptr {
		field = field.Elem()
	}

	switch field.Kind() {
	case reflect.Struct:
		return field.Interface().(Style)
	default:
		// TODO: Should we panic here or return a default style
		panic(fmt.Sprintf("The given object [%v] does not have a style registered in the given theme [%v] [%v]", ts.Name(), spew.Sdump(t), spew.Sdump(field)))
	}
}

var DefaultTheme = &RootTheme{
	Element: StyleDefault,
	Border:  StyleDefault,
	Title:   StyleDefault,
}

var ClassicTheme = &RootTheme{
	Element: StyleClassicTerminal,
	Border:  StyleClassicTerminal,
	Title:   StyleClassicTerminal,
}
