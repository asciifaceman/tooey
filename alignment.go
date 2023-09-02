package tooey

// Alignment provides a mechanism for Styles to control
// alignment of certain elements
// Not widely employed but can be at the widget level
type Alignment uint

const (
	AlignLeft Alignment = iota
	AlignCenter
	AlignRight
	AlignFull
)
