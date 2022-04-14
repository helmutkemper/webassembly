package browserMouse

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.channels-go/mouse"
	"syscall/js"
)

var mouseDoubleClickEvt js.Func

// SetMouseDoubleClickEvent
//
// English:
//
//  Mouse double click coupling function, passing (x, y) in mouse
//  channel.BrowserMouseDoubleClickToPlatformMouseDoubleClickEvent
//
// Português:
//
//  Função de acoplamento do clique duplo do mouse, transmitindo (x, y) no canal
//  mouse.BrowserMouseDoubleClickToPlatformMouseDoubleClickEvent
func SetMouseDoubleClickEvent() js.Func {
	mouseDoubleClickEvt = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		X = e.Get("clientX").Int()
		Y = e.Get("clientY").Int()

		mouse.BrowserMouseDoubleClickToPlatformMouseDoubleClickEvent <- mouse.DoubleClick{X: X, Y: Y}

		return nil
	})

	return mouseDoubleClickEvt
}

func ReleaseMouseDoubleClickEvent() {
	mouseDoubleClickEvt.Release()
}
