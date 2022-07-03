// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/polyline
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/polyline
//
//  <svg viewBox="0 0 200 100" xmlns="http://www.w3.org/2000/svg">
//    <!-- Example of a polyline with the default fill -->
//    <polyline points="0,100 50,25 50,75 100,0" />
//
//    <!-- Example of the same polyline shape with stroke and no fill -->
//    <polyline points="100,100 150,25 150,75 200,0"
//              fill="none" stroke="black" />
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().
		ViewBox([]float64{0, 0, 200, 100}).
		Append(

			// Example of a polyline with the default fill
			factoryBrowser.NewTagSvgPolyline().Points(factoryBrowser.NewPoints([]html.Point{{0, 100}, {50, 25}, {50, 75}, {100, 0}})),

			// Example of the same polyline shape with stroke and no fill
			factoryBrowser.NewTagSvgPolyline().Points(factoryBrowser.NewPoints([]html.Point{{100, 100}, {150, 25}, {150, 75}, {200, 0}})).Fill(nil).Stroke(factoryColor.NewBlack()),
		)

	stage.Append(s1)

	<-done
}
