package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

func (el *Canvas) GetCanvas() js.Value {
	return el.SelfElement
}
