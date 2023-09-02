package widgets

import (
	"github.com/asciifaceman/tooey"
	"github.com/gdamore/tcell/v2"
)

// Container is an Element which holds
// other Elements within its bounds and
// can have a styled border with a title
type Container struct {
	tooey.Element
	Contents []tooey.Drawable

	Border      bool
	BorderStyle tooey.Style

	BorderLeft   bool
	BorderRight  bool
	BorderTop    bool
	BorderBottom bool

	Title            string
	TitleStyle       tooey.Style
	TitlePaddingLeft int
}

// NewContainer ...
// TODO: Themeing redo
func NewContainer() *Container {
	return &Container{
		Border:           true,
		BorderStyle:      tooey.StyleDefault,
		BorderLeft:       true,
		BorderRight:      true,
		BorderTop:        true,
		BorderBottom:     true,
		TitleStyle:       tooey.StyleDefault,
		TitlePaddingLeft: 2,
	}
}

func (c *Container) Append(d tooey.Drawable) {
	// d.SetRect(c.X1()+d.X1()+c.PaddingLeft,
	//
	//	c.Y1()+d.Y1()-c.PaddingLeft,
	//	c.X1()+d.X2()+c.PaddingRight,
	//	c.Y2()+d.Y2()-c.PaddingBottom)
	//
	// c.Contents = append(c.Contents, d)
}

func (c *Container) drawBorder(s tcell.Screen) {

	// TODO: fill background / container style
	//for row := c.Rectangle.Min.Y; row <= c.Rectangle.Max.Y; row++ {
	//	for col := c.Rectangle.Min.X; col <= c.Rectangle.Max.X; col++ {
	//		s.SetContent(col, row, ' ', nil, c.BorderStyle.Style)
	//	}
	//}

	// draw borders
	for col := c.Rectangle.Min.X; col <= c.Rectangle.Max.X; col++ {
		s.SetContent(col, c.Rectangle.Min.Y, tcell.RuneHLine, nil, c.BorderStyle.Style)
		s.SetContent(col, c.Rectangle.Max.Y, tcell.RuneHLine, nil, c.BorderStyle.Style)
	}

	for row := c.Rectangle.Min.Y; row <= c.Rectangle.Max.Y; row++ {
		s.SetContent(c.Rectangle.Min.X, row, tcell.RuneVLine, nil, c.BorderStyle.Style)
		s.SetContent(c.Rectangle.Max.X, row, tcell.RuneVLine, nil, c.BorderStyle.Style)
	}

}

func (c *Container) drawTitle(s tcell.Screen) {
	row := c.Rectangle.Min.Y
	col := c.Rectangle.Min.X + c.TitlePaddingLeft

	for _, r := range c.Title {
		// TODO: Handle zero width characters

		s.SetContent(col, row, r, nil, c.TitleStyle.Style)
		col++

		// TODO: TitlePaddingRight
		if col > c.Rectangle.Max.X-c.TitlePaddingLeft {
			break
		}

		// draw corners if necessary
		if c.Rectangle.Min.Y != c.Rectangle.Max.Y && c.Rectangle.Min.X != c.Rectangle.Max.X {
			s.SetContent(c.X1(), c.Y1(), tcell.RuneULCorner, nil, c.BorderStyle.Style)
			s.SetContent(c.X2(), c.Y1(), tcell.RuneURCorner, nil, c.BorderStyle.Style)
			s.SetContent(c.X1(), c.Y2(), tcell.RuneLLCorner, nil, c.BorderStyle.Style)
			s.SetContent(c.X2(), c.Y2(), tcell.RuneLRCorner, nil, c.BorderStyle.Style)
		}
	}
}

func (c *Container) Draw(s tcell.Screen) {
	if c.Border {
		c.drawBorder(s)
	}

	c.drawTitle(s)

	for _, element := range c.Contents {
		element.Draw(s)
	}

}
