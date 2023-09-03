// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package tooey

import (
	"fmt"
	"math"
	"reflect"
	"unicode"

	"github.com/davecgh/go-spew/spew"
	rw "github.com/mattn/go-runewidth"
)

// InterfaceSlice takes an []interface{} represented as an interface{} and converts it
// https://stackoverflow.com/questions/12753805/type-converting-slices-of-interfaces-in-go
func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic(fmt.Sprintf("InterfaceSlice() given a non-slice type: %v [%v]", s.Kind(), spew.Sdump(slice)))
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

// TrimString trims a string to a max length and adds '…' to the end if it was trimmed.
func TrimString(s string, w int) string {
	if w <= 0 {
		return ""
	}
	if rw.StringWidth(s) > w {
		return rw.Truncate(s, w, string(ELLIPSES))
	}
	return s
}

// ShiftRuneSliceRight takes a []rune and shifts everything right, wrapping the right
// most character back to the left most position
func ShiftRuneSliceRight(slice []rune) []rune {
	if len(slice) < 1 {
		return slice
	}

	contentLength := len(slice)
	newSlice := make([]rune, 0)

	newSlice = append(newSlice, slice[contentLength-1])
	for i, r := range slice {
		if i == contentLength-1 {
			break
		}
		newSlice = append(newSlice, r)
	}

	return newSlice

}

// ShiftRuneSliceLeft takes a []rune and shifts everything left, wrapping the
// right most character back to the right most position
func ShiftRuneSliceLeft(slice []rune) []rune {
	if len(slice) < 1 {
		return slice
	}

	newSlice := append(slice[1:], slice[0])

	return newSlice

}

// ShiftRuneWhitespaceToLeft takes a []rune and moves all whitespace to the left
func ShiftRuneWhitespaceToLeft(slice []rune) []rune {
	if len(slice) < 1 {
		return slice
	}

	contentLength := len(slice)
	newSlice := slice

	if !ContainsNonWhitespace(slice) {
		return slice
	}

	for {
		if unicode.IsSpace(newSlice[contentLength-1]) {
			newSlice = ShiftRuneSliceRight(newSlice)
		} else {
			break
		}
	}
	return newSlice
}

// ShiftRuneWhitespaceToRight takes a []rune and moves all whitespace to the right
func ShiftRuneWhitespaceToRight(slice []rune) []rune {
	if !ContainsNonWhitespace(slice) || len(slice) < 1 {
		return slice
	}

	newSlice := slice

	for {
		if unicode.IsSpace(newSlice[0]) {
			newSlice = ShiftRuneSliceLeft(newSlice)
		} else {
			break
		}
	}
	return newSlice
}

/*
	SpreadWhitespaceAcrossSliceInterior takes a []rune

and attempts to distribute it's whitespace across the
width of the slice interior but not at the outside edges

Take the string "abc def   gh ij   "

"abc_def___gh_ij___" would try to make "abc___def__gh___ij"
"__abc_def___gh_ij" would try to make "abc__def__gh___ij"
*/
func SpreadWhitespaceAcrossSliceInterior(slice []rune) []rune {
	wordCount := CountWordsInRuneSlice(slice)

	if !ContainsNonWhitespace(slice) || len(slice) < 1 || wordCount < 2 {
		return slice
	}

	// Shift all right whitespace to the left
	newSlice := ShiftRuneWhitespaceToLeft(slice)

	// Move left whitespace inwards
	newSlice = MoveLeftWhitespaceInwards(newSlice)

	return newSlice
}

// CheckWhichPositionHasFewest returns index of the position with the lowest
// count
func CheckWhichPositionHasFewest(positions []int) int {
	lowestCount := 0

	for _, i := range positions {
		if lowestCount >= i {
			continue
		}
		lowestCount = i
	}
	return lowestCount
}

