// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/line
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/line
//
//  <svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
//    <line x1="0" y1="80" x2="100" y2="20" stroke="black" />
//
//    <!-- If you do not specify the stroke
//         color the line will not be visible -->
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Append(
		factoryBrowser.NewTagSvgLine().X1(0).Y1(80).X2(100).Y2(20).Stroke(factoryColor.NewBlack()),
	)

	stage.Append(s1)

	<-done
}
