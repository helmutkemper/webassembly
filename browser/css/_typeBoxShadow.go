package css

import (
	"fmt"
	"image/color"
)

type BoxShadow struct {
	// X and Y offsets relative to the element, blur and spread radius, and color
	xOffset       int
	_xOffset      bool
	yOffset       int
	_yOffset      bool
	blur          int
	_blur         bool
	spreadRadius  int
	_spreadRadius bool
	color         color.RGBA
	_color        bool
	inset         Inset
	_inset        bool
}

func (b *BoxShadow) SetXOffset(xOffset int) {
	b.xOffset = xOffset
	b._xOffset = true
}

func (b *BoxShadow) SetYOffset(yOffset int) {
	b.yOffset = yOffset
	b._yOffset = true
}

func (b *BoxShadow) SetBlur(blur int) {
	b.blur = blur
	b._blur = true
}

func (b *BoxShadow) SetSpreadRadius(spreadRadius int) {
	b.spreadRadius = spreadRadius
	b._spreadRadius = true
}

func (b *BoxShadow) SetColor(color color.RGBA) {
	b.color = color
	b._color = true
}

func (b *BoxShadow) SetInset(inset Inset) {
	b.inset = inset
	b._inset = true
}

func (b BoxShadow) String() string {
	var inset string
	if b._inset {
		inset = string(b.inset) + " "
	}

	var xOffset string
	if b._xOffset {
		xOffset = fmt.Sprintf("%dpx ", b.xOffset)
	}

	var yOffset string
	if b._yOffset {
		yOffset = fmt.Sprintf("%dpx ", b.yOffset)
	}

	var blur string
	if b._blur {
		blur = fmt.Sprintf("%dpx ", b.blur)
	}

	var spreadRadius string
	if b._spreadRadius {
		spreadRadius = fmt.Sprintf("%dpx ", b.spreadRadius)
	}

	var colorStr string
	if b._color {
		colorStr = fmt.Sprintf("rgba(%d, %d, %d, %0.2f) ", b.color.R, b.color.G, b.color.B, float64(b.color.A)/255)
	}

	return fmt.Sprintf("%s%s%s%s%s%s", inset, xOffset, yOffset, blur, spreadRadius, colorStr)
}
