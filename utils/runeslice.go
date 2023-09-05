// package words bundles some utilities and helpers for dealing with rune slice strings
// and text operations like formatting / alignment
package utils

import (
	"unicode"
)

// NewRowsFromString returns a []Rows of width honoring wrap
func NewRowsFromString(s string, width int, wrap bool) []Row {
	rows := make([]Row, 0)

	return rows
}

type Row []rune

// Len returns the length of the rune slice
func (r Row) Len() int {
	return len(r)
}

// ShiftLeft shifts the rune slice one to the left and returns a copy
func (r Row) ShiftLeft() Row {
	row := r

	if row.Len() < 1 || row.OnlyWhitespace() {
		return r
	}

	row = append(row[1:], row[0])
	return row
}

// ShiftRight shifts the rune slice one to the right and returns a copy
func (r Row) ShiftRight() Row {
	row := r

	if row.Len() < 1 || row.OnlyWhitespace() {
		return r
	}

	row = append(row[len(row)-1:], row[:len(row)-1]...)
	return row
}

// ShiftWhitespaceLeft shifts any whitespace on the right to the left
func (r Row) ShiftWhitespaceLeft() []rune {
	row := r
	if row.Len() < 1 || row.OnlyWhitespace() {
		return row
	}

	for {
		if unicode.IsSpace(row[r.Len()-1]) {
			row = r.ShiftRight()
		} else {
			break
		}
	}
	return row
}

// ShiftWhitespaceRight shifts any whitespace on the left to the right
func (r Row) ShiftWhitespaceRight() []rune {
	row := r
	if row.Len() < 1 || row.OnlyWhitespace() {
		return row
	}

	for {
		if unicode.IsSpace(row[0]) {
			row = r.ShiftLeft()
		} else {
			break
		}
	}
	return row
}

// OnlyWhitespace returns true if the rune slice contains only whitespace
func (r Row) OnlyWhitespace() bool {
	for _, char := range r {
		if !unicode.IsSpace(char) {
			return false
		}
	}
	return true
}

// CountWhitespace returns a count of the white space characters
func (r Row) CountWhiteSpace() int {
	count := 0
	for _, char := range r {
		if unicode.IsSpace(char) {
			count++
		}
	}
	return count
}

// CountWords returns a count of how many groupings of non-whitespace characters
// are inside the rune slice separated by whitespace
func (r Row) CountWords() int {
	count := 0
	word := false

	for _, char := range r {
		if !unicode.IsSpace(char) {
			if !word {
				word = true
				count++
			}
		} else {
			if word {
				word = false
			}
		}
	}
	return count
}

// NormalizeWhitespace attempts to distribute edge whitespace across the interior
// of the rune slice between blocks of non-whitespace characters
func (r Row) NormalizeWhitespace() []rune {
	row := r

	if r.OnlyWhitespace() || r.Len() < 1 || r.CountWords() < 2 {
		return row
	}

	row = r.ShiftWhitespaceLeft()

	return row
}
