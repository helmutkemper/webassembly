package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Creates a quadratic Bézier curve
//     cpx: The x-axis coordinate of the control point.
//     cpy: The y-axis coordinate of the control point.
//     x:   The x-axis coordinate of the end point.
//     y:   The y-axis coordinate of the end point.
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.beginPath();
//     ctx.moveTo(20, 20);
//     ctx.quadraticCurveTo(20, 100, 200, 20);
//     ctx.stroke();
func (el *Canvas) QuadraticCurveTo(cpx, cpy, x, y iotmaker_types.Coordinate) {
	el.SelfContext.Call("quadraticCurveTo", cpx, cpy, x, y)
}
