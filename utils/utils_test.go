package utils

import (
	"testing"
)

type TestCase struct {
	TestID         string
	Purpose        string
	F              func([]rune) []rune
	Case           string
	ExpectedResult string
}

func DoSliceTest(t *testing.T, c *TestCase) {
	rContent := []rune(c.Case)
	result := c.F(rContent)

	if string(result) != c.ExpectedResult {
		t.Fatalf("[%s] [%s]\n - Expected [%s] (%v) to equal [%s] (%v) \nSource: [%s]", c.TestID, c.Purpose, string(result), result, c.ExpectedResult, []rune(c.ExpectedResult), c.Case)
	}
}

func TestShiftRuneSliceRight(t *testing.T) {
	testID := "TestShiftRuneSliceRight"
	tests := []*TestCase{
		{
			TestID:         testID,
			Purpose:        "Ensure func shifts runes to the right",
			F:              ShiftRuneSliceRight,
			Case:           "abcdefg",
			ExpectedResult: "gabcdef",
		},
		{
			TestID:         testID,
			Purpose:        "Ensure func shifts runes to the right",
			F:              ShiftRuneSliceRight,
			Case:           "abcdef ",
			ExpectedResult: " abcdef",
		},
		{
			TestID:         testID,
			Purpose:        "Ensure func shifts runes to the right",
			F:              ShiftRuneSliceRight,
			Case:           "  f  e   h sdfd34 %1",
			ExpectedResult: "1  f  e   h sdfd34 %",
		},
	}

	for _, test := range tests {
		DoSliceTest(t, test)
	}
}

func TestShiftRuneSliceLeft(t *testing.T) {
	testID := "TestShiftRuneSliceLeft"
	tests := []*TestCase{
		{
			TestID:         testID,
			Purpose:        "Ensure func shifts runes to the left",
			F:              ShiftRuneSliceLeft,
			Case:           "abcdef",
			ExpectedResult: "bcdefa",
		},
		{
			TestID:         testID,
			Purpose:        "Ensure func shifts runes to the left",
			F:              ShiftRuneSliceLeft,
			Case:           "abcdef ",
			ExpectedResult: "bcdef a",
		},
		{
			TestID:         testID,
			Purpose:        "Ensure func shifts runes to the left",
			F:              ShiftRuneSliceLeft,
			Case:           "Oaaaaa",
			ExpectedResult: "aaaaaO",
		},
	}

	for _, test := range tests {
		DoSliceTest(t, test)
	}

}

func TestShiftRuneWhitespaceToLeft(t *testing.T) {
	testID := "TestShiftRuneWhitespaceToLeft"
	tests := []*TestCase{
		{
			TestID:         testID,
			Purpose:        "Ensure func shifts whitespace to the left",
			F:              ShiftRuneWhitespaceToLeft,
			Case:           "abcdef",
			ExpectedResult: "abcdef",
		},
		{
			TestID:         testID,
			Purpose:        "Ensure func shifts whitespace to the left",
			F:              ShiftRuneWhitespaceToLeft,
			Case:           "abcdef     ",
			ExpectedResult: "     abcdef",
		},
		{
			TestID:         testID,
			Purpose:        "Ensure func doesn't spinlock on string of only whitespace",
			F:              ShiftRuneWhitespaceToLeft,
			Case:           "     ",
			ExpectedResult: "     ",
		},
		{
			TestID:         testID,
			Purpose:        "Ensure func shifts whitespace to the left",
			F:              ShiftRuneWhitespaceToLeft,
			Case:           "     abcdef          ",
			ExpectedResult: "               abcdef",
		},
	}

	for _, test := range tests {
		DoSliceTest(t, test)
	}

}

func TestShiftRuneWhitespaceToRight(t *testing.T) {
	testID := "ShiftRuneWhitespaceToRight"
	tests := []*TestCase{
		{
			TestID:         testID,
			Purpose:        "Ensure func shifts whitespace to the right",
			F:              ShiftRuneWhitespaceToRight,
			Case:           "abcdef",
			ExpectedResult: "abcdef",
		},
		{
			TestID:         testID,
			Purpose:        "Ensure func shifts whitespace to the right",
			F:              ShiftRuneWhitespaceToRight,
			Case:           "     abcdef",
			ExpectedResult: "abcdef     ",
		},
		{
			TestID:         testID,
			Purpose:        "Ensure func doesn't spinlock on string of only whitespace",
			F:              ShiftRuneWhitespaceToRight,
			Case:           "     ",
			ExpectedResult: "     ",
		},
		{
			TestID:         testID,
			Purpose:        "Ensure func shifts whitespace to the right",
			F:              ShiftRuneWhitespaceToRight,
			Case:           "          abcdef     ",
			ExpectedResult: "abcdef               ",
		},
	}

	for _, test := range tests {
		DoSliceTest(t, test)
	}
}

