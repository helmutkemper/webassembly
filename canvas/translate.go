package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Remaps the (0,0) position on the canvas
//     x: The value to add to horizontal (x) coordinates
//     y: The value to add to vertical (y) coordinates
//
//     The translate() method remaps the (0,0) position on the canvas.
//     Note: When you call a method such as fillRect() after translate(), the value is added to the x- and y-coordinate
//     values.
//     JavaScript syntax: context.translate(x, y);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.fillRect(10, 10, 100, 50);
//     ctx.translate(70, 70);
//     ctx.fillRect(10, 10, 100, 50);
func (el *Canvas) Translate(x, y iotmaker_types.Pixel) {
	el.SelfContext.Call("translate", x, y)
}
