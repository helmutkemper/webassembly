package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// todo: tem que saber que id é um canvas
func (el *Canvas) InitializeContext2DById(id string) {
	el.Document.Initialize()
	el.SelfElement = el.Element.NewCanvas(id)
	el.SelfContextType = 1
	el.SelfContext = el.SelfElement.Call("getContext", "2d")
}
