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

// en: Sets or returns the vertical distance of the shadow from the shape
//     The shadowOffsetY property sets or returns the vertical distance of the shadow from the shape.
//     shadowOffsety = 0 indicates that the shadow is right behind the shape.
//     shadowOffsetY = 20 indicates that the shadow starts 20 pixels below the shape's top position.
//     shadowOffsetY = -20 indicates that the shadow starts 20 pixels above the shape's top position.
//     Tip: To adjust the horizontal distance of the shadow from the shape, use the shadowOffsetX property.
//     Default value: 0
//     JavaScript syntax: context.shadowOffsetY = number;
func (el *Canvas) ShadowOffsetY(value string) {
	el.selfContext.Set("shadowOffsetY", value)
}

// en: Creates a linear gradient (to use on canvas content)
//     x0: The x-coordinate of the start point of the gradient
//     y0: The y-coordinate of the start point of the gradient
//     x1: The x-coordinate of the end point of the gradient
//     y1: The y-coordinate of the end point of the gradient
//
//     The createLinearGradient() method creates a linear gradient object.
//     The gradient can be used to fill rectangles, circles, lines, text, etc.
//     Tip: Use this object as the value to the strokeStyle or fillStyle properties.
//     Tip: Use the addColorStop() method to specify different colors, and where to position the colors in the gradient object.
//     JavaScript syntax:	context.createLinearGradient(x0, y0, x1, y1);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     var grd = ctx.createLinearGradient(0, 0, 170, 0);
//     grd.addColorStop(0, "black");
//     grd.addColorStop(1, "white");
//     ctx.fillStyle = grd;
//     ctx.fillRect(20, 20, 150, 100);
func (el *Canvas) CreateLinearGradient(x0, y0, x1, y1 float64) {
	el.selfContext.Call("createLinearGradient", x0, y0, x1, y1)
}

// en: Repeats a specified element in the specified direction
//     image: Specifies the image, canvas, or video element of the pattern to use
//     repeatedElement
//          repeat: Default. The pattern repeats both horizontally and vertically
//          repeat-x: The pattern repeats only horizontally
//          repeat-y: The pattern repeats only vertically
//          no-repeat: The pattern will be displayed only once (no repeat)
//
//     The createPattern() method repeats the specified element in the specified direction.
//     The element can be an image, video, or another <canvas> element.
//     The repeated element can be used to draw/fill rectangles, circles, lines etc.
//     JavaScript syntax:	context.createPattern(image, "repeat|repeat-x|repeat-y|no-repeat");
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     var img = document.getElementById("lamp");
//     var pat = ctx.createPattern(img, "repeat");
//     ctx.rect(0, 0, 150, 100);
//     ctx.fillStyle = pat;
//     ctx.fill();
func (el *Canvas) CreatePattern(image js.Value, repeatRule CanvasRepeatRule) {
	el.selfContext.Call("createPattern", image, repeatRule)
}

// en: Creates a radial/circular gradient (to use on canvas content)
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     var grd = ctx.createRadialGradient(75, 50, 5, 90, 60, 100);
//     grd.addColorStop(0, "red");
//     grd.addColorStop(1, "white");
//     // Fill with gradient
//     ctx.fillStyle = grd;
//     ctx.fillRect(10, 10, 150, 100);
func (el *Canvas) CreateRadialGradient(image js.Value, repeatRule CanvasRepeatRule) {
	el.selfContext.Call("createRadialGradient", image, repeatRule)
}

//addColorStop()
// en: Specifies the colors and stop positions in a gradient object
