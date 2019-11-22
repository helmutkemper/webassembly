package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Moves the path to the specified point in the canvas, without creating a line
//     x: The x-coordinate of where to move the path to
//     y: The y-coordinate of where to move the path to
//     The moveTo() method moves the path to the specified point in the canvas, without creating a line.
//     Tip: Use the stroke() method to actually draw the path on the canvas.
func (el *Canvas) MoveTo(x, y iotmaker_types.Coordinate) {
	el.SelfContext.Call("moveTo", x, y)
}
