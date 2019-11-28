package iotmaker_platform_webbrowser

import (
	"image/color"
	"syscall/js"
)

// en: Specifies the colors and stop positions in a gradient object
//     stop:  A value between 0.0 and 1.0 that represents the position between start and end in a gradient
//     color: A CSS color value to display at the stop position
//
//     The addColorStop() method specifies the colors and position in a gradient object.
//     The addColorStop() method is used together with createLinearGradient() or createRadialGradient().
//     Note: You can call the addColorStop() method multiple times to change a gradient. If you omit this method for
//     gradient objects, the gradient will not be visible. You need to create at least one color stop to have a visible
//     gradient.
//
//     Example:
//     var c = document.getElementById('myCanvas');
//     var ctx = c.getContext('2d');
//     var grd = ctx.createLinearGradient(0, 0, 170, 0);
//     grd.addColorStop(0, "black");
//     grd.addColorStop(1, "white");
//     ctx.fillStyle = grd;
//     ctx.fillRect(20, 20, 150, 100);
func (el *Canvas) AddColorStop(gradient interface{}, stop float64, color color.RGBA) {
	gradient.(js.Value).Call("addColorStop", stop, RGBAToJs(color))
}
