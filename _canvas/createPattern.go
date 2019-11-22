package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Repeats a specified element in the specified direction
//     image: Specifies the image, canvas, or video element of the pattern to use
//     repeatedElement
//          repeat: Default. The pattern repeats both horizontally and vertically
//          repeat-x: The pattern repeats only horizontally
//          repeat-y: The pattern repeats only vertically
//          no-repeat: The pattern will be displayed only once (no repeat)
//
//     The createPattern() method repeats the specified element in the specified direction.
//     The element can be an image, video, or another <canvas> element.
//     The repeated element can be used to draw/fill rectangles, circles, lines etc.
//     JavaScript syntax:	context.createPattern(image, "repeat|repeat-x|repeat-y|no-repeat");
//
//     Example:
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     var img = document.getElementById("lamp");
//     var pat = ctx.createPattern(img, "repeat");
//     ctx.rect(0, 0, 150, 100);
//     ctx.fillStyle = pat;
//     ctx.fill();
func (el *Canvas) CreatePattern(image js.Value, repeatRule CanvasRepeatRule) {
	el.SelfContext.Call("createPattern", image, repeatRule)
}
