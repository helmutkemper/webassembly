package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

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
func (el *Canvas) IsPointInPath(path js.Value, x, y iotmaker_types.Coordinate, fillRule CanvasFillRule) bool {
	return el.SelfContext.Call("isPointInPath", path, x, y, fillRule.String()).Bool()
}
