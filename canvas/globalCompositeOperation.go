package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Sets or returns how a new image are drawn onto an existing image
//
//     The globalCompositeOperation property sets or returns how a source (new) image are drawn onto a destination
//     (existing) image.
//     source image = drawings you are about to place onto the canvas.
//     destination image = drawings that are already placed onto the canvas.
//
//     Default value: source-over
//     JavaScript syntax: context.globalCompositeOperation = "source-in";
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.fillStyle = "red";
//     ctx.fillRect(20, 20, 75, 50);
//     ctx.globalCompositeOperation = "source-over";
//     ctx.fillStyle = "blue";
//     ctx.fillRect(50, 50, 75, 50);
//     ctx.fillStyle = "red";
//     ctx.fillRect(150, 20, 75, 50);
//     ctx.globalCompositeOperation = "destination-over";
//     ctx.fillStyle = "blue";
//     ctx.fillRect(180, 50, 75, 50);
func (el *Canvas) GlobalCompositeOperation(value CanvasCompositeOperationsRule) {
	el.SelfContext.Set("globalCompositeOperation", value.String())
}
