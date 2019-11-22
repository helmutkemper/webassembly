package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Sets or returns the style of the end caps for a line
//     PlatformBasicType: "butt|round|square"
//
//     The lineCap property sets or returns the style of the end caps for a line.
//     Note: The value "round" and "square" make the lines slightly longer.
//
//     Default value: butt
//     JavaScript syntax: context.lineCap = "butt|round|square";
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.beginPath();
//     ctx.lineCap = "round";
//     ctx.moveTo(20, 20);
//     ctx.lineTo(20, 200);
//     ctx.stroke();
func (el *Canvas) LineCap(value CanvasCapRule) {
	el.SelfContext.Set("lineCap", value.String())
}
