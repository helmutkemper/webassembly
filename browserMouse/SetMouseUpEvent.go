package browserMouse

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.channels-go/mouse"
	"syscall/js"
)

var mouseUpEvt js.Func

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
