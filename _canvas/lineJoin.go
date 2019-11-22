package canvas

// en: Sets or returns the type of corner created, when two lines meet
//     PlatformBasicType: "bevel|round|miter"
//
//     The lineJoin property sets or returns the type of corner created, when two lines meet.
//     Note: The "miter" value is affected by the miterLimit property.
//     Default value:	miter
//     JavaScript syntax:	context.lineJoin = "bevel|round|miter";
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.beginPath();
//     ctx.lineJoin = "round";
//     ctx.moveTo(20, 20);
//     ctx.lineTo(100, 50);
//     ctx.lineTo(20, 100);
//     ctx.stroke();
func (el *Canvas) LineJoin(value CanvasJoinRule) {
	el.SelfContext.Set("lineJoin", value.String())
}
