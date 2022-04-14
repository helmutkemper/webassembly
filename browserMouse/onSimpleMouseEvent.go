package browserMouse

import (
	"syscall/js"
)

type SimpleManager func()

// SetMouseSimpleEventManager
//
// English:
//
//  Registers a Golang function to fire when a mouse event happens.
//
//   Input:
//     manager: Golang function like func().
//
// Português:
//
//  Registra uma função Golang para ser disparada quando um evento do mouse acontece.
//
//   Entrada:
//     manager: função Golang tipo func().
func SetMouseSimpleEventManager(manager SimpleManager) js.Func {
	mouseMoveEvt = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if manager != nil {
			manager()
		}

		return nil
	})

	return mouseMoveEvt
}
