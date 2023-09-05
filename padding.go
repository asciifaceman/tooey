package tooey

// Padding provides an inset from an outer edge
type Padding struct {
	Left   int
	Top    int
	Right  int
	Bottom int
}

// TotalWidth returns Left + Right
func (p *Padding) TotalWidth() int {
	return p.Left + p.Right
}

// TotalHeight returns Top + Bottom
func (p *Padding) TotalHeight() int {
	return p.Top + p.Bottom
}

// NewDefaultPadding returns a global padding of 1 all around
// which can account for a basic border
func NewDefaultPadding() *Padding {
	return &Padding{
		Left:   1,
		Top:    1,
		Right:  1,
		Bottom: 1,
	}
}

// NewTitlePadding returns a left & right padding of 2 to give the title
// more room to breathe
func NewTitlePadding() *Padding {
	return &Padding{
		Left:  2,
		Right: 2,
	}
}

// NewPadding returns an empty padding
func NewPadding() *Padding {
	return &Padding{
		Left:   0,
		Top:    0,
		Right:  0,
		Bottom: 0,
	}
}
