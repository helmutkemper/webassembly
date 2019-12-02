package mouse

import (
	"syscall/js"
)

type PointerFunction func(x, y int)

func AddFunctionPointer(f PointerFunction) int {
	if len(listFunctions) == 0 {
		listFunctions = make([]PointerFunction, 0)
	}

	listFunctions = append(listFunctions, f)

	return len(listFunctions) - 1
}

var x, y int
var listFunctions []PointerFunction
var mouseMoveEvt js.Func

func GetDefaultFunction() js.Func {
	StartMouseMoveListener()
	return mouseMoveEvt
}

func StartMouseMoveListener() js.Func {
	mouseMoveEvt = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		x = e.Get("clientX").Int()
		y = e.Get("clientY").Int()

		for _, f := range listFunctions {
			if f != nil {
				f(x, y)
			}
		}

		return nil
	})
	//mouseMoveEvt.Release()
	return mouseMoveEvt
}

func ReleaseMouseMoveListener() {
	mouseMoveEvt.Release()
}
