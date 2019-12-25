package mouse

import (
	"syscall/js"
)

type Manager func(x, y float64)

var mouseMoveEvt js.Func

func SetMouseMoveManager(manager Manager) js.Func {
	mouseMoveEvt = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		x := e.Get("clientX").Float()
		y := e.Get("clientY").Float()

		if manager != nil {
			manager(x, y)
		}

		return nil
	})
	//mouseMoveEvt.Release()
	return mouseMoveEvt
}

func ReleaseMouseMoveListener() {
	mouseMoveEvt.Release()
}
