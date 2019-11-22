package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Sets or returns the current alignment for text content
//
//     The textAlign property sets or returns the current alignment for text content, according to the anchor point.
//     Normally, the text will START in the position specified, however, if you set textAlign="right" and place the text in position 150, it means that the text should END in position 150.
//     Tip: Use the fillText() or the strokeText() method to actually draw and position the text on the canvas.
//     Default value: start
//     JavaScript syntax: context.textAlign = "center | end | left | right | start";
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     // Create a red line in position 150
//     ctx.strokeStyle = "red";
//     ctx.moveTo(150, 20);
//     ctx.lineTo(150, 170);
//     ctx.stroke();
//     ctx.font = "15px Arial";
//     // Show the different textAlign values
//     ctx.textAlign = "start";
//     ctx.fillText("textAlign = start", 150, 60);
//     ctx.textAlign = "end";
//     ctx.fillText("textAlign = end", 150, 80);
//     ctx.textAlign = "left";
//     ctx.fillText("textAlign = left", 150, 100);
//     ctx.textAlign = "center";
//     ctx.fillText("textAlign = center", 150, 120);
//     ctx.textAlign = "right";
//     ctx.fillText("textAlign = right", 150, 140);
func (el *Canvas) TextAlign(value CanvasFontAlignRule) {
	el.SelfContext.Set("textAlign", value.String())
}
