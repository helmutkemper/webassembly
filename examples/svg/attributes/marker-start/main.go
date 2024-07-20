// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/marker-start
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/marker-start
//
//  <svg viewBox="0 0 120 120" xmlns="http://www.w3.org/2000/svg">
//    <defs>
//      <marker id="triangle" viewBox="0 0 10 10"
//            refX="1" refY="5"
//            markerUnits="strokeWidth"
//            markerWidth="10" markerHeight="10"
//            orient="auto">
//        <path d="M 0 0 L 10 5 L 0 10 z" fill="#f00"/>
//      </marker>
//    </defs>
//    <polyline fill="none" stroke="black"
//        points="20,100 40,60 70,80 100,20" marker-start="url(#triangle)"/>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 120, 120}).Append(
		factoryBrowser.NewTagSvgDefs().Append(
			factoryBrowser.NewTagSvgMarker().Id("triangle").ViewBox([]float64{0, 0, 10, 10}).RefX(1).RefY(5).MarkerUnits(html.KSvgMarkerUnitsStrokeWidth).MarkerWidth(10).MarkerHeight(10).Orient(html.KSvgOrientAuto).Append(
				factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(0, 0).L(10, 5).L(0, 10).Z()).Fill("#F00"),
			),
		),

		factoryBrowser.NewTagSvgPolyline().Fill(nil).Stroke(factoryColor.NewBlack()).Points([][]float64{{20, 100}, {40, 60}, {70, 80}, {100, 20}}).MarkerStart("url(#triangle)"),
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
