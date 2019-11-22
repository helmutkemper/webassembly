package iotmaker_platform_webbrowser

import (
	"github.com/helmutkemper/iotmaker.platform.webbrowser/canvas"
)

func NewExistentElementById(id string) canvas.Element {
	el := canvas.Element{}
	el.InitializeExistentElementById(id)

	return el
}
