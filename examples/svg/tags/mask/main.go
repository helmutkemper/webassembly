// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/mask
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/mask
//
//  <svg viewBox="-10 -10 120 120">
//    <mask id="myMask">
//      <!-- Everything under a white pixel will be visible -->
//      <rect x="0" y="0" width="100" height="100" fill="white" />
//
//      <!-- Everything under a black pixel will be invisible -->
//      <path d="M10,35 A20,20,0,0,1,50,35 A20,20,0,0,1,90,35 Q90,65,50,95 Q10,65,10,35 Z" fill="black" />
//    </mask>
//
//    <polygon points="-10,110 110,110 110,-10" fill="orange" />
//
//    <!-- with this mask applied, we "punch" a heart shape hole into the circle -->
//    <circle cx="50" cy="50" r="50" mask="url(#myMask)" />
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{-10, -10, 120, 120}).Append(

		factoryBrowser.NewTagSvgMask().Id("myMask").Append(

			//Everything under a white pixel will be visible
			factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(100).Height(100).Fill(factoryColor.NewWhite()),

			//Everything under a black pixel will be invisible
			factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(10, 35).A(20, 20, 0, 0, 1, 50, 35).A(20, 20, 0, 0, 1, 90, 35).Q(90, 65, 50, 95).Q(10, 65, 10, 35).Z()),
		),

		factoryBrowser.NewTagSvgPolygon().Points(factoryBrowser.NewPoints([]html.Point{{-10, 110}, {110, 110}, {110, -10}})).Fill(factoryColor.NewOrange()),

		//with this mask applied, we "punch" a heart shape hole into the circle
		factoryBrowser.NewTagSvgCircle().Cx(50).Cy(50).Cy(50).R(50).Mask("url(#myMask)"),
	)

	stage.Append(s1)

	<-done
}
