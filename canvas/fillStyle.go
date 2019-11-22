package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Sets or returns the color, gradient, or pattern used to fill the drawing
//     The fillStyle property sets or returns the color, gradient, or pattern used to fill the drawing.
//     Default value:	#000000
//     JavaScript syntax:	context.fillStyle = color|gradient|pattern;
func (el *Canvas) FillStyle(value string) {
	el.SelfContext.Set("fillStyle", value)
}
