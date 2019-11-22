package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

func (el *Canvas) Get(jsParam string, value ...interface{}) iotmaker_types.Coordinate {
	return el.selfDocument.Get(jsParam, value)
}
