package iotmaker_platform_webbrowser

import (
	"github.com/helmutkemper/iotmaker.platform.webbrowser/canvas"
)

func NewDocument() canvas.Document {
	el := canvas.Document{}
	el.Initialize()

	return el
}
