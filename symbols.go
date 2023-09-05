package tooey

import "github.com/gdamore/tcell/v2"

const (
	DOT      = '•'
	ELLIPSES = '…'

	UP_ARROW   = '▲'
	DOWN_ARROW = '▼'

	COLLAPSED = '+'
	EXPANDED  = '−'
)

const (
	StylizedUL = '╒'
	DefaultUL  = tcell.RuneULCorner
	DefaultUR  = tcell.RuneURCorner
	DefaultLL  = tcell.RuneLLCorner
	DefaultLR  = tcell.RuneLRCorner
	DefaultH   = tcell.RuneHLine
	DefaultV   = tcell.RuneVLine
	RoundedUL  = '╭'
	RoundedUR  = '╮'
	RoundedLL  = '╰'
	RoundedLR  = '╯'
	DoubleH    = '═'
	DoubleV    = '║'
	DoubleUL   = '╔'
	DoubleUR   = '╗'
	DoubleLL   = '╚'
	DoubleLR   = '╝'
)

var (
	BARS = [...]rune{' ', '▁', '▂', '▃', '▄', '▅', '▆', '▇', '█'}

	SHADED_BLOCKS = [...]rune{' ', '░', '▒', '▓', '█'}

	IRREGULAR_BLOCKS = [...]rune{
		' ', '▘', '▝', '▀', '▖', '▌', '▞', '▛',
		'▗', '▚', '▐', '▜', '▄', '▙', '▟', '█',
	}

	BRAILLE_OFFSET = '\u2800'
	BRAILLE        = [4][2]rune{
		{'\u0001', '\u0008'},
		{'\u0002', '\u0010'},
		{'\u0004', '\u0020'},
		{'\u0040', '\u0080'},
	}

	DOUBLE_BRAILLE = map[[2]int]rune{
		{0, 0}: '⣀',
		{0, 1}: '⡠',
		{0, 2}: '⡐',
		{0, 3}: '⡈',

		{1, 0}: '⢄',
		{1, 1}: '⠤',
		{1, 2}: '⠔',
		{1, 3}: '⠌',

		{2, 0}: '⢂',
		{2, 1}: '⠢',
		{2, 2}: '⠒',
		{2, 3}: '⠊',

		{3, 0}: '⢁',
		{3, 1}: '⠡',
		{3, 2}: '⠑',
		{3, 3}: '⠉',
	}

	SINGLE_BRAILLE_LEFT  = [4]rune{'\u2840', '⠄', '⠂', '⠁'}
	SINGLE_BRAILLE_RIGHT = [4]rune{'\u2880', '⠠', '⠐', '⠈'}
)
