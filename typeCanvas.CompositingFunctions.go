package iotmaker_platform_webbrowser

// en: Sets or returns the current alpha or transparency value of the drawing
//     number: The transparency value. Must be a number between 0.0 (fully transparent) and 1.0 (no transparancy)
//
//     Default value: 1.0
//     JavaScript syntax: context.globalAlpha = number;
//
//     The globalAlpha property sets or returns the current alpha or transparency value of the drawing.
//     The globalAlpha property value must be a number between 0.0 (fully transparent) and 1.0 (no transparancy)
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.fillStyle = "red";
//     ctx.fillRect(20, 20, 75, 50);
//     // Turn transparency on
//     ctx.globalAlpha = 0.2;
//     ctx.fillStyle = "blue";
//     ctx.fillRect(50, 50, 75, 50);
//     ctx.fillStyle = "green";
//     ctx.fillRect(80, 80, 75, 50);
func (el *Canvas) GlobalAlpha(value float64) {
	el.selfContext.Set("globalAlpha", value)
}

//globalCompositeOperation
// Sets or returns how a new image are drawn onto an existing image
