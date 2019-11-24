package iotmaker_platform_webbrowser

// en: Actually draws the path you have defined
//     The stroke() method actually draws the path you have defined with all those moveTo() and lineTo() methods. The
//     default color is black.
//     Tip: Use the strokeStyle property to draw with another color/gradient.
func (el *Canvas) Stroke() {
	el.SelfContext.Call("stroke")
}
