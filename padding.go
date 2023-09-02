package tooey

// Padding provides an inset from an outer edge
type Padding struct {
	Left   int
	Top    int
	Right  int
	Bottom int
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
		Left:   2,
		Top:    1,
		Right:  2,
		Bottom: 1,
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
