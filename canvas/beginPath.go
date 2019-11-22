package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

//	en: Begins a path, or resets the current path
//      The beginPath() method begins a path, or resets the current path.
//      Tip: Use moveTo(), lineTo(), quadricCurveTo(), bezierCurveTo(), arcTo(), and arc(), to create paths.
//      Tip: Use the stroke() method to actually draw the path on the canvas.
func (el *Canvas) BeginPath() {
	el.SelfContext.Call("beginPath")
}
