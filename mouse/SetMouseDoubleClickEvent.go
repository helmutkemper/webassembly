package mouse

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.channels-go/mouse"
	"syscall/js"
)

var mouseDoubleClickEvt js.Func

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
