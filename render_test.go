package tooey

import "testing"

func TestRender(t *testing.T) {
	err := InitSim()
	if err != nil {
		t.Fatal("Failed to initialize simulation screen")
	}
	defer Close()
	s := GetRootScreen()

	e := NewElement(nil)
	e.SetRect(0, 0, 15, 10)
	e.Title.Content = "test"

	Render(e)

	runeT, _, _, _ := s.GetContent(3, 0)
	check := rune(e.Title.Content[0])
	if runeT != check {
		t.Fatalf("Expected drawn rune at [3,0] to be [%v][%s] but got [%v][%s]", check, string(check), runeT, string(runeT))
	}
	runeE, _, _, _ := s.GetContent(4, 0)
	check = rune(e.Title.Content[1])
	if runeE != check {
		t.Fatalf("Expected drawn rune at [4,0] to be [%v][%s] but got [%v][%s]", check, string(check), runeE, string(runeE))
	}
	runeS, _, _, _ := s.GetContent(5, 0)
	check = rune(e.Title.Content[2])
	if runeS != check {
		t.Fatalf("Expected drawn rune at [5,0] to be [%v][%s] but got [%v][%s]", check, string(check), runeS, string(runeS))
	}
	runeT2, _, _, _ := s.GetContent(6, 0)
	check = rune(e.Title.Content[3])
	if runeT2 != check {
		t.Fatalf("Expected drawn rune at [6,0] to be [%v][%s] but got [%v][%s]", check, string(check), runeT2, string(runeT2))
	}
}
