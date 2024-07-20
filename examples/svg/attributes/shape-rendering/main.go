// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/shape-rendering
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/shape-rendering
//
//  <svg viewBox="0 0 420 200" xmlns="http://www.w3.org/2000/svg">
//    <circle cx="100" cy="100" r="100" shape-rendering="geometricPrecision"/>
//    <circle cx="320" cy="100" r="100" shape-rendering="crispEdges"/>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 420, 200}).Append(
		factoryBrowser.NewTagSvgCircle().Cx(100).Cy(100).R(100).ShapeRendering(html.KSvgShapeRenderingGeometricPrecision),
		factoryBrowser.NewTagSvgCircle().Cx(320).Cy(100).R(100).ShapeRendering(html.KSvgShapeRenderingCrispEdges),
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
