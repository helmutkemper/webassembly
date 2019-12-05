package mouse

import (
	"syscall/js"
)

var currentCursor CursorType = KCursorAuto

// Example:
// mouse.SetCursor(stage.SelfElement, mouse.KCursorAuto)
// mouse.SetCursor(stage.SelfElement, mouse.KCursorColResize)
func SetCursor(element js.Value, cursor CursorType) {
	if cursor == currentCursor {
		return
	}

	currentCursor = cursor

	element.Call("setAttribute", "style", cursor.String())
}
