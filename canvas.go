package termui

import (
	"image"

	"github.com/asciifaceman/termui/drawille"
)

type Canvas struct {
	Block
	drawille.Canvas
}

func NewCanvas() *Canvas {
	return &Canvas{
		Block:  *NewBlock(),
		Canvas: *drawille.NewCanvas(),
	}
}

func (c *Canvas) SetPoint(p image.Point, color Color) {
	c.Canvas.SetPoint(p, drawille.Color(color))
}

func (c *Canvas) SetLine(p0, p1 image.Point, color Color) {
	c.Canvas.SetLine(p0, p1, drawille.Color(color))
}

func (c *Canvas) Draw(buf *Buffer) {
	for point, cell := range c.Canvas.GetCells() {
		if point.In(c.Rectangle) {
			convertedCell := Cell{
				cell.Rune,
				Style{
					Color(cell.Color),
					ColorClear,
					ModifierClear,
				},
			}
			buf.SetCell(convertedCell, point)
		}
	}
}
