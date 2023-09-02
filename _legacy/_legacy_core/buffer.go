// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package tooey

import (
	"image"

	"github.com/mattn/go-runewidth"
	rw "github.com/mattn/go-runewidth"
)

// Cell represents a viewable terminal cell
type Cell struct {
	Rune  rune
	Style Style
}

// IsZeroWidth returns true if the rune is a zwc
func (c *Cell) IsZeroWidth() bool {
	w := c.Width()
	if w == 0 {
		return true
	} else {
		return false
	}
}

// Width returns the width of the cell's rune
func (c *Cell) Width() int {
	return runewidth.RuneWidth(c.Rune)
}

var CellClear = Cell{
	Rune:  ' ',
	Style: StyleClear,
}

// NewCell takes 1 to 2 arguments
// 1st argument = rune
// 2nd argument = optional style
func NewCell(rune rune, args ...interface{}) Cell {
	style := StyleClear
	if len(args) == 1 {
		style = args[0].(Style)
	}
	return Cell{
		Rune:  rune,
		Style: style,
	}
}

// Buffer represents a section of a terminal and is a renderable rectangle of cells.
type Buffer struct {
	image.Rectangle
	CellMap map[image.Point]Cell
}

func NewBuffer(r image.Rectangle) *Buffer {
	buf := &Buffer{
		Rectangle: r,
		CellMap:   make(map[image.Point]Cell),
	}
	buf.Fill(CellClear, r) // clears out area
	return buf
}

func (b *Buffer) GetCell(p image.Point) Cell {
	return b.CellMap[p]
}

func (b *Buffer) SetCell(c Cell, p image.Point) {
	b.CellMap[p] = c
}

func (b *Buffer) Fill(c Cell, rect image.Rectangle) {
	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			b.SetCell(c, image.Pt(x, y))
		}
	}
}

func (b *Buffer) SetString(s string, style Style, p image.Point) {
	runes := []rune(s)
	x := 0

	for i := 0; i < len(runes); i++ {
		char := runes[i]
		b.SetCell(Cell{char, style}, image.Pt(p.X+x, p.Y))
		x += rw.RuneWidth(char)
	}

}
