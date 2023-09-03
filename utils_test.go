package tooey

import (
	"testing"
)

func DoSliceTestMapWithFunc(t *testing.T, tests map[string]string, f func([]rune) []rune) {
	for content, expected := range tests {
		rContent := []rune(content)
		result := f(rContent)

		if string(result) != expected {
			t.Fatalf("Expected [%s] (%v) to equal [%s] (%v)", string(result), result, expected, []rune(expected))
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
		"  abc def":          "abc   def",
		"     abc def ghi  ": "abc        def ghi",
		"     ":              "     ",
	}

	DoSliceTestMapWithFunc(t, tests, SpreadWhitespaceAcrossSliceInterior)
}

func TestMoveLeftWhitespaceInwards(t *testing.T) {
	tests := map[string]string{
		"  abc def":          "abc   def",
		"     abc def ghi  ": "abc      def ghi  ",
	}

	DoSliceTestMapWithFunc(t, tests, MoveLeftWhitespaceInwards)
}

func TestCountWordsInRuneSlice(t *testing.T) {
	tests := map[string]int{
		"  abc  def  ghi": 3,
		"abc def":         2,
	}

	for content, expected := range tests {
		rContent := []rune(content)
		c := CountWordsInRuneSlice(rContent)
		if c != expected {
			t.Fatalf("Expected [%d] words in [%s] but counted [%d]", expected, content, c)
		}
	}
}
