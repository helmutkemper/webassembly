package iotmaker_platform_webbrowser

// en: Fills the current drawing (path)
//     The fill() method fills the current drawing (path). The default color is black.
//     Tip: Use the fillStyle property to fill with another color/gradient.
//     Note: If the path is not closed, the fill() method will add a line from the last point to the startpoint of the
//     path to close the path (like closePath()), and then fill the path.
func (el *Canvas) Fill() {
	el.SelfContext.Call("fill")
}
