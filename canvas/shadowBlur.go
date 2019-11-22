package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Sets or returns the blur level for shadows
//     The shadowBlur property sets or returns the blur level for shadows.
//     Default value: 0
//     JavaScript syntax: context.shadowBlur = number;
func (el *Canvas) ShadowBlur(value string) {
	el.SelfContext.Set("shadowBlur", value)
}