func MoveLeftWhitespaceInwards(slice []rune) []rune {
	wordCount := CountWordsInRuneSlice(slice)

	if !ContainsNonWhitespace(slice) || len(slice) < 1 || wordCount < 2 {
		return slice
	}

	newSlice := slice
	position := make([]int, wordCount)

OUTER:
	for {
		insideWord := false
		//interior := false
		if unicode.IsSpace(newSlice[0]) {
			wordIndex := 0
			for i, r := range newSlice {
				bestIndex := CheckWhichPositionHasFewest(position)
				// range through the runes and detect words
				// place left spaces throughout the rune slice between words interior

				if !unicode.IsSpace(r) {
					if !insideWord {
						insideWord = true
					}
				} else {
					if insideWord {
						insideWord = false
						wordIndex++
					}
				}

				// we have to wait to do anything until we've passed our first non
				// whitespace character so it's "inside"
				//if !unicode.IsSpace(r) {
				//	interior = true
				//}
				if i == len(newSlice) {
					// if we reached the end
					// without finding a spot
					// give up
					//continue
					break OUTER
				}
				if bestIndex == wordIndex {
					if unicode.IsSpace(r) {
						newSlice = newSlice[1:]
						newSlice = append(newSlice[:i+1], newSlice[i:]...)
						newSlice[i] = rune(' ')
						position[wordIndex]++
						break
					}
				}
			}
		} else {
			break
		}
	}

	return newSlice
}

// CountWordsInRuneSlice counts how many blocks of non-whitespace characters
// are inside the rune
//
// This can be used to traverse whitespace between the left and right boundaries
// use count-1 for right bound to prevent traversing outside of the word bounds
func CountWordsInRuneSlice(slice []rune) int {
	count := 0
	insideWord := false

	for _, r := range slice {
		if !unicode.IsSpace(r) {
			if !insideWord {
				insideWord = true
				count++
			}
		} else {
			if insideWord {
				insideWord = false
			}
		}

	}

	return count
}

// CountWhiteSpace returns a count of the number of space characters in a []rune
func CountWhiteSpace(slice []rune) int {
	count := 0

	for _, r := range slice {
		if unicode.IsSpace(r) {
			count++
		}
	}

	return count
}

// ContainsNonWhitespace returns true if a given rune slice has any non-whitespace
// runes
func ContainsNonWhitespace(slice []rune) bool {
	for _, r := range slice {
		if !unicode.IsSpace(r) {
			return true
		}
	}
	return false
}

func SelectColor(colors []Color, index int) Color {
	return colors[index%len(colors)]
}

func SelectStyle(styles []Style, index int) Style {
	return styles[index%len(styles)]
}

// Math ------------------------------------------------------------------------

func SumIntSlice(slice []int) int {
	sum := 0
	for _, val := range slice {
		sum += val
	}
	return sum
}

func SumFloat64Slice(data []float64) float64 {
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum
}

func GetMaxIntFromSlice(slice []int) (int, error) {
	if len(slice) == 0 {
		return 0, fmt.Errorf("cannot get max value from empty slice")
	}
	var max int
	for _, val := range slice {
		if val > max {
			max = val
		}
	}
	return max, nil
}

func GetMaxFloat64FromSlice(slice []float64) (float64, error) {
	if len(slice) == 0 {
		return 0, fmt.Errorf("cannot get max value from empty slice")
	}
	var max float64
	for _, val := range slice {
		if val > max {
			max = val
		}
	}
	return max, nil
}

func GetMaxFloat64From2dSlice(slices [][]float64) (float64, error) {
	if len(slices) == 0 {
		return 0, fmt.Errorf("cannot get max value from empty slice")
	}
	var max float64
	for _, slice := range slices {
		for _, val := range slice {
			if val > max {
				max = val
			}
		}
	}
	return max, nil
}

func RoundFloat64(x float64) float64 {
	return math.Floor(x + 0.5)
}

func FloorFloat64(x float64) float64 {
	return math.Floor(x)
}

func AbsInt(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func MinFloat64(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func MaxFloat64(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
