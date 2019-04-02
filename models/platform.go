package models

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type PlatformType int

const (
	PLATFORM_TOP PlatformType = iota
	PLATFORM_BOTTOM
)

type Platform struct {
	Rect  pixel.Rect
	Color color.Color
	Type  PlatformType
}

func (p *Platform) Draw(imd *imdraw.IMDraw) {
	imd.Color = p.Color
	imd.Push(p.Rect.Min, p.Rect.Max)
	imd.Rectangle(0)
}
