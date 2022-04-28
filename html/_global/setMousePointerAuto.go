package _global

// SetMousePointerAuto
//
// English:
//
//  Sets the mouse pointer to auto.
//
// Português:
//
//  Define o ponteiro do mouse como automático.
func (e *GlobalAttributes) SetMousePointerAuto() (ref *GlobalAttributes) {
	e.selfElement.Get("body").Set("style", mouse.KCursorAuto.String())
	e.cursor = mouse.KCursorAuto

	return e
}
