package canvas

import (
	"syscall/js"
)

func (el *Canvas) GetCanvas() js.Value {
	return el.SelfElement
}
