package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Scales the current drawing bigger or smaller
//     scaleWidth:  Scales the width of the current drawing (1=100%, 0.5=50%, 2=200%, etc.)
//     scaleHeight: Scales the height of the current drawing (1=100%, 0.5=50%, 2=200%, etc.)
//
//     The scale() method scales the current drawing, bigger or smaller.
//     Note: If you scale a drawing, all future drawings will also be scaled. The positioning will also be scaled. If
//     you scale(2,2); drawings will be positioned twice as far from the left and top of the canvas as you specify.
//     JavaScript syntax: context.scale(scalewidth, scaleheight);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.strokeRect(5, 5, 25, 15);
//     ctx.scale(2, 2);
//     ctx.strokeRect(5, 5, 25, 15);
func (el *Canvas) Scale(scaleWidth, scaleHeight iotmaker_types.Pixel) {
	el.SelfContext.Call("scale", scaleWidth, scaleHeight)
}
