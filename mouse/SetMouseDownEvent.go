package mouse

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.channels-go/mouse"
	"syscall/js"
)

var mouseDownEvt js.Func

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
