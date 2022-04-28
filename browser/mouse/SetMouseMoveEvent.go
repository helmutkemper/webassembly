package mouse

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.channels-go/mouse"
	"syscall/js"
)

var mouseMoveEvt js.Func
var X int
var Y int

// SetMouseMoveEvent
//
// English:
//
//  Mouse move coupling function, passing (x, y) in mouse
//  channel.BrowserMouseToPlatformMouseCoordinate
//
// Português:
//
//  Função de acoplamento do mouse move, transmitindo (x, y) no canal
//  mouse.BrowserMouseToPlatformMouseCoordinate
func SetMouseMoveEvent() js.Func {
	mouseMoveEvt = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		X = e.Get("clientX").Int()
		Y = e.Get("clientY").Int()

		mouse.BrowserMouseToPlatformMouseCoordinate <- mouse.Move{X: X, Y: Y}

		return nil
	})

	return mouseMoveEvt
}

func ReleaseMouseMoveListener() {
	mouseMoveEvt.Release()
}
