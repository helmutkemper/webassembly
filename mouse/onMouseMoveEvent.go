package mouse

import (
	"syscall/js"
)

type Manager func(x, y int)

var mouseMoveEvt js.Func

func SetMouseMoveManager(manager Manager) js.Func {
	mouseMoveEvt = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		x := e.Get("clientX").Int()
		y := e.Get("clientY").Int()

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
