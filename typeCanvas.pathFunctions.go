package iotmaker_platform_webbrowser

import (
	"syscall/js"
)

// en: Fills the current drawing (path)
//     The fill() method fills the current drawing (path). The default color is black.
//     Tip: Use the fillStyle property to fill with another color/gradient.
//     Note: If the path is not closed, the fill() method will add a line from the last point to the startpoint of the
//     path to close the path (like closePath()), and then fill the path.
func (el *Canvas) Fill() {
	el.selfContext.Call("fill")
}

// en: Actually draws the path you have defined
//     The stroke() method actually draws the path you have defined with all those moveTo() and lineTo() methods. The
//     default color is black.
//     Tip: Use the strokeStyle property to draw with another color/gradient.
func (el *Canvas) Stroke() {
	el.selfContext.Call("stroke")
}

//	en: Begins a path, or resets the current path
//      The beginPath() method begins a path, or resets the current path.
//      Tip: Use moveTo(), lineTo(), quadricCurveTo(), bezierCurveTo(), arcTo(), and arc(), to create paths.
//      Tip: Use the stroke() method to actually draw the path on the canvas.
func (el *Canvas) BeginPath() {
	el.selfContext.Call("beginPath")
}

// en: Moves the path to the specified point in the canvas, without creating a line
//     x: The x-coordinate of where to move the path to
//     y: The y-coordinate of where to move the path to
//     The moveTo() method moves the path to the specified point in the canvas, without creating a line.
//     Tip: Use the stroke() method to actually draw the path on the canvas.
func (el *Canvas) MoveTo(x, y float64) {
	el.selfContext.Call("moveTo", x, y)
}

// en: Creates a path from the current point back to the starting point
//     The closePath() method creates a path from the current point back to the starting point.
//     Tip: Use the stroke() method to actually draw the path on the canvas.
//     Tip: Use the fill() method to fill the drawing (black is default). Use the fillStyle property to fill with
//     another color/gradient.
func (el *Canvas) ClosePath(x, y float64) {
	el.selfContext.Call("closePath", x, y)
}

// en: Adds a new point and creates a line from that point to the last specified point in the canvas
//     x: The x-coordinate of where to create the line to
//     y: The y-coordinate of where to create the line to
//     The lineTo() method adds a new point and creates a line from that point to the last specified point in the canvas
//     (this method does not draw the line).
//     Tip: Use the stroke() method to actually draw the path on the canvas.
func (el *Canvas) LineTo(x, y float64) {
	el.selfContext.Call("lineTo", x, y)
}

// en: Clips a region of any shape and size from the original canvas
//     The clip() method clips a region of any shape and size from the original canvas.
//     Tip: Once a region is clipped, all future drawing will be limited to the clipped region (no access to other
//     regions on the canvas). You can however save the current canvas region using the save() method before using the
//     clip() method, and restore it (with the restore() method) any time in the future.
func (el *Canvas) Clip(x, y float64) {
	el.selfContext.Call("clip", x, y)
}

// en: Creates a quadratic Bézier curve
//     cpx: The x-axis coordinate of the control point.
//     cpy: The y-axis coordinate of the control point.
//     x:   The x-axis coordinate of the end point.
//     y:   The y-axis coordinate of the end point.
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.beginPath();
//     ctx.moveTo(20, 20);
//     ctx.quadraticCurveTo(20, 100, 200, 20);
//     ctx.stroke();
func (el *Canvas) QuadraticCurveTo(cpx, cpy, x, y float64) {
	el.selfContext.Call("quadraticCurveTo", cpx, cpy, x, y)
}

// en: Creates a cubic Bézier curve
//     cp1x: The x-axis coordinate of the first control point.
//     cp1y: The y-axis coordinate of the first control point.
//     cp2x: The x-axis coordinate of the second control point.
//     cp2y: The y-axis coordinate of the second control point.
//     x:    The x-axis coordinate of the end point.
//     y:    The y-axis coordinate of the end point.
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.beginPath();
//     ctx.moveTo(20, 20);
//     ctx.bezierCurveTo(20, 100, 200, 100, 200, 20);
//     ctx.stroke();
func (el *Canvas) BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y float64) {
	el.selfContext.Call("bezierCurveTo", cp1x, cp1y, cp2x, cp2y, x, y)
}

// en: Creates an arc/curve (used to create circles, or parts of circles)
//     x:             The horizontal coordinate of the arc's center.
//     y:             The vertical coordinate of the arc's center.
//     radius:        The arc's radius. Must be positive.
//     startAngle:    The angle at which the arc starts in radians, measured from the positive x-axis.
//     endAngle:      The angle at which the arc ends in radians, measured from the positive x-axis.
//     anticlockwise: [Optional] An optional Boolean. If true, draws the arc counter-clockwise between the start and end angles. The default is false (clockwise).
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.beginPath();
//     ctx.arc(100, 75, 50, 0, 2 * Math.PI);
//     ctx.stroke();
func (el *Canvas) Arc(x, y, radius, startAngle, endAngle float64, anticlockwise bool) {
	el.selfContext.Call("arc", x, y, radius, startAngle, endAngle, anticlockwise)
}

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
func (el *Canvas) ArcTo(x1, y1, x2, y2, radius float64) {
	el.selfContext.Call("arcTo", x1, y1, x2, y2, radius)
}

// en: Returns true if the specified point is in the current path, otherwise false
//     x: The x-axis coordinate of the point to check.
//     y: The y-axis coordinate of the point to check.
//     fillRule: The algorithm by which to determine if a point is inside or outside the path.
//          "nonzero": The non-zero winding rule. Default rule.
//          "evenodd": The even-odd winding rule.
//     path: A Path2D path to check against. If unspecified, the current path is used.
//
//    Example:
//    var c = document.getElementById("myCanvas");
//    var ctx = c.getContext("2d");
//    ctx.rect(20, 20, 150, 100);
//    if (ctx.isPointInPath(20, 50)) {
//      ctx.stroke();
//    };
func (el *Canvas) IsPointInPath(path js.Value, x, y float64, fillRule CanvasFillRule) {
	el.selfContext.Call("isPointInPath", path, x, y, fillRule.String())
}
