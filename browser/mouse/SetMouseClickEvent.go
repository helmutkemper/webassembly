package mouse

import (
	mouse "github.com/helmutkemper/iotmaker.webassembly/platform/channels"
	"syscall/js"
)

var mouseClickEvt js.Func

// SetMouseClickEvent
//
// English:
//
//  Mouse click coupling function, passing (x, y) in mouse
//  channel.BrowserMouseClickToPlatformMouseClickEvent
//
// Português:
//
//  Função de acoplamento do clique do mouse, transmitindo (x, y) no canal
//  mouse.BrowserMouseClickToPlatformMouseClickEvent
func SetMouseClickEvent() js.Func {
	mouseClickEvt = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		X = e.Get("clientX").Int()
		Y = e.Get("clientY").Int()

		mouse.BrowserMouseClickToPlatformMouseClickEvent <- mouse.Click{X: X, Y: Y}

		return nil
	})

	return mouseClickEvt
}

func ReleaseMouseClickEvent() {
	mouseClickEvt.Release()
}
