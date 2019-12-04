package mouse

import (
	"syscall/js"
)

type PointerCollisionFunction func(x, y int) bool
type PointerPositiveEventFunction func(x, y int, collision bool)
type PointerComplexFunction func(ritFunction PointerCollisionFunction, positiveEventFunction PointerPositiveEventFunction)

func AddFunctionPointer(collisionFunction PointerCollisionFunction, positiveEventFunction PointerPositiveEventFunction) int {
	if len(listCollisionFunctions) == 0 {
		listCollisionFunctions = make([]PointerCollisionFunction, 0)
		listPositEventFunctions = make([]PointerPositiveEventFunction, 0)
	}

	listCollisionFunctions = append(listCollisionFunctions, collisionFunction)
	listPositEventFunctions = append(listPositEventFunctions, positiveEventFunction)

	return len(listCollisionFunctions) - 1
}

var x, y int
var listCollisionFunctions []PointerCollisionFunction
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

		for k, f := range listCollisionFunctions {
			if f != nil {
				retValue := f(x, y)
				if listPositEventFunctions[k] != nil {
					listPositEventFunctions[k](x, y, retValue)
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
