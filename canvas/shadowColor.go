package canvas

import (
	"image/color"
)

// en: Sets or returns the color to use for shadows
//     The shadowColor property sets or returns the color to use for shadows.
//     Note: Use the shadowColor property together with the shadowBlur property to create a shadow.
//     Tip: Adjust the shadow by using the shadowOffsetX and shadowOffsetY properties.
//     Default value: #000000
//     JavaScript syntax: context.shadowColor = color;
func (el *Canvas) ShadowColor(value color.RGBA) {
	el.SelfContext.Set("shadowColor", RGBAToJs(value))
}
