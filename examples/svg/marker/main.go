// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/marker
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/marker
//
//  <svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
//    <defs>
//      <!-- arrowhead marker definition -->
//      <marker id="arrow" viewBox="0 0 10 10" refX="5" refY="5"
//          markerWidth="6" markerHeight="6"
//          orient="auto-start-reverse">
//        <path d="M 0 0 L 10 5 L 0 10 z" />
//      </marker>
//
//      <!-- simple dot marker definition -->
//      <marker id="dot" viewBox="0 0 10 10" refX="5" refY="5"
//          markerWidth="5" markerHeight="5">
//        <circle cx="5" cy="5" r="5" fill="red" />
//      </marker>
//    </defs>
//
//    <!-- Coordinate axes with a arrowhead in both direction -->
//    <polyline points="10,10 10,90 90,90" fill="none" stroke="black"
//     marker-start="url(#arrow)" marker-end="url(#arrow)" />
//
//    <!-- Data line with polymarkers -->
//    <polyline points="15,80 29,50 43,60 57,30 71,40 85,15" fill="none" stroke="grey"
//     marker-start="url(#dot)" marker-mid="url(#dot)"  marker-end="url(#dot)" />
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
		ViewBox([]float64{0, 0, 100, 100}).
		Append(

			factoryBrowser.NewTagSvgDefs().
				Append(

					// arrowhead marker definition
					factoryBrowser.NewTagSvgMarker().
						Id("arrow").
						ViewBox([]float64{0, 0, 10, 10}).
						RefX(5).
						RefY(5).
						MarkerWidth(6).
						MarkerHeight(6).
						Orient(html.KSvgOrientAutoStartReverse).
						Append(

							factoryBrowser.NewTagSvgPath().
								D(

									factoryBrowser.NewPath().
										M(0, 0).
										L(10, 5).
										L(0, 10).
										Z(),
								),
						),

					// simple dot marker definition
					factoryBrowser.NewTagSvgMarker().
						Id("dot").
						ViewBox([]float64{0, 0, 10, 10}).
						RefX(5).
						RefY(5).
						MarkerWidth(5).
						MarkerHeight(5).
						Append(

							factoryBrowser.NewTagSvgCircle().
								Cx(5).
								Cy(5).
								R(5).
								Fill(factoryColor.NewRed()),
						),
				),

			// Coordinate axes with a arrowhead in both direction
			factoryBrowser.NewTagSvgPolyline().
				Points(

					factoryBrowser.NewPoints(
						[]html.Point{
							{10, 10}, {10, 90}, {90, 90},
						},
					),
				).
				Fill("none").
				Stroke(factoryColor.NewBlack()).
				MarkerStart("url(#arrow)").
				MarkerEnd("url(#arrow)"),

			// Data line with polymarkers
			factoryBrowser.NewTagSvgPolyline().
				Points(

					factoryBrowser.NewPoints(
						[]html.Point{
							{15, 80}, {29, 50}, {43, 60}, {57, 30}, {71, 40}, {85, 15},
						},
					),
				).
				Fill("none").
				Stroke(factoryColor.NewGrey()).
				MarkerStart("url(#dot)").
				MarkerMid("url(#dot)").
				MarkerEnd("url(#dot)"),
		)

	stage.Append(s1)

	<-done
}
