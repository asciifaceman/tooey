package tooey

import (
	"testing"
)

func TestContainers(t *testing.T) {
	/*
		Test that containers render and that a general container in
		flex row with two children will fill the screen
		half and half
	*/

	err := InitSim()
	if err != nil {
		t.Fatal("Failed to initialize simulation screen")
	}
	defer Close()
	w, h := DrawableDimensions()

	s := GetRootScreen()

	e := NewContainer(DefaultTheme)
	e.Direction = FlexRow
	e.SetRect(0, 0, w-1, h)

	e1 := NewElement(DefaultTheme)

	e2 := NewElement(DefaultTheme)

	e.Wrap(
		NewFlexChild(1, e1),
		NewFlexChild(1, e2),
	)

	Render(e)

	//	runeE1TopLeft, _, _, _ := s.GetContent(1, 0)
	//	if runeE1TopLeft != DefaultULCorner {
	//		t.Fatalf("Expected 1,0 to be [%s][%v] but got [%s][%v]\n", string(DefaultULCorner), DefaultULCorner, string(runeE1TopLeft), runeE1TopLeft)
	//	}

	topRow := make([]rune, w)
	for rowX := 0; rowX < w; rowX++ {
		if rowX == 0 {
			topRow[rowX] = DefaultULCorner
		} else if rowX == 78 {
			topRow[rowX] = DefaultURCorner
		} else {
			topRow[rowX] = DefaultHLine
		}
	}

	for x := 0; x < w; x++ {
		r, _, _, _ := s.GetContent(x, 0)
		if r != topRow[x] {
			t.Fatalf("Expected (%d, %d) rune to equal [%s][%v] but got [%s][%v]\n", x, 0, string(topRow[x]), topRow[x], string(r), r)
		}
	}

	secondRow := make([]rune, w)
	for rowX := 0; rowX < w; rowX++ {
		if rowX == 0 {
			secondRow[rowX] = DefaultVLine
		} else if rowX == 1 {
			secondRow[rowX] = DefaultULCorner
		} else if rowX == 39 {
			secondRow[rowX] = DefaultURCorner
		} else if rowX == 40 {
			secondRow[rowX] = DefaultULCorner
		} else if rowX == 77 {
			secondRow[rowX] = DefaultURCorner
		} else if rowX == 78 {
			secondRow[rowX] = DefaultVLine
		} else {
			secondRow[rowX] = DefaultHLine
		}
	}

	for x := 0; x < w; x++ {
		r, _, _, _ := s.GetContent(x, 1)
		if r != secondRow[x] {
			t.Fatalf("Expected (%d, %d) rune to equal [%s][%v] but got [%s][%v]\n", x, 1, string(secondRow[x]), secondRow[x], string(r), r)
		}
	}

	/*
		Used to discover chars
	*/

	//	for x := 0; x < w; x++ {
	//		r, _, _, _ := s.GetContent(x, 1)
	//		t.Logf("X: %d | [%s]\n", x, string(r))
	//	}

}
