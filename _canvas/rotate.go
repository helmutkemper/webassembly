package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Rotates the current drawing
//     angle: The rotation angle, in radians.
//            To calculate from degrees to radians: degrees*Math.PI/180.
//            Example: to rotate 5 degrees, specify the following: 5*Math.PI/180
//
//     The rotate() method rotates the current drawing.
//     Note: The rotation will only affect drawings made AFTER the rotation is done.
//     JavaScript syntax: context.rotate(angle);
//
//     Example:
//     var c = document.getElementById("my Canvas");
//     var ctx = c.getContext("2d");
//     ctx.rotate(20 * Math.PI / 180);
//     ctx.fillRect(50, 20, 100, 50);
func (el *Canvas) Rotate(angle iotmaker_types.Coordinate) {
	el.SelfContext.Call("rotate", angle)
}
