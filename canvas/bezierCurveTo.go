package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Creates a cubic Bézier curve
//     cp1x: The x-axis coordinate of the first control point.
//     cp1y: The y-axis coordinate of the first control point.
//     cp2x: The x-axis coordinate of the second control point.
//     cp2y: The y-axis coordinate of the second control point.
//     x:    The x-axis coordinate of the end point.
//     y:    The y-axis coordinate of the end point.
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.beginPath();
//     ctx.moveTo(20, 20);
//     ctx.bezierCurveTo(20, 100, 200, 100, 200, 20);
//     ctx.stroke();
func (el *Canvas) BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y iotmaker_types.Pixel) {
	el.SelfContext.Call("bezierCurveTo", cp1x, cp1y, cp2x, cp2y, x, y)
}
