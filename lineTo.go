package iotmaker_platform_webbrowser

// en: Adds a new point and creates a line from that point to the last specified point in the canvas
//     x: The x-coordinate of where to create the line to
//     y: The y-coordinate of where to create the line to
//     The lineTo() method adds a new point and creates a line from that point to the last specified point in the canvas
//     (this method does not draw the line).
//     Tip: Use the stroke() method to actually draw the path on the canvas.
func (el *Canvas) LineTo(x, y int) {
	el.SelfContext.Call("lineTo", x, y)
}
