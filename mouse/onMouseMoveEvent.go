package mouse

import (
	"syscall/js"
)

var mouseMoveEvt js.Func
var X int
var Y int

func SetMouseMoveManager() js.Func {
	mouseMoveEvt = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		X = e.Get("clientX").Int()
		Y = e.Get("clientY").Int()

		return nil
	})

	return mouseMoveEvt
}

func ReleaseMouseMoveListener() {
	mouseMoveEvt.Release()
}
