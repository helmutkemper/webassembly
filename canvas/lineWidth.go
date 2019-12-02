package canvas

// en: Sets or returns the current line width
//     PlatformBasicType: The current line width, in pixels
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
func (el *Canvas) LineWidth(value int) {
	el.SelfContext.Set("lineWidth", value)
}
