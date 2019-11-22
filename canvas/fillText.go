package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Draws "filled" text on the canvas
//     text:     Specifies the text that will be written on the canvas
//     x:        The x coordinate where to start painting the text (relative to the canvas)
//     y:        The y coordinate where to start painting the text (relative to the canvas)
//     maxWidth: Optional. The maximum allowed width of the text, in pixels
//
//     The fillText() method draws filled text on the canvas. The default color of the text is black.
//     Tip: Use the font property to specify font and font size, and use the fillStyle property to render the text in another color/gradient.
//     JavaScript syntax: context.fillText(text, x, y, maxWidth);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.font = "20px Georgia";
//     ctx.fillText("Hello World!", 10, 50);
//     ctx.font = "30px Verdana";
//     // Create gradient
//     var gradient = ctx.createLinearGradient(0, 0, c.width, 0);
//     gradient.addColorStop("0", "magenta");
//     gradient.addColorStop("0.5", "blue");
//     gradient.addColorStop("1.0", "red");
//     // Fill with gradient
//     ctx.fillStyle = gradient;
//     ctx.fillText("Big smile!", 10, 90);
func (el *Canvas) FillText(text string, x, y, maxWidth iotmaker_types.Coordinate) {
	el.SelfContext.Call("fillText", text, x, y, maxWidth)
}
