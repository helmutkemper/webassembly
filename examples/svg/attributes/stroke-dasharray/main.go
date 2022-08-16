// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stroke-dasharray
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stroke-dasharray
//
//  <svg viewBox="0 0 30 10" xmlns="http://www.w3.org/2000/svg">
//    <!-- No dashes nor gaps -->
//    <line x1="0" y1="1" x2="30" y2="1" stroke="black" />
//
//    <!-- Dashes and gaps of the same size -->
//    <line x1="0" y1="3" x2="30" y2="3" stroke="black"
//            stroke-dasharray="4" />
//
//    <!-- Dashes and gaps of different sizes -->
//    <line x1="0" y1="5" x2="30" y2="5" stroke="black"
//            stroke-dasharray="4 1" />
//
//    <!-- Dashes and gaps of various sizes with an odd number of values -->
//    <line x1="0" y1="7" x2="30" y2="7" stroke="black"
//            stroke-dasharray="4 1 2" />
//
//    <!-- Dashes and gaps of various sizes with an even number of values -->
//    <line x1="0" y1="9" x2="30" y2="9" stroke="black"
//            stroke-dasharray="4 1 2 3" />
//  </svg>

//go:build js

// bug: o original não apararece nada

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
		// No dashes nor gaps
		factoryBrowser.NewTagSvgLine().X1(0).Y1(1).X2(30).Y2(1).Stroke(factoryColor.NewBlack()),

		// Dashes and gaps of the same size
		factoryBrowser.NewTagSvgLine().X1(0).Y1(3).X2(30).Y2(3).Stroke(factoryColor.NewBlack()).StrokeDasharray(4),

		// Dashes and gaps of different sizes
		factoryBrowser.NewTagSvgLine().X1(0).Y1(5).X2(30).Y2(5).Stroke(factoryColor.NewBlack()).StrokeDasharray([]float64{4, 1}),

		// Dashes and gaps of various sizes with an odd number of values
		factoryBrowser.NewTagSvgLine().X1(0).Y1(7).X2(30).Y2(7).Stroke(factoryColor.NewBlack()).StrokeDasharray([]float64{4, 1, 2}),

		// Dashes and gaps of various sizes with an even number of values
		factoryBrowser.NewTagSvgLine().X1(0).Y1(9).X2(30).Y2(9).Stroke(factoryColor.NewBlack()).StrokeDasharray([]float64{4, 1, 2, 3}),
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
