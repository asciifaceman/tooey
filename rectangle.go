package tooey

import (
	"image"
	"sync"
)

// NewRectangle returns a rectangle
func NewRectangle() Rectangle {
	return Rectangle{
		Padding: NewDefaultPadding(),
	}
}

// Rectangle is a convenience extension of the stdlib
// image.Rectangle which adds locking, interior padding, etc
//
// Objects cannot be drawn outside of a Rectangle, overflowing
// isn't supported in general right now to attempt to improve stability
type Rectangle struct {
	image.Rectangle
	// Padding should only affect objects within a rectangle
	// not the outer bounds of the rectangle which will fill
	// whatever the given rect is
	Padding *Padding

	sync.Mutex
}

// Width returns the outer width of the rectangle
func (r *Rectangle) Width() int {
	return r.Max.X - r.Min.X
}

// Height returns the outer Height of the rectangle
func (r *Rectangle) Height() int {
	return r.Max.Y - r.Min.Y
}

// Rect applies the given image.Rect to the Rectangle defining its bounds
func (r *Rectangle) Rect(rect image.Rectangle) {
	r.Rectangle = rect
}

// SetRect takes the given coordinates and creates an image.Rectangle which is
// applied to the given Rectangle
func (r *Rectangle) SetRect(x1 int, y1 int, x2 int, y2 int) {
	r.Rect(image.Rect(x1, y1, x2, y2))
}

// GetRect returns the current underlying image.Rectangle
func (r *Rectangle) GetRect() image.Rectangle {
	return r.Rectangle
}

// GetInnerRect returns the bounds of the inner padded rectangle
func (r *Rectangle) GetInnerRect() image.Rectangle {
	return image.Rect(r.InnerX1(), r.InnerY1(), r.InnerX2(), r.InnerY2())
}

// SetPadding sets the padding of the rectangle
func (r *Rectangle) SetPadding(padding *Padding) {
	r.Padding = padding
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
	if r.Width() > 0 {
		return r.X1() + r.Padding.Left
	}
	return r.X1()
}

// InnerX2 returns X2 with padding
func (r *Rectangle) InnerX2() int {
	if r.Width() > 0 {
		return r.X2() - r.Padding.Right
	}
	return r.X2()
}

// InnerY1 returns Y1 with padding
func (r *Rectangle) InnerY1() int {
	if r.Height() > 0 {
		return r.Y1() + r.Padding.Top
	}
	return r.Y1()
}

// InnerY2 returns Y2 with padding
func (r *Rectangle) InnerY2() int {
	if r.Height() > 0 {
		return r.Y2() - r.Padding.Bottom
	}
	return r.Y2()
}

// DrawableWidth returns the max width of the rectangle minus padding
func (r *Rectangle) DrawableWidth() int {
	if r.Width() > 0 {
		return r.InnerX2() - r.InnerX1()
	}
	return 0
}

// DrawableHeight returns the max height of the rectangle minux padding
func (r *Rectangle) DrawableHeight() int {
	if r.Height() > 0 {
		return r.InnerY2() - r.InnerY1()
	}
	return 0
}

// ZeroSize returns true if the rectangle has no size
func (r *Rectangle) ZeroSize() bool {
	if r.Width() == 0 && r.Height() == 0 {
		return true
	}
	return false
}
