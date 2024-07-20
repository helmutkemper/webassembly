package mouse

import (
	mouse "github.com/helmutkemper/webassembly/platform/channels"
	"syscall/js"
)

var mouseDownEvt js.Func

// SetMouseDownEvent
//
// English:
//
//  Mouse down coupling function, passing (x, y) in mouse
//  channel.BrowserMouseDownToPlatformMouseDownEvent
//
// Português:
//
//  Função de acoplamento do mouse down, transmitindo (x, y) no canal
//  mouse.BrowserMouseDownToPlatformMouseDownEvent
func SetMouseDownEvent() js.Func {
	mouseDownEvt = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		X = e.Get("clientX").Int()
		Y = e.Get("clientY").Int()

		mouse.BrowserMouseDownToPlatformMouseDownEvent <- mouse.Press{X: X, Y: Y}

		return nil
	})

	return mouseDownEvt
}

func ReleaseMousePressEvent() {
	mouseDownEvt.Release()
}
