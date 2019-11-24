package iotmaker_platform_webbrowser

import (
	"syscall/js"
)

func (el *Canvas) GetContext() js.Value {
	return el.SelfContext
}
