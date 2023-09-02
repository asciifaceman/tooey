// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package tooey

import "github.com/gdamore/tcell/v2"

var StandardColors = []tcell.Color{
	ColorRed,
	ColorGreen,
	ColorYellow,
	ColorBlue,
	ColorMagenta,
	ColorCyan,
	ColorWhite,
}

var StandardStyles = []Style{
	NewStyle(ColorRed),
	NewStyle(ColorGreen),
	NewStyle(ColorYellow),
	NewStyle(ColorBlue),
	NewStyle(ColorMagenta),
	NewStyle(ColorCyan),
	NewStyle(ColorWhite),
}

type RootOldTheme struct {
	Default Style

	Block BlockOldTheme

	BarChart        BarChartOldTheme
	Gauge           GaugeOldTheme
	Plot            PlotOldTheme
	List            ListOldTheme
	Tree            TreeOldTheme
	Paragraph       ParagraphOldTheme
	PieChart        PieChartOldTheme
	Sparkline       SparklineOldTheme
	StackedBarChart StackedBarChartOldTheme
	Tab             TabOldTheme
	Table           TableOldTheme
}

type BlockOldTheme struct {
	Title  Style
	Border Style
}

type BarChartOldTheme struct {
	Bars   []tcell.Color
	Nums   []Style
	Labels []Style
}

type GaugeOldTheme struct {
	Bar   tcell.Color
	Label Style
}

type PlotOldTheme struct {
	Lines []tcell.Color
	Axes  tcell.Color
}

type ListOldTheme struct {
	Text Style
}

type TreeOldTheme struct {
	Text      Style
	Collapsed rune
	Expanded  rune
}

type ParagraphOldTheme struct {
	Text Style
}

type PieChartOldTheme struct {
	Slices []tcell.Color
}

type SparklineOldTheme struct {
	Title Style
	Line  tcell.Color
}

type StackedBarChartOldTheme struct {
	Bars   []tcell.Color
	Nums   []Style
	Labels []Style
}

type TabOldTheme struct {
	Active   Style
	Inactive Style
}

type TableOldTheme struct {
	Text Style
}

// OldTheme holds the default Styles and Colors for all widgets.
// You can set default widget Styles by modifying the OldTheme before creating the widgets.
var OldTheme = RootOldTheme{
	Default: NewStyle(ColorWhite),

	Block: BlockOldTheme{
		Title:  NewStyle(ColorWhite),
		Border: NewStyle(ColorWhite),
	},

	BarChart: BarChartOldTheme{
		Bars:   StandardColors,
		Nums:   StandardStyles,
		Labels: StandardStyles,
	},

	Paragraph: ParagraphOldTheme{
		Text: NewStyle(ColorWhite),
	},

	PieChart: PieChartOldTheme{
		Slices: StandardColors,
	},

	List: ListOldTheme{
		Text: NewStyle(ColorWhite),
	},

	Tree: TreeOldTheme{
		Text:      NewStyle(ColorWhite),
		Collapsed: COLLAPSED,
		Expanded:  EXPANDED,
	},

	StackedBarChart: StackedBarChartOldTheme{
		Bars:   StandardColors,
		Nums:   StandardStyles,
		Labels: StandardStyles,
	},

	Gauge: GaugeOldTheme{
		Bar:   ColorWhite,
		Label: NewStyle(ColorWhite),
	},

	Sparkline: SparklineOldTheme{
		Title: NewStyle(ColorWhite),
		Line:  ColorWhite,
	},

	Plot: PlotOldTheme{
		Lines: StandardColors,
		Axes:  ColorWhite,
	},

	Table: TableOldTheme{
		Text: NewStyle(ColorWhite),
	},

	Tab: TabOldTheme{
		Active:   NewStyle(ColorRed),
		Inactive: NewStyle(ColorWhite),
	},
}
