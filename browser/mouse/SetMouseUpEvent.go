package mouse

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.channels-go/mouse"
	"syscall/js"
)

var mouseUpEvt js.Func

// SetMouseUpEvent
//
// English:
//
//  Mouse up coupling function, passing (x, y) in mouse
//  channel.BrowserMouseUpToPlatformMouseUpEvent
//
// Português:
//
//  Função de acoplamento do mouse up, transmitindo (x, y) no canal
//  mouse.BrowserMouseUpToPlatformMouseUpEvent
func SetMouseUpEvent() js.Func {
	mouseUpEvt = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		X = e.Get("clientX").Int()
		Y = e.Get("clientY").Int()

		mouse.BrowserMouseUpToPlatformMouseUpEvent <- mouse.Release{X: X, Y: Y}

		return nil
	})

	return mouseUpEvt
}

func ReleaseMouseUpEvent() {
	mouseUpEvt.Release()
}
