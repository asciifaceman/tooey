package tooey

import (
	"testing"
)

func DoSliceTestMapWithFunc(t *testing.T, tests map[string]string, f func([]rune) []rune) {
	for content, expected := range tests {
		rContent := []rune(content)
		result := f(rContent)

		if string(result) != expected {
			t.Fatalf("Expected [%s] (%v) to equal [%s] (%v) | Test Case: %s", string(result), result, expected, []rune(expected), content)
		}
	}
}

func TestShiftRuneSliceRight(t *testing.T) {
	tests := map[string]string{
		"abcdefg": "gabcdef",
		"abcdef ": " abcdef",
		"aaaaaO":  "Oaaaaa",
	}

	DoSliceTestMapWithFunc(t, tests, ShiftRuneSliceRight)
}

func TestShiftRuneSliceLeft(t *testing.T) {
	tests := map[string]string{
		"abcdef":  "bcdefa",
		" abcdef": "abcdef ",
		"Oaaaaa":  "aaaaaO",
	}

	DoSliceTestMapWithFunc(t, tests, ShiftRuneSliceLeft)

}

func TestShiftRuneWhitespaceToLeft(t *testing.T) {
	tests := map[string]string{
		"abcdef":     "abcdef",
		"abcdef    ": "    abcdef",
		"     ":      "     ",
	}

	DoSliceTestMapWithFunc(t, tests, ShiftRuneWhitespaceToLeft)

}

func TestShiftRuneWhitespaceToRight(t *testing.T) {
	tests := map[string]string{
		"abcdef":     "abcdef",
		"    abcdef": "abcdef    ",
		"     ":      "     ",
	}

	DoSliceTestMapWithFunc(t, tests, ShiftRuneWhitespaceToRight)

}

func TestSpreadWhitespaceAcrossSliceInterior(t *testing.T) {
	tests := map[string]string{
		"  abc def":                     "abc   def",
		"     abc def ghi  ":            "abc        def ghi",
		"     ":                         "     ",
		"          abc def ghi        ": " ",
	}

	DoSliceTestMapWithFunc(t, tests, SpreadWhitespaceAcrossSliceInterior)
}

func TestNormalizeLeftWhitespace(t *testing.T) {
	tests := map[string]string{
		"  abc def":          "abc   def",
		"     abc def ghi  ": "abc    def   ghi  ",
	}

	DoSliceTestMapWithFunc(t, tests, NormalizeLeftWhitespace)
}

func TestCountWordsInRuneSlice(t *testing.T) {
	tests := map[string]int{
		"  abc  def  ghi":               3,
		"abc def":                       2,
		"          abc def ghi        ": 3,
	}

	for content, expected := range tests {
		rContent := []rune(content)
		c := CountWordsInRuneSlice(rContent)
		if c != expected {
			t.Fatalf("Expected [%d] words in [%s] but counted [%d]", expected, content, c)
		}
	}
}

func TestCheckWhichPositionHasFewest(t *testing.T) {
	source := make([]int, 5)
	source[0] = 5
	source[1] = 4
	source[2] = 2
	source[3] = 3
	source[4] = 7

	d := CheckWhichPositionHasFewest(source)

	if d != 2 {
		t.Fatalf("Failed to detect position with lowest count, got [%d]", d)
	}

}
