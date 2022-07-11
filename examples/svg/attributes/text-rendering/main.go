// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/text-rendering
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/text-rendering
//
//  <svg viewBox="0 0 140 40" xmlns="http://www.w3.org/2000/svg">
//    <text y="15" text-rendering="geometricPrecision">Geometric precision</text>
//    <text y="35" text-rendering="optimizeLegibility">Optimized legibility</text>
//  </svg>

//go:build js
// +build js

// bug: o original não apararece nada

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 140, 40}).Append(
		factoryBrowser.NewTagSvgText().Y(15).TextRendering(html.KSvgTextRenderingGeometricPrecision).Text("Geometric precision"),
		factoryBrowser.NewTagSvgText().Y(35).TextRendering(html.KSvgTextRenderingOptimizeLegibility).Text("Optimized legibility"),
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
