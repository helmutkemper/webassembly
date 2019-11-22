package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Returns previously saved path state and attributes
func (el *Canvas) Restore() {
	el.selfDocument.Call("restore")
}
