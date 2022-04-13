package browserMouse

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.channels-go/mouse"
	"syscall/js"
)

var mouseClickEvt js.Func

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
