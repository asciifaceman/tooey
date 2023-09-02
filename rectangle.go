package tooey

import (
	"image"
	"sync"
)

func NewRectangle(padding *Padding) Rectangle {
	if padding == nil {
		padding = NewDefaultPadding()
	}

	return Rectangle{
		Padding: padding,
	}
}

// Rectangle is a convenience extension of the stdlib
// image.Rectangle
type Rectangle struct {
	image.Rectangle
	// Padding should only affect objects within a rectangle
	// not the outer bounds of the rectangle
	Padding *Padding

	sync.Mutex
}

// SetRect defines the boundaries of the rectangle
func (r *Rectangle) SetRect(x1 int, y1 int, x2 int, y2 int) {
	r.Rectangle = image.Rect(x1, y1, x2, y2)
}

// GetRect returns the current underlying image.Rectangle
func (r *Rectangle) GetRect() image.Rectangle {
	return r.Rectangle
}

// GetInnerRect returns the bounds of the inner padded rectangle
func (r *Rectangle) GetInnerRect() image.Rectangle {
	return image.Rect(r.InnerX1(), r.InnerY1(), r.InnerX2(), r.InnerY2())
}

// X1 returns the rectangle's Min.X point
func (r *Rectangle) X1() int {
	return r.Min.X
}

// X2 returns the rectangle's Max.X point
func (r *Rectangle) X2() int {
	return r.Max.X
}

// Y1 returns the rectangle's Min.Y point
func (r *Rectangle) Y1() int {
	return r.Min.Y
}

// Y2 returns the rectangle's Max.Y point
func (r *Rectangle) Y2() int {
	return r.Max.Y
}

// InnerX1 returns X1 with padding
func (r *Rectangle) InnerX1() int {
	return r.X1() + r.Padding.Left
}

// InnerX2 returns X2 with padding
func (r *Rectangle) InnerX2() int {
	return r.X2() - r.Padding.Right
}

// InnerY1 returns Y1 with padding
func (r *Rectangle) InnerY1() int {
	return r.Y1() + r.Padding.Top
}

// InnerY2 returns Y2 with padding
func (r *Rectangle) InnerY2() int {
	return r.Y2() - r.Padding.Bottom
}

// DrawableWidth returns the max width of the rectangle minus padding
func (r *Rectangle) DrawableWidth() int {
	return r.X2() - r.X1() - r.Padding.Left - r.Padding.Right
}

// DrawableHeight returns the max height of the rectangle minux padding
func (r *Rectangle) DrawableHeight() int {
	return r.Y2() - r.Y1() - r.Padding.Top - r.Padding.Bottom
}

// ZeroSize returns true if the rectangle has no size
func (r *Rectangle) ZeroSize() bool {
	if r.Y1() == r.Y2() && r.X1() == r.X2() {
		return true
	} else {
		return false
	}
}
