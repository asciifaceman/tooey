package widgets

import (
	"testing"

	"github.com/asciifaceman/tooey"
)

func TestText(t *testing.T) {
	/*
		Test that text draws

		I think this test could be dramatically improved but I wanted
		to have something here
	*/

	err := tooey.InitSim()
	if err != nil {
		t.Fatal("Failed to initialize simulation screen")
	}
	defer tooey.Close()
	//w, h := tooey.DrawableDimensions()

	s := tooey.GetRootScreen()

	txt := NewText(tooey.DefaultTheme)
	txt.Content = "TEST"
	txt.SetRect(0, 0, 20, 20)
	txt.Draw(s)

	tests := map[int]rune{
		2: 'T',
		3: 'E',
		4: 'S',
		5: 'T',
	}

	for x, r := range tests {
		f, _, _, _ := s.GetContent(x, 2)
		if f != r {
			t.Fatalf("Expected rune at (%d,%d) to be [%s][%v] but got [%s][%v]", x, 2, string(r), r, string(f), f)
		}

	}

}
