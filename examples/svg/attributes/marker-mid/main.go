// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/marker-mid
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/marker-mid
//
//  <svg viewBox="0 0 120 120" xmlns="http://www.w3.org/2000/svg">
//    <defs>
//      <marker id="circle" markerWidth="8" markerHeight="8" refX="4" refY="4">
//          <circle cx="4" cy="4" r="4" stroke="none" fill="#f00"/>
//      </marker>
//    </defs>
//    <polyline fill="none" stroke="black"
//        points="20,100 40,60 70,80 100,20" marker-mid="url(#circle)"/>
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 120, 120}).Append(
		factoryBrowser.NewTagSvgDefs().Append(
			factoryBrowser.NewTagSvgMarker().Id("circle").MarkerWidth(8).MarkerHeight(8).RefX(4).RefY(4).Append(
				factoryBrowser.NewTagSvgCircle().Cx(4).Cy(4).R(4).Stroke(nil).Fill("#F00"),
			),
		),

		factoryBrowser.NewTagSvgPolyline().Fill(nil).Stroke(factoryColor.NewBlack()).Points([][]float64{{20, 100}, {40, 60}, {70, 80}, {100, 20}}).MarkerMid("url(#circle)"),
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
