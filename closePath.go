package iotmaker_platform_webbrowser

// en: Creates a path from the current point back to the starting point
//     The closePath() method creates a path from the current point back to the starting point.
//     Tip: Use the stroke() method to actually draw the path on the canvas.
//     Tip: Use the fill() method to fill the drawing (black is default). Use the fillStyle property to fill with
//     another color/gradient.
func (el *Canvas) ClosePath(x, y int) {
	el.SelfContext.Call("closePath", x, y)
}
