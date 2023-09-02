// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package tooey

import (
	"image"
	"sync"

	"github.com/gdamore/tcell/v2"
)

// Drawable represents a renderable item
type Drawable interface {
	GetRect() image.Rectangle
	GetInnerRect() image.Rectangle
	// SetRect x1, y1, x2, y2
	SetRect(int, int, int, int)
	Draw(tcell.Screen)
	SetTheme(*Theme)
	sync.Locker
}

// Drawable represents an item that can be Rendered
//type Drawable interface {
//	GetRect() image.Rectangle
//	// SetRect x1, y1, x2, y2
//	SetRect(int, int, int, int)
//	X1() int
//	X2() int
//	Y1() int
//	Y2() int
//	DrawableWidth() int
//	DrawableHeight() int
//	Draw(tcell.Screen)
//	sync.Locker
//}

func Render(items ...Drawable) {
	for _, item := range items {
		item.Lock()
		item.Draw(scrn)
		item.Unlock()
	}
	scrn.Show()
}

//func Render2(items ...Drawable) {
//	for _, item := range items {
//		buf := NewBuffer(item.GetRect())
//		item.Lock()
//		item.Draw(buf)
//		item.Unlock()
//		//for point, cell := range buf.CellMap {
//		//	if point.In(buf.Rectangle) {
//		//		fmt.Println(cell)
//		//		//tb.SetCell(
//		//		//	point.X, point.Y,
//		//		//	cell.Rune,
//		//		//	tb.Attribute(cell.Style.Fg+1)|tb.Attribute(cell.Style.Modifier), tb.Attribute(cell.Style.Bg+1),
//		//		//)
//		//	}
//		//}
//	}
//	//tcell.Fl
//}
//
