package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
)

// en: Creates an arc/curve between two tangents
//     x1:     The x-axis coordinate of the first control point.
//     y1:     The y-axis coordinate of the first control point.
//     x2:     The x-axis coordinate of the second control point.
//     y2:     The y-axis coordinate of the second control point.
//     radius: The arc's radius. Must be non-negative.
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.beginPath();
//     ctx.moveTo(20, 20);              // Create a starting point
//     ctx.lineTo(100, 20);             // Create a horizontal line
//     ctx.arcTo(150, 20, 150, 70, 50); // Create an arc
//     ctx.lineTo(150, 120);            // Continue with vertical line
//     ctx.stroke();                    // Draw it
func (el *Canvas) ArcTo(x1, y1, x2, y2, radius iotmaker_types.Coordinate) {
	el.SelfContext.Call("arcTo", x1, y1, x2, y2, radius)
}
