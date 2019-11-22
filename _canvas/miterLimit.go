package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Sets or returns the maximum miter length
//     PlatformBasicType: A positive number that specifies the maximum miter length. If the current miter length exceeds the
//            miterLimit, the corner will display as lineJoin "bevel"
//
//     The miterLimit property sets or returns the maximum miter length.
//     The miter length is the distance between the inner corner and the outer corner where two lines meet.
//
//     Default value: 10
//     JavaScript syntax: context.miterLimit = number;
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.lineWidth = 10;
//     ctx.lineJoin = "miter";
//     ctx.miterLimit = 5;
//     ctx.moveTo(20, 20);
//     ctx.lineTo(50, 27);
//     ctx.lineTo(20, 34);
//     ctx.stroke();
func (el *Canvas) MiterLimit(value iotmaker_types.Coordinate) {
	el.SelfContext.Set("miterLimit", value)
}
