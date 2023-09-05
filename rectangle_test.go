package tooey

import (
	"image"
	"testing"
)

func TestRectangle(t *testing.T) {
	r := NewRectangle()

	// Brand new empty rect should have no bounds
	if r.Width() != 0 || r.Height() != 0 {
		t.Fatalf("Expected raw Rectangle to have 0 bounds but found Width: %d Height: %d", r.Width(), r.Height())
	}

	if !r.ZeroSize() {
		t.Fatal("Expected ZeroSize to return true on raw Rectangle")
	}

	inner := r.GetInnerRect()
	if inner.Dx() > 0 || inner.Dy() > 0 {
		t.Fatal("Expected GetInnerRect to return zero size rect for raw Rectangle")
	}

	tests := []image.Rectangle{
		image.Rect(0, 0, 10, 10),
		image.Rect(10, 10, 25, 25),
	}

	p := NewDefaultPadding()

	for iter, test := range tests {

		r.SetRect(test.Min.X, test.Min.Y, test.Max.X, test.Max.Y)

		if r.Width() != test.Dx() {
			t.Fatalf("Expected tests[%d] to have width [%d] but got [%d]", iter, test.Dx(), r.Width())
		}

		if r.Height() != test.Dy() {
			t.Fatalf("Expected tests[%d] to have height [%d] but got [%d]", iter, test.Dy(), r.Height())
		}

		paddingWidthCalc := test.Dx() - (p.Left + p.Right)
		paddingHeightCalc := test.Dy() - (p.Top + p.Bottom)

		if r.DrawableWidth() != paddingWidthCalc {
			t.Fatalf("Expected tests[%d] to have inner drawable width [%d] with default padding but got [%d]", iter, paddingWidthCalc, r.DrawableWidth())
		}

		if r.DrawableHeight() != paddingHeightCalc {
			t.Fatalf("Expected tests[%d] to have inner drawable height [%d] with default padding but got [%d]", iter, paddingHeightCalc, r.DrawableHeight())
		}

		if r.ZeroSize() {
			t.Fatal("Did not expect ZeroSize to return true on Rectangle with bounds")
		}

	}

}
