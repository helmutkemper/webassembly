package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Creates an arc/curve (used to create circles, or parts of circles)
//     x:             The horizontal coordinate of the arc's center.
//     y:             The vertical coordinate of the arc's center.
//     radius:        The arc's radius. Must be positive.
//     startAngle:    The angle at which the arc starts in radians, measured from the positive x-axis.
//     endAngle:      The angle at which the arc ends in radians, measured from the positive x-axis.
//     anticlockwise: [Optional] An optional Boolean. If true, draws the arc counter-clockwise between the start and end angles. The default is false (clockwise).
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.beginPath();
//     ctx.arc(100, 75, 50, 0, 2 * Math.PI);
//     ctx.stroke();
func (el *Canvas) Arc(x, y, radius, startAngle, endAngle iotmaker_types.Pixel, anticlockwise bool) {
	el.SelfContext.Call("arc", x, y, radius, startAngle, endAngle, anticlockwise)
}
