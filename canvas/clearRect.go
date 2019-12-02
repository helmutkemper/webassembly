package canvas

// en: Clears the specified pixels within a given rectangle
//     x:      The x-coordinate of the upper-left corner of the rectangle to clear
//     y:      The y-coordinate of the upper-left corner of the rectangle to clear
//     width:  The width of the rectangle to clear, in pixels
//     height: The height of the rectangle to clear, in pixels
//
//     The clearRect() method clears the specified pixels within a given rectangle.
//     JavaScript syntax: context.clearRect(x, y, width, height);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.fillStyle = "red";
//     ctx.fillRect(0, 0, 300, 150);
//     ctx.clearRect(20, 20, 100, 50);
func (el *Canvas) ClearRect(x, y, width, height int) {
	el.SelfContext.Call("clearRect", x, y, width, height)
}
