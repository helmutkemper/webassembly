package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Resets the current transform to the identity matrix. Then runs transform()
//     a: Scales the drawings horizontally
//     b: Skews the drawings horizontally
//     c: Skews the drawings vertically
//     d: Scales the drawings vertically
//     e: Moves the the drawings horizontally
//     f: Moves the the drawings vertically
//
//     Each object on the canvas has a current transformation matrix.
//     The setTransform() method resets the current transform to the identity matrix, and then runs transform() with the
//     same arguments.
//     In other words, the setTransform() method lets you scale, rotate, move, and skew the current context.
//     Note: The transformation will only affect drawings made after the setTransform method is called.
//     JavaScript syntax: context.setTransform(a, b, c, d, e, f);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.fillStyle = "yellow";
//     ctx.fillRect(0, 0, 250, 100)
//     ctx.setTransform(1, 0.5, -0.5, 1, 30, 10);
//     ctx.fillStyle = "red";
//     ctx.fillRect(0, 0, 250, 100);
//     ctx.setTransform(1, 0.5, -0.5, 1, 30, 10);
//     ctx.fillStyle = "blue";
//     ctx.fillRect(0, 0, 250, 100);
func (el *Canvas) SetTransform(a, b, c, d, e, f iotmaker_types.Coordinate) {
	el.SelfContext.Call("setTransform", a, b, c, d, e, f)
}
