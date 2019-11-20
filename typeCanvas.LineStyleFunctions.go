package iotmaker_platform_webbrowser

// en: Sets or returns the style of the end caps for a line
//     Value: "butt|round|square"
//
//     The lineCap property sets or returns the style of the end caps for a line.
//     Note: The value "round" and "square" make the lines slightly longer.
//
//     Default value: butt
//     JavaScript syntax: context.lineCap = "butt|round|square";
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.beginPath();
//     ctx.lineCap = "round";
//     ctx.moveTo(20, 20);
//     ctx.lineTo(20, 200);
//     ctx.stroke();
func (el *Canvas) LineCap(value CanvasCapRule) {
	el.SelfContext.Set("lineCap", value.String())
}

// en: Sets or returns the type of corner created, when two lines meet
//     Value: "bevel|round|miter"
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

// en: Sets or returns the current line width
//     Value: The current line width, in pixels
//
//     The lineWidth property sets or returns the current line width, in pixels.
//     Default value: 1
//     JavaScript syntax: context.lineWidth = number;
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.lineWidth = 10;
//     ctx.strokeRect(20, 20, 80, 100);
func (el *Canvas) LineWidth(value float64) {
	el.SelfContext.Set("lineWidth", value)
}

// en: Sets or returns the maximum miter length
//     Value: A positive number that specifies the maximum miter length. If the current miter length exceeds the
//            miterLimit, the corner will display as lineJoin "bevel"
//
//     The miterLimit property sets or returns the maximum miter length.
//     The miter length is the distance between the inner corner and the outer corner where two lines meet.
//
//     Default value: 10
//     JavaScript syntax: context.miterLimit = number;
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.lineWidth = 10;
//     ctx.lineJoin = "miter";
//     ctx.miterLimit = 5;
//     ctx.moveTo(20, 20);
//     ctx.lineTo(50, 27);
//     ctx.lineTo(20, 34);
//     ctx.stroke();
func (el *Canvas) MiterLimit(value float64) {
	el.SelfContext.Set("miterLimit", value)
}
