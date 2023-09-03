package widgets

import (
	"testing"

	"github.com/asciifaceman/tooey"
	"github.com/asciifaceman/tooey/themes"
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
	w, _ := tooey.DrawableDimensions()

	s := tooey.GetRootScreen()

	txt := NewText(themes.ThemeRetroTerminalGreen)
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

	s.Clear()

	fullAlignTheme := *themes.ThemeRetroTerminalGreen
	fullAlignTheme.Text.Align = tooey.AlignFull
	txt2 := NewText(&fullAlignTheme)
	txt2.Content = "     TEST TEST TEST TEST TEST TEST TEST  TEST TEST TEST TEST  TEST TEST TEST TEST TEST TEST TEST     "
	txt2.SetRect(0, 0, 20, 10)
	txt2.Draw(s)

	for y := 0; y < 10; y++ {
		if y == 0 || y == 1 {
			continue
		}
		for x := 0; x < w; x++ {
			if x > 20 {
				continue
			}
			v, _, _, _ := s.GetContent(x, y)
			t.Logf("X: %d Y: %d| rune: [%s]", x, y, string(v))
		}
	}

	s.Clear()

	rightAlignTheme := *themes.ThemeRetroTerminalGreen
	rightAlignTheme.Text.Align = tooey.AlignRight
	txt3 := NewText(&rightAlignTheme)
	txt3.Content = "This should right align but I don't know how to test that just yet"
	txt3.SetRect(0, 0, 10, 5)
	txt3.Draw(s)

	// need to find a way to actually write tests for the visual elements
	// don't quite know how to approach this

}
