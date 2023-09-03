package tooey

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestDrawableDimensions(t *testing.T) {
	err := InitSim()
	if err != nil {
		t.Fatal("Failed to initialize simulation screen")
	}
	defer Close()

	x, y := DrawableDimensions()

	if x != 79 && y != 24 {
		t.Fatalf("Expected drawable dimensions of sim screen to be x[79] y[24] but got x[%d] y[%d]", x, y)
	}

	s := GetRootScreen()

	s.SetContent(1, 2, rune('#'), nil, tcell.StyleDefault)

	result, _, _, _ := s.GetContent(1, 2)

	if result != rune('#') {
		t.Fatal("Did not get back rune that was set")
	}

	Clear()

	result, _, _, _ = s.GetContent(1, 2)

	if result != rune(' ') {
		t.Fatal("Screen was not cleared by clear call")
	}

}
