// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/font-size
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/font-size
//
//  <svg viewBox="0 0 200 30" xmlns="http://www.w3.org/2000/svg">
//    <text y="20" font-size="smaller">smaller</text>
//    <text x="100" y="20" font-size="2em">2em</text>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 200, 30}).Append(
		factoryBrowser.NewTagSvgText().Y(20).FontSize("smaller").Text("smaller"),
		factoryBrowser.NewTagSvgText().X(100).Y(20).FontSize("2em").Text("2em"),
	)

	stage.Append(s1)

	<-done
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
