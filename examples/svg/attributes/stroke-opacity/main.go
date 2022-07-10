// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stroke-opacity
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stroke-opacity
//
//  <svg viewBox="0 0 40 10" xmlns="http://www.w3.org/2000/svg">
//    <!-- Default stroke opacity: 1 -->
//    <circle cx="5" cy="5" r="4" stroke="green" />
//
//    <!-- Stroke opacity as a number -->
//    <circle cx="15" cy="5" r="4" stroke="green"
//            stroke-opacity="0.7" />
//
//    <!-- Stroke opacity as a percentage -->
//    <circle cx="25" cy="5" r="4" stroke="green"
//            stroke-opacity="50%" />
//
//    <!-- Stroke opacity as a CSS property -->
//    <circle cx="35" cy="5" r="4" stroke="green"
//            style="stroke-opacity: .3;" />
//  </svg>

//go:build js
// +build js

// bug: o original não apararece nada

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 40, 10}).Append(
		// Default stroke opacity: 1
		factoryBrowser.NewTagSvgCircle().Cx(5).Cy(5).R(4).Stroke(factoryColor.NewGreen()),

		// Stroke opacity as a number
		factoryBrowser.NewTagSvgCircle().Cx(15).Cy(5).R(4).Stroke(factoryColor.NewGreen()).StrokeOpacity(0.7),

		// Stroke opacity as a percentage
		factoryBrowser.NewTagSvgCircle().Cx(25).Cy(5).R(4).Stroke(factoryColor.NewGreen()).StrokeOpacity(float32(0.5)),

		// Stroke opacity as a CSS property
		factoryBrowser.NewTagSvgCircle().Cx(35).Cy(5).R(4).Stroke(factoryColor.NewGreen()).Style("stroke-opacity: .3;"),
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
