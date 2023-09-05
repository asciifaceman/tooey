//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/gdamore/tcell/v2"
)

type Border struct{}
type Title struct{}
type None struct{}

type Theme interface {
	GetCellStyle(interface{}) tcell.Style
}

func NewTheme() *RootTheme {
	return &RootTheme{
		Default: tcell.StyleDefault,
		Border:  tcell.StyleDefault.Foreground(tcell.ColorWheat),
		Title:   tcell.StyleDefault,
	}
}

type RootTheme struct {
	Default tcell.Style
	Border  tcell.Style
	Title   tcell.Style
}

type NewObject struct{}

type TestTheme struct {
	RootTheme
	NewObject tcell.Style
}

// GetCellStyle returns a tcell.Style that matches sub if it exists within Theme
// else returns tcell.StyleDefault
func GetCellStyle(r Theme, sub interface{}) tcell.Style {

	ty := reflect.TypeOf(sub)
	rv := reflect.ValueOf(r)

	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}

	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	resp := rv.FieldByName(ty.Name())

	switch resp.Kind() {
	case reflect.Struct:
		return resp.Interface().(tcell.Style)
	default:
		return tcell.StyleDefault
	}
}

func (r *RootTheme) GetCellStyle(sub interface{}) tcell.Style {
	ty := reflect.TypeOf(sub)
	rv := reflect.ValueOf(*r)

	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}

	resp := rv.FieldByName(ty.Name())

	switch resp.Kind() {
	case reflect.Struct:
		return resp.Interface().(tcell.Style)
	default:
		return tcell.StyleDefault
	}

}

func main() {
	r := NewTheme()

	b := &Border{}

	derp := r.GetCellStyle(b)

	spew.Dump(derp)

	fmt.Println("====")

	r2 := &TestTheme{
		RootTheme: *NewTheme(),
		NewObject: tcell.StyleDefault.Foreground(tcell.ColorDarkMagenta),
	}

	b2 := &NewObject{}

	derp2 := GetCellStyle(r2, b2)

	spew.Dump(r2.GetCellStyle(b))
	spew.Dump(derp2)
}
