package canvas

// en: This method of the Canvas 2D API creates a gradient along the line connecting two given coordinates, starting at
// (x0, y0) point and ending at (x1, y1) point
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
func (el *Canvas) CreateLinearGradient(x0, y0, x1, y1 interface{}) interface{} {
	return el.SelfContext.Call("createLinearGradient", x0, y0, x1, y1)
}
