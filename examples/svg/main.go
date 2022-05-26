//go:build js
// +build js

//
package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/stage"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	// browser stage
	var bs = stage.Stage{}
	bs.Init()

	factoryBrowser.NewTagSvg("svg").
		AppendToStage()

	factoryBrowser.NewTagSvgRect("rect").
		Width(200).
		Height(200).
		Fill(factoryColor.NewBlueHalfTransparent()).
		AppendById("svg")

	<-done
}
