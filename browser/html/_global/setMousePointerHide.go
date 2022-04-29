package _global

// SetMousePointerHide
//
// English:
//
//  Sets the mouse pointer to hide.
//
// Português:
//
//  Define o ponteiro do mouse como oculto.
func (e *GlobalAttributes) SetMousePointerHide() (ref *GlobalAttributes) {
	e.selfElement.Get("body").Set("style", mouse.KCursorNone.String())
	e.cursor = mouse.KCursorNone

	return e
}
