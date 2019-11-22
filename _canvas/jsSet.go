package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

func (el *Canvas) Set(jsParam string, value ...interface{}) {
	el.selfDocument.Set(jsParam, value)
}
