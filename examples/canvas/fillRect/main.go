//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
)

var canvas *html.TagCanvas

func main() {

	canvas = factoryBrowser.NewTagCanvas(800, 600).
		FillRect(20, 20, 150, 100)

	var stage = factoryBrowser.NewStage()
	stage.Append(canvas)

	done := make(chan struct{}, 0)
	<-done
}
