package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

func (el *Canvas) CreateEvent() {
	el.selfDocument.Call("createEvent")
}
