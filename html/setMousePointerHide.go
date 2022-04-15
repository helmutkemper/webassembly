package html

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/browserMouse"

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
	//e.selfElement.Get("body").Set("style", mouse.KCursorNone.String())
	e.cursor = browserMouse.KCursorNone

	return e
}
