//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
)

var canvas *html.TagCanvas

func main() {

	canvas = factoryBrowser.NewTagCanvas(800, 600).
		CreateLinearGradient(0, 0, 170, 0).
		AddColorStopPosition(0.0, factoryColor.NewBlack()).
		AddColorStopPosition(0.5, factoryColor.NewOrangered()).
		AddColorStopPosition(1.0, factoryColor.NewWhite()).
		FillStyleGradient().
		FillRect(20, 20, 150, 100)

	var stage = factoryBrowser.NewStage()
	stage.Append(canvas)

	done := make(chan struct{}, 0)
	<-done
}
