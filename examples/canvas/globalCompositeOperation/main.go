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
		FillStyle(factoryColor.NewRed()).
		FillRect(20, 20, 75, 50).
		GlobalCompositeOperation(html.KCompositeOperationsRuleSourceOver).
		FillStyle(factoryColor.NewBlue()).
		FillRect(50, 50, 75, 50).
		FillStyle(factoryColor.NewRed()).
		FillRect(150, 20, 75, 50).
		GlobalCompositeOperation(html.KCompositeOperationsRuleDestinationOver).
		FillStyle(factoryColor.NewBlue()).
		FillRect(180, 50, 75, 50)

	var stage = factoryBrowser.NewStage()
	stage.Append(canvas)

	done := make(chan struct{}, 0)
	<-done
}
