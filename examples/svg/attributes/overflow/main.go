// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/overflow
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/overflow
//
//  <svg viewBox="0 0 200 30" xmlns="http://www.w3.org/2000/svg" overflow="auto">
//    <text y="20">This text is wider than the SVG, so there should be a scrollbar shown.</text>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 200, 30}).Overflow(html.KOverflowAuto).Append(
		factoryBrowser.NewTagSvgText().Y(20).Text("This text is wider than the SVG, so there should be a scrollbar shown."),
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
