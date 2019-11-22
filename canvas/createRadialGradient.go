package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Creates a radial/circular gradient (to use on canvas content)
//     x0: The x-coordinate of the starting circle of the gradient
//     y0: The y-coordinate of the starting circle of the gradient
//     r0: The radius of the starting circle
//     x1: The x-coordinate of the ending circle of the gradient
//     y1: The y-coordinate of the ending circle of the gradient
//     r1: The radius of the ending circle
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
func (el *Canvas) CreateRadialGradient(x0, y0, r0, x1, y1, r1 iotmaker_types.Coordinate) {
	el.SelfContext.Call("createRadialGradient", x0, y0, r0, x1, y1, r1)
}
