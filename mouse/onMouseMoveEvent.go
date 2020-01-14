package mouse

import (
	"fmt"
	"syscall/js"
)

var mouseMoveEvt js.Func
var X int
var Y int

var mouseChannelCoordinate chan []int

func SetMouseMoveManager() js.Func {
	mouseMoveEvt = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		X = e.Get("clientX").Int()
		Y = e.Get("clientY").Int()

		mouseChannelCoordinate <- []int{X, Y}

		return nil
	})

	return mouseMoveEvt
}

func ReleaseMouseMoveListener() {
	mouseMoveEvt.Release()
}

func init() {
	mouseChannelCoordinate = make(chan []int)

	go func() {
		//var coordinate = make( chan mouseChnn.CoordinateChannel )
		for {

			fmt.Printf("(%v)\n", <-mouseChannelCoordinate)

		}
	}()
}
