//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"math"
)

var canvas *html.TagCanvas

func main() {

	canvas = factoryBrowser.NewTagCanvas(800, 600).
		BeginPath().
		Arc(100, 75, 50, 0, 2*math.Pi, false).
		Stroke()

	var stage = factoryBrowser.NewStage()
	stage.Append(canvas)

	done := make(chan struct{}, 0)
	<-done
}
