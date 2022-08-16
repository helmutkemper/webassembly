//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
)

var canvas *html.TagCanvas

func main() {

	var img = factoryBrowser.NewTagImg().
		Alt("spacecraft").
		Src("./small.png", true).
		Width(29).
		Height(50)

	canvas = factoryBrowser.NewTagCanvas(800, 600).
		CreatePattern(img, html.KRepeatRuleRepeat).
		Rect(0, 0, 300, 300).
		FillStylePattern().
		Fill()

	var stage = factoryBrowser.NewStage()
	stage.Append(canvas)

	done := make(chan struct{}, 0)
	<-done
}
