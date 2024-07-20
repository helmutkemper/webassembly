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
		FillStyle(factoryColor.NewRed()).
		FillRect(20, 20, 75, 50).
		GlobalAlpha(0.2).
		FillStyle(factoryColor.NewBlue()).
		FillRect(50, 50, 75, 50).
		FillStyle(factoryColor.NewGreen()).
		FillRect(80, 80, 75, 50)

	var stage = factoryBrowser.NewStage()
	stage.Append(canvas)

	done := make(chan struct{}, 0)
	<-done
}
