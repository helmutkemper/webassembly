package iotmaker_platform_webbrowser

import (
	"syscall/js"
)

// en: Sets or returns the color, gradient, or pattern used to fill the drawing
//     The fillStyle property sets or returns the color, gradient, or pattern used to fill the drawing.
//     Default value:	#000000
//     JavaScript syntax:	context.fillStyle = color|gradient|pattern;
func (el *Canvas) FillStyle(value string) {
	el.selfContext.Set("fillStyle", value)
}

// en: Sets or returns the color, gradient, or pattern used for strokes
//     The strokeStyle property sets or returns the color, gradient, or pattern used for strokes.
//     Default value: #000000
//     JavaScript syntax: context.strokeStyle = color|gradient|pattern;
func (el *Canvas) StrokeStyle(value string) {
	el.selfContext.Set("strokeStyle", value)
}

// en: Sets or returns the color to use for shadows
//     The shadowColor property sets or returns the color to use for shadows.
//     Note: Use the shadowColor property together with the shadowBlur property to create a shadow.
//     Tip: Adjust the shadow by using the shadowOffsetX and shadowOffsetY properties.
//     Default value: #000000
//     JavaScript syntax: context.shadowColor = color;
func (el *Canvas) ShadowColor(value string) {
	el.selfContext.Set("shadowColor", value)
}

// en: Sets or returns the blur level for shadows
//     The shadowBlur property sets or returns the blur level for shadows.
//     Default value: 0
//     JavaScript syntax: context.shadowBlur = number;
func (el *Canvas) ShadowBlur(value string) {
	el.selfContext.Set("shadowBlur", value)
}

// en: Sets or returns the horizontal distance of the shadow from the shape
//     The shadowOffsetX property sets or returns the horizontal distance of the shadow from the shape.
//     shadowOffsetX = 0 indicates that the shadow is right behind the shape.
//     shadowOffsetX = 20 indicates that the shadow starts 20 pixels to the right (from the shape's left position).
//     shadowOffsetX = -20 indicates that the shadow starts 20 pixels to the left (from the shape's left position).
//     Tip: To adjust the vertical distance of the shadow from the shape, use the shadowOffsetY property.
//     Default value: 0
//     JavaScript syntax: context.shadowOffsetX = number;
func (el *Canvas) ShadowOffsetX(value string) {
	el.selfContext.Set("shadowOffsetX", value)
}

//shadowOffsetY
// en: Sets or returns the vertical distance of the shadow from the shape
//createLinearGradient()
// en: Creates a linear gradient (to use on canvas content)
//createPattern()
// en: Repeats a specified element in the specified direction
//createRadialGradient()
// en: Creates a radial/circular gradient (to use on canvas content)
//addColorStop()
// en: Specifies the colors and stop positions in a gradient object
