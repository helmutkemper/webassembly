package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Returns an object that contains the width of the specified text
//     text: The text to be measured
//
//     The measureText() method returns an object that contains the width of the specified text, in pixels.
//     Tip: Use this method if you need to know the width of a text, before writing it on the canvas.
//     JavaScript syntax: context.measureText(text).width;
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.font = "30px Arial";
//     var txt = "Hello World"
//     ctx.fillText("width:" + ctx.measureText(txt).width, 10, 50)
//     ctx.fillText(txt, 10, 100);
func (el *Canvas) MeasureText(text string) {
	el.SelfContext.Call("measureText", text)
}
