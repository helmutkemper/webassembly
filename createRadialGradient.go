package iotmaker_platform_webbrowser

// en: Creates a radial gradient (to use on canvas content). The parameters represent two circles, one with its center
// at (x0, y0) and a radius of r0, and the other with its center at (x1, y1) with a radius of r1.
//     x0: The x-coordinate of the starting circle of the gradient
//     y0: The y-coordinate of the starting circle of the gradient
//     r0: The radius of the starting circle. Must be non-negative and finite. (note: radius is a width, not a degrees angle)
//     x1: The x-coordinate of the ending circle of the gradient
//     y1: The y-coordinate of the ending circle of the gradient
//     r1: The radius of the ending circle. Must be non-negative and finite. (note: radius is a width, not a degrees angle)
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
func (el *Canvas) CreateRadialGradient(x0, y0, r0, x1, y1, r1 interface{}) interface{} {
	return el.SelfContext.Call("createRadialGradient", x0, y0, r0, x1, y1, r1)
}
