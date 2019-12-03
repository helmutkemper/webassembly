package mouse

import (
	"syscall/js"
)

type PointerRitFunction func(x, y int) uint8
type PointerPositiveEventFunction func(x, y int)
type PointerComplexFunction func(ritFunction PointerRitFunction, positiveEventFunction PointerPositiveEventFunction)

func AddFunctionPointer(ritFunction PointerRitFunction, positiveEventFunction PointerPositiveEventFunction) int {
	if len(listRitFunctions) == 0 {
		listRitFunctions = make([]PointerRitFunction, 0)
		listPositEventFunctions = make([]PointerPositiveEventFunction, 0)
	}

	listRitFunctions = append(listRitFunctions, ritFunction)
	listPositEventFunctions = append(listPositEventFunctions, positiveEventFunction)

	return len(listRitFunctions) - 1
}

var x, y int
var listRitFunctions []PointerRitFunction
var listPositEventFunctions []PointerPositiveEventFunction
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

		for k, f := range listRitFunctions {
			if f != nil {
				retValue := f(x, y)
				if retValue > 0 && listPositEventFunctions[k] != nil {
					listPositEventFunctions[k](x, y)
				}
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
