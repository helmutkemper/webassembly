package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Sets or returns the color, gradient, or pattern used for strokes
//     The strokeStyle property sets or returns the color, gradient, or pattern used for strokes.
//     Default value: #000000
//     JavaScript syntax: context.strokeStyle = color|gradient|pattern;
func (el *Canvas) StrokeStyle(value string) {
	el.SelfContext.Set("strokeStyle", value)
}
