package iotmaker_platform_webbrowser

// en: Sets or returns the color, gradient, or pattern used to fill the drawing
//     The fillStyle property sets or returns the color, gradient, or pattern used to fill the drawing.
//     Default value:	#000000
//     JavaScript syntax:	context.fillStyle = color|gradient|pattern;
func (el *Canvas) FillStyle(value interface{}) {
	el.SelfContext.Set("fillStyle", value)
}
