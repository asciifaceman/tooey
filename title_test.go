package tooey

import "testing"

func TestTitle(t *testing.T) {

	err := InitSim()
	if err != nil {
		t.Fatal("Failed to initialize simulation screen")
	}
	defer Close()

	s := GetRootScreen()

	title := NewTitle()

	testRect := NewRectangle()
	testRect.SetRect(0, 0, 10, 2)

	title.Draw(s, &testRect)

	for x := 0; x < testRect.Dx(); x++ {
		c, _, _, _ := s.GetContent(x, 0)
		if c != rune(' ') {
			t.Fatalf("Expected x[%d] to be empty character, but got [%s][%v]", x, string(c), c)
		}
	}

	s.Clear()

	// 4 char string will be padded to total length 8
	title.Set("TEST")

	title.Draw(s, &testRect)

	// start point should be 1 + 2
	start := testRect.InnerX1() + title.Padding.Left

	d, _, _, _ := s.GetContent(start, 0)
	if d != ELLIPSES {
		t.Fatalf("Expected %s to be drawn at %d,0 but got [%s](%v)", string(ELLIPSES), start, string(d), d)
	}

	s.Clear()
	testRect.SetRect(0, 0, 20, 2)

	// 4 char string will be padded to total length 8
	title.Set("TEST")

	title.Draw(s, &testRect)

	d, _, _, _ = s.GetContent(start, 0)
	if d != rune('T') {
		t.Fatalf("Expected 'T' to be drawn at %d,0 but got [%s](%v)", start, string(d), d)
	}

	d, _, _, _ = s.GetContent(start+3, 0)
	if d != rune('T') {
		t.Fatalf("Expected 'T' to be drawn at %d,0 but got [%s](%v)", start, string(d), d)
	}
}
