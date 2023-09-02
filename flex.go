package tooey

/*
	This is not meant to be an exhaustive implementation
	of flex
*/

type FlexibleElement struct {
	Children  interface{}
	Direction FlexDirection
	XRatio    float64
	YRatio    float64
	WRatio    float64
	HRatio    float64
	Leaf      bool
	ratio     float64
}

// NewFlexColumn accepts a height percentage weight and Drawable children
func NewFlexColumn(ratio float64, i ...interface{}) FlexibleElement {
	_, ok := i[0].(Drawable)
	child := i[0]
	if !ok {
		child = i
	}
	return FlexibleElement{
		Children:  child,
		Direction: FlexColumn,
		Leaf:      ok,
		ratio:     ratio,
	}
}

// NewFlexibleRow accepts a width percentage weight and drawable children
func NewFlexibleRow(ratio float64, i ...interface{}) FlexibleElement {
	_, ok := i[0].(Drawable)
	child := i[0]
	if !ok {
		child = i
	}
	return FlexibleElement{
		Children:  child,
		Direction: FlexRow,
		Leaf:      ok,
		ratio:     ratio,
	}
}
