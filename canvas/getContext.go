package canvas

import (
	"syscall/js"
)

func (el *Canvas) GetContext() js.Value {
	return el.SelfContext
}
