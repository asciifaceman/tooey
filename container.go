package tooey

import (
	"github.com/gdamore/tcell/v2"
)

type FlexDirection uint

const (
	FlexColumn        FlexDirection = iota
	FlexColumnReverse               // NOT IMPLEMENTED
	FlexRow
	FlexRowReverse // NOT IMPLEMENTED
)

// NewContainer returns a new default configuration Container
//
// A theme can be passed at this time or nil for defaults or to configure later
func NewContainer() *Container {
	return &Container{
		Direction: FlexRow,
		Element:   *NewElement(),
		//Theme:     DefaultTheme,
	}
}

// Container is an element that holds other elements within
type Container struct {
	Element
	Direction FlexDirection
	Children  []ContainerChild
	Theme     *Theme
}

// ContainerChild which represents a member of the flex
// Grow is added up then divided by the total grow
// to find the ratio of space each child will consume
type ContainerChild struct {
	Drawable bool
	Contents interface{}
	Grow     float64
}

// NewFlexChild produces a ContainerChild which can be Wrapped
func NewFlexChild(grow float64, i ...interface{}) ContainerChild {
	_, ok := i[0].(Drawable)
	child := i[0]
	if !ok {
		child = i
	}

	return ContainerChild{
		Drawable: ok,
		Contents: child,
		Grow:     grow,
	}
}

//

// Wrap embeds the given objects within the container
// using a top-level container that will fill it's available space
func (c *Container) Wrap(children ...interface{}) {
	child := ContainerChild{
		Drawable: false,
		Contents: children,
		Grow:     1.0,
	}
	c.RecursiveWrap(child)
}

// RecursiveWrap wraps a tree of children
func (c *Container) RecursiveWrap(child ContainerChild) {

	if child.Drawable {
		c.Children = append(c.Children, child)
	} else {

		children := InterfaceSlice(child.Contents)

		for i := 0; i < len(children); i++ {
			if children[i] == nil {
				continue
			}
			ch, _ := children[i].(ContainerChild)
			c.RecursiveWrap(ch)
		}

	}

}

// DrawFlexRow will draw the contents as a flexible row
func (c *Container) DrawFlexRow(s tcell.Screen) {

	totalFlex := 0.0

	for _, child := range c.Children {
		totalFlex += child.Grow
	}

	width := float64(c.GetInnerRect().Dx())

	lastPosition := c.InnerX1()

	for _, child := range c.Children {
		childRatio := child.Grow / totalFlex // mult by available width
		childWidth := width * childRatio

		drawableChild := child.Contents.(Drawable)

		x := lastPosition
		y := c.InnerY1()
		w := int(childWidth)
		h := c.InnerY2()

		if x+w > c.GetInnerRect().Dx() {
			w--
		}

		drawableChild.SetRect(x, y, x+w, h)

		drawableChild.Lock()
		drawableChild.Draw(s)
		drawableChild.Unlock()

		lastPosition = x + w + 1

	}
}

// DrawFlexColumn will draw the contents as a flexible column
func (c *Container) DrawFlexColumn(s tcell.Screen) {
	totalFlex := c.calcFlex()

	height := float64(c.GetInnerRect().Dy())

	lastPosition := c.InnerY1()

	for _, child := range c.Children {
		childRatio := child.Grow / totalFlex
		childHeight := height * childRatio

		drawableChild := child.Contents.(Drawable)

		x := c.InnerX1()
		y := lastPosition
		w := c.InnerX2()
		h := int(childHeight)

		if y+h > c.GetInnerRect().Dy() {
			h--
		}

		drawableChild.SetRect(x, y, w, y+h)

		drawableChild.Lock()
		drawableChild.Draw(s)
		drawableChild.Unlock()

		lastPosition = y + h + 1
	}
}

// Draw draws the row or col flex and their children
func (c *Container) Draw(s tcell.Screen) {
	c.Element.Draw(s)

	switch c.Direction {
	case FlexColumn:
		c.DrawFlexColumn(s)
	case FlexColumnReverse:
		panic("FlexColumnReverse not yet implemented")
	case FlexRow:
		c.DrawFlexRow(s)
	case FlexRowReverse:
		panic("FlexRowReverse not yet implemented")
	default:
		panic("No flex direction selected")
	}

}

// calcFlex just adds up the flex Grows across the container's children
func (c *Container) calcFlex() float64 {
	totalFlex := 0.0

	for _, child := range c.Children {
		totalFlex += child.Grow
	}

	return totalFlex
}
