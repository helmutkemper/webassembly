package html

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/browserMouse"

// SetMousePointer
//
// English:
//
//  Defines the shape of the mouse pointer.
//
//   Input:
//     value: mouse pointer shape.
//       Example: SetMousePointer(mouse.KCursorCell) // Use mouse.K... and let autocomplete do the
//                rest
//
// Português:
//
//  Define o formato do ponteiro do mouse.
//
//   Entrada:
//     V: formato do ponteiro do mouse.
//       Exemplo: SetMousePointer(mouse.KCursorCell) // Use mouse.K... e deixe o autocompletar fazer
//                o resto
func (e *GlobalAttributes) SetMousePointer(pointer browserMouse.CursorType) (ref *GlobalAttributes) {
	e.selfElement.Set("style", pointer.String())
	return e
}
