package iotmaker_platform_webbrowser

import (
	"syscall/js"
)

// todo: selfContextType deve ser um enum
type Canvas struct {
	selfCanvas      js.Value
	selfContext     js.Value
	selfContextType int
	Element
}

func (el *Canvas) CreateNewWith3DContext(width, height float64) {
	el.selfCanvas = el.selfDocument.Call("getElementsById", "myCanvas")
	el.selfCanvas.Set("width", width)
	el.selfCanvas.Set("height", height)
	el.selfCanvas.Call("getContext", "3d")
}

// todo: tem que saber que id é um canvas
func (el *Canvas) InitializeContext2DById(id string) {
	el.Element.InitializeById(id)
	el.selfContextType = 1
	el.selfContext = el.selfCanvas.Call("getContext", "2d")
}

// todo: tem que saber que id é um canvas
func (el *Canvas) InitializeContext3DById(id string) {
	el.Element.InitializeById(id)
	el.selfContextType = 2
	el.selfContext = el.selfCanvas.Call("getContext", "3d")
}
