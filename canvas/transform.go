package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Replaces the current transformation matrix for the drawing
//     a: Scales the drawing horizontally
//     b: Skew the the drawing horizontally
//     c: Skew the the drawing vertically
//     d: Scales the drawing vertically
//     e: Moves the the drawing horizontally
//     f: Moves the the drawing vertically
//
//     Each object on the canvas has a current transformation matrix.
//     The transform() method replaces the current transformation matrix. It multiplies the current transformation
//     matrix with the matrix described by:
//
//     a | c | e
//    -----------
//     b | d | f
//    -----------
//     0 | 0 | 1
//
//     In other words, the transform() method lets you scale, rotate, move, and skew the current context.
//     Note: The transformation will only affect drawings made after the transform() method is called.
//     Note: The transform() method behaves relatively to other transformations made by rotate(), scale(), translate(),
//     or transform(). Example: If you already have set your drawing to scale by two, and the transform() method scales
//     your drawings by two, your drawings will now scale by four.
//     Tip: Check out the setTransform() method, which does not behave relatively to other transformations.
//     JavaScript syntax: context.transform(a, b, c, d, e, f);
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.fillStyle = "yellow";
//     ctx.fillRect(0, 0, 250, 100)
//     ctx.transform(1, 0.5, -0.5, 1, 30, 10);
//     ctx.fillStyle = "red";
//     ctx.fillRect(0, 0, 250, 100);
//     ctx.transform(1, 0.5, -0.5, 1, 30, 10);
//     ctx.fillStyle = "blue";
//     ctx.fillRect(0, 0, 250, 100);
func (el *Canvas) Transform(a, b, c, d, e, f iotmaker_types.Pixel) {
	el.SelfContext.Call("transform", a, b, c, d, e, f)
}
