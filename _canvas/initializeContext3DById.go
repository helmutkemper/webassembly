package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// todo: tem que saber que id é um canvas
func (el *Canvas) InitializeContext3DById(id string) {
	el.Element.NewCanvas(id)
	el.SelfContextType = 2
	el.SelfContext = el.SelfElement.Call("getContext", "3d")
}
