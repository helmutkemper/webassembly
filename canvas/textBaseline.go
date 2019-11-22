package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Sets or returns the current text baseline used when drawing text
//     PlatformBasicType:
//          alphabetic:  Default. The text baseline is the normal alphabetic baseline
//          top:         The text baseline is the top of the em square
//          hanging:     The text baseline is the hanging baseline
//          middle:      The text baseline is the middle of the em square
//          ideographic: The text baseline is the ideographic baseline
//          bottom:      The text baseline is the bottom of the bounding box
//
//     The textBaseline property sets or returns the current text baseline used when drawing text.
//     Note: The fillText() and strokeText() methods will use the specified textBaseline value when positioning the text
//     on the canvas.
//     Default value: alphabetic
//     JavaScript syntax: context.textBaseline = "alphabetic|top|hanging|middle|ideographic|bottom";
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     //Draw a red line at y=100
//     ctx.strokeStyle = "red";
//     ctx.moveTo(5, 100);
//     ctx.lineTo(395, 100);
//     ctx.stroke();
//     ctx.font = "20px Arial"
//     //Place each word at y=100 with different textBaseline values
//     ctx.textBaseline = "top";
//     ctx.fillText("Top", 5, 100);
//     ctx.textBaseline = "bottom";
//     ctx.fillText("Bottom", 50, 100);
//     ctx.textBaseline = "middle";
//     ctx.fillText("Middle", 120, 100);
//     ctx.textBaseline = "alphabetic";
//     ctx.fillText("Alphabetic", 190, 100);
//     ctx.textBaseline = "hanging";
//     ctx.fillText("Hanging", 290, 100);
func (el *Canvas) TextBaseline(value CanvasTextBaseLineRule) {
	el.SelfContext.Set("textBaseline", value.String())
}
