// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stroke-width
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stroke-width
//
//  <svg viewBox="0 0 30 10" xmlns="http://www.w3.org/2000/svg">
//    <!-- Default stroke width: 1 -->
//    <circle cx="5" cy="5" r="3" stroke="green" />
//
//    <!-- Stroke width as a number -->
//    <circle cx="15" cy="5" r="3" stroke="green"
//            stroke-width="3" />
//
//    <!-- Stroke width as a percentage -->
//    <circle cx="25" cy="5" r="3" stroke="green"
//            stroke-width="2%" />
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
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
		// Default stroke width: 1
		factoryBrowser.NewTagSvgCircle().Cx(5).Cy(5).R(3).Stroke(factoryColor.NewGreen()),

		// Stroke width as a number
		factoryBrowser.NewTagSvgCircle().Cx(15).Cy(5).R(3).Stroke(factoryColor.NewGreen()).StrokeWidth(3),

		// Stroke width as a percentage
		factoryBrowser.NewTagSvgCircle().Cx(25).Cy(5).R(3).Stroke(factoryColor.NewGreen()).StrokeWidth(float32(0.02)),
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