func TestSpreadWhitespaceAcrossSliceInterior(t *testing.T) {
	testID := "TestSpreadWhitespaceAcrossSliceInterior"
	tests := []*TestCase{
		{
			TestID:         testID,
			Purpose:        "Ensure func evenly distributes left whitespace across row",
			F:              SpreadWhitespaceAcrossSliceInterior,
			Case:           "     123 56  9",
			ExpectedResult: "123    56    9",
		},
		{
			TestID:         testID,
			Purpose:        "Ensure func evenly distributes left whitespace across row",
			F:              SpreadWhitespaceAcrossSliceInterior,
			Case:           "     a b cd e  f    g  ",
			ExpectedResult: "a   b   cd  e   f     g",
		},
		{
			TestID:         testID,
			Purpose:        "Ensure func doesn't spinlock on string of only whitespace",
			F:              SpreadWhitespaceAcrossSliceInterior,
			Case:           "          ",
			ExpectedResult: "          ",
		},
		{
			TestID:         testID,
			Purpose:        "Ensure func evenly distributes left whitespace across row",
			F:              SpreadWhitespaceAcrossSliceInterior,
			Case:           "          abc def ghi        ",
			ExpectedResult: "abc          def          ghi",
		},
	}

	for _, test := range tests {
		DoSliceTest(t, test)
	}
}

func TestNormalizeLeftWhitespace(t *testing.T) {
	testID := "TestNormalizeLeftWhitespace"
	tests := []*TestCase{
		{
			TestID:         testID,
			Purpose:        "Ensure func evenly distributes whitespace across width",
			F:              NormalizeLeftWhitespace,
			Case:           "  abc def",
			ExpectedResult: "abc   def",
		},
		{
			TestID:         testID,
			Purpose:        "Ensure func evenly distributes whitespace across width",
			F:              NormalizeLeftWhitespace,
			Case:           "     abc def ghi  ",
			ExpectedResult: "abc    def   ghi  ",
		},
		{
			TestID:         testID,
			Purpose:        "Ensure func doesn't spinlock on string of only whitespace",
			F:              NormalizeLeftWhitespace,
			Case:           "     ",
			ExpectedResult: "     ",
		},
		{
			TestID:         testID,
			Purpose:        "Ensure func doesn't touch whitespace without gaps to respec it into",
			F:              NormalizeLeftWhitespace,
			Case:           "          abcdef     ",
			ExpectedResult: "          abcdef     ",
		},
	}

	for _, test := range tests {
		DoSliceTest(t, test)
	}
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

func TestInterfaceSlice(t *testing.T) {
	testID := "TestInterfaceSlice"
	test1Content := []string{
		"this is a string",
		"this is another string",
	}
	test2Content := 5
	test3Content := rune('#')
	test1 := []interface{}{
		test1Content,
		test2Content,
		test3Content,
	}
	var i interface{} = test1

	d := InterfaceSlice(i)

	if len(d) != 3 {
		t.Fatalf("[%s] - Expected length of result to be 3 but got [%d]", testID, len(d))
	}

	for i, result := range d {
		if i == 0 {
			val, ok := result.([]string)
			if !ok {
				t.Fatal("Failed to unpack expected string slice from interface")
			}
			if len(val) != 2 {
				t.Fatalf("Count of string slice is wrong. Expected 2 but got %d", len(val))
			}
			for i, str := range val {
				if test1Content[i] != str {
					t.Fatalf("Expected value at index [%d] to be [%s] but got [%s]", i, test1Content[i], str)
				}
			}
		} else if i == 1 {
			val, ok := result.(int)
			if !ok {
				t.Fatal("Failed to unpack expected int from interface")
			}
			if val != 5 {
				t.Fatalf("Expected to get int 5 but got %d", val)
			}
		} else if i == 2 {
			val, ok := result.(rune)
			if !ok {
				t.Fatal("Failed to unpack expected rune from interface")
			}
			if val != test3Content {
				t.Fatalf("Expected to get rune %s but got %s", string(test3Content), string(val))
			}
		}
	}

}

func TestCountWhiteSpace(t *testing.T) {
	test1 := "this should have 4 whitespaces"
	test1r := []rune(test1)
	d := CountWhiteSpace(test1r)
	if d != 4 {
		t.Fatalf("Expected 4 whitespaces but counted %d", d)
	}
}

func TestTrimString(t *testing.T) {
	test1Content := "first string"
	test1Expected := "first…"

	d := TrimString(test1Content, 6)
	if d != test1Expected {
		t.Fatalf("Expected %s to match %s", d, test1Expected)
	}

	d2 := TrimString(test1Content, 0)
	if d2 != "" {
		t.Fatal("Expected string to be truncated and it wasn't")
	}

	d3 := TrimString(test1Content, len(test1Content))
	if d3 != test1Content {
		t.Fatal("Expected string to be untouched but it was altered")
	}
}
