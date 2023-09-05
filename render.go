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
	SetTheme(Theme)
	sync.Locker
}

// Render locks and draws the passed Drawables
func Render(items ...Drawable) {
	for _, item := range items {
		item.Lock()
		item.Draw(scrn)
		item.Unlock()
	}
	scrn.Show()
}
