package mouse

import (
	"github.com/helmutkemper/iotmaker.platform/mouse"
	"syscall/js"
)

var currentCursor mouse.CursorType = mouse.KCursorAuto

// Example:
// mouse.SetCursor(stage.SelfElement, mouse.KCursorAuto)
// mouse.SetCursor(stage.SelfElement, mouse.KCursorColResize)
func SetCursor(element js.Value, cursor mouse.CursorType) {
	if cursor == currentCursor {
		return
	}

	currentCursor = cursor

	element.Call("setAttribute", "style", cursor.String())
}
