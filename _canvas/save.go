package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Saves the state of the current context
func (el *Canvas) Save() {
	el.selfDocument.Call("save")
}
