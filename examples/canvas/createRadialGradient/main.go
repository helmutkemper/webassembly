//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

var canvas *html.TagCanvas

func main() {

	canvas = factoryBrowser.NewTagCanvas(800, 600).
		CreateRadialGradient(75, 50, 5, 90, 60, 100).
		AddColorStopPosition(0.0, factoryColor.NewRed()).
		AddColorStopPosition(1.0, factoryColor.NewWhite()).
		FillStyleGradient().
		FillRect(10, 10, 150, 100)

	var stage = factoryBrowser.NewStage()
	stage.Append(canvas)

	done := make(chan struct{}, 0)
	<-done
}
