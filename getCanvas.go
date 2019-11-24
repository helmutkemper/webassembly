package iotmaker_platform_webbrowser

import (
	"syscall/js"
)

func (el *Canvas) GetCanvas() js.Value {
	return el.SelfElement
}
