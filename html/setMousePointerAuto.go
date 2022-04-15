package html

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/browserMouse"

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
	//e.selfElement.Get("body").Set("style", mouse.KCursorAuto.String())
	e.cursor = browserMouse.KCursorAuto

	return e
}
