package iotmaker_platform_webbrowser

import (
	"github.com/helmutkemper/iotmaker.platform.webbrowser/canvas"
)

func NewElement() canvas.Element {
	el := canvas.Element{}
	el.InitializeDocument()

	return el
}
