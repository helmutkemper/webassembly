package mouse

import (
	"syscall/js"
)

type SimpleManager func()

func SetMouseSimpleEventManager(manager SimpleManager) js.Func {
	mouseMoveEvt = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if manager != nil {
			manager()
		}

		return nil
	})

	return mouseMoveEvt
}
