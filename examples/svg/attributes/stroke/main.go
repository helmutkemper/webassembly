// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stitchTiles
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stitchTiles
//
//  <svg viewBox="0 0 20 10" xmlns="http://www.w3.org/2000/svg">
//    <!-- Simple color stroke -->
//    <circle cx="5" cy="5" r="4" fill="none"
//            stroke="green" />
//
//    <!-- Stroke a circle with a gradient -->
//    <defs>
//      <linearGradient id="myGradient">
//        <stop offset="0%"   stop-color="green" />
//        <stop offset="100%" stop-color="white" />
//      </linearGradient>
//    </defs>
//
//    <circle cx="15" cy="5" r="4" fill="none"
//            stroke="url(#myGradient)" />
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
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 20, 10}).Append(
		// Simple color stroke
		factoryBrowser.NewTagSvgCircle().Cx(5).Cy(5).R(4).Fill(nil).Stroke(factoryColor.NewGreen()),

		// Stroke a circle with a gradient
		factoryBrowser.NewTagSvgDefs().Append(
			factoryBrowser.NewTagSvgLinearGradient().Id("myGradient").Append(
				factoryBrowser.NewTagSvgStop().Offset(float32(0)).StopColor(factoryColor.NewGreen()),
				factoryBrowser.NewTagSvgStop().Offset(float32(1)).StopColor(factoryColor.NewWhite()),
			),
		),

		factoryBrowser.NewTagSvgCircle().Cx(15).Cy(5).R(4).Fill(nil).Stroke("url(#myGradient)"),
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
