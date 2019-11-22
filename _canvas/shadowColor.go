package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Sets or returns the color to use for shadows
//     The shadowColor property sets or returns the color to use for shadows.
//     Note: Use the shadowColor property together with the shadowBlur property to create a shadow.
//     Tip: Adjust the shadow by using the shadowOffsetX and shadowOffsetY properties.
//     Default value: #000000
//     JavaScript syntax: context.shadowColor = color;
func (el *Canvas) ShadowColor(value string) {
	el.SelfContext.Set("shadowColor", value)
}
