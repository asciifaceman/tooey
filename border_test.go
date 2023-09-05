package tooey

import (
	"testing"
)

func TestBasicBorder(t *testing.T) {

	err := InitSim()
	if err != nil {
		t.Fatal("Failed to initialize simulation screen")
	}
	defer Close()

	s := GetRootScreen()

	b := NewBorder()

	testRect := NewRectangle()
	testRect.SetRect(0, 0, 2, 2)

	b.Draw(s, &testRect)

	// Sorry for my ugly test

	if d, _, _, _ := s.GetContent(0, 0); d != DefaultUL {
		t.Fatal("UL not correct")
	} else if d, _, _, _ := s.GetContent(1, 0); d != DefaultH {
		t.Fatal("Top H not correct")
	} else if d, _, _, _ := s.GetContent(2, 0); d != DefaultUR {
		t.Fatal("UR not correct")
	} else if d, _, _, _ := s.GetContent(0, 1); d != DefaultV {
		t.Fatal("L V not correct")
	} else if d, _, _, _ := s.GetContent(2, 1); d != DefaultV {
		t.Fatal("R V not correct")
	} else if d, _, _, _ := s.GetContent(0, 2); d != DefaultLL {
		t.Fatal("LL not correct")
	} else if d, _, _, _ := s.GetContent(1, 2); d != DefaultH {
		t.Fatal("L H not correct")
	} else if d, _, _, _ := s.GetContent(2, 2); d != DefaultLR {
		t.Fatal("LR not correct")
	} else if d, _, _, _ := s.GetContent(1, 1); d != rune(' ') {
		t.Fatal("center cell not empty")
	}
}
