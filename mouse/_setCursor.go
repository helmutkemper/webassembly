package mouse

import (
	"syscall/js"
)

var currentCursor CursorType = KCursorAuto

// Example:
// mouse.SetCursor(stage.SelfElement, mouse.KCursorAuto)
// mouse.SetCursor(stage.SelfElement, mouse.KCursorColResize)
func _SetCursor(element interface{}, cursor interface{}) {
	if cursor == currentCursor {
		return
	}

	currentCursor = cursor.(CursorType)

	element.(js.Value).Call("setAttribute", "style", cursor.(CursorType).String())
}
