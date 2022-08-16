// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/font-size-adjust
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/font-size-adjust
//
//  <svg width="600" height="80" viewBox="0 0 500 80"
//      xmlns="http://www.w3.org/2000/svg">
//    <text y="20" font-family="Times, serif" font-size="10px">
//      This text uses the Times font (10px), which is hard to read in small sizes.
//    </text>
//    <text y="40" font-family="Verdana, sans-serif" font-size="10px">
//      This text uses the Verdana font (10px), which has relatively large lowercase letters.
//    </text>
//    <text y="60" font-family="Times, serif" font-size="10px" font-size-adjust="0.58">
//      This is the 10px Times, but now adjusted to the same aspect ratio as the Verdana.
//    </text>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().Width(600).Height(80).ViewBox([]float64{0, 0, 500, 80}).Append(
		factoryBrowser.NewTagSvgText().Y(20).FontFamily("Times, serif").FontSize("10px").Text("This text uses the Times font (10px), which is hard to read in small sizes."),
		factoryBrowser.NewTagSvgText().Y(40).FontFamily("Verdana, sans-serif").FontSize("10px").Text("This text uses the Verdana font (10px), which has relatively large lowercase letters."),
		factoryBrowser.NewTagSvgText().Y(60).FontFamily("Times, serif").FontSize("10px").FontSizeAdjust(0.58).Text("This is the 10px Times, but now adjusted to the same aspect ratio as the Verdana."),
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
