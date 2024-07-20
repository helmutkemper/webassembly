// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/d
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/d
//
//  <svg viewBox="0 0 200 100" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
//
//    <!-- Quadratic Bézier curve with implicit repetition -->
//    <path fill="none" stroke="red"
//          d="M 10,50
//             Q 25,25 40,50
//             t 30,0 30,0 30,0 30,0 30,0" />
//
//    <!-- Highlight the curve vertex and control points -->
//    <g>
//      <polyline points="10,50 25,25 40,50" stroke="rgba(0,0,0,.2)" fill="none" />
//      <circle cx="25" cy="25" r="1.5" />
//
//      <!-- Curve vertex points -->
//      <circle cx="10" cy="50" r="1.5"/>
//      <circle cx="40" cy="50" r="1.5"/>
//
//      <g id="SmoothQuadraticDown">
//        <polyline points="40,50 55,75 70,50" stroke="rgba(0,0,0,.2)" stroke-dasharray="2" fill="none" />
//        <circle cx="55" cy="75" r="1.5" fill="lightgrey" />
//        <circle cx="70" cy="50" r="1.5" />
//      </g>
//
//      <g id="SmoothQuadraticUp">
//        <polyline points="70,50 85,25 100,50" stroke="rgba(0,0,0,.2)" stroke-dasharray="2" fill="none" />
//        <circle cx="85" cy="25" r="1.5" fill="lightgrey" />
//        <circle cx="100" cy="50" r="1.5" />
//      </g>
//
//      <use xlink:href="#SmoothQuadraticDown" x="60" />
//      <use xlink:href="#SmoothQuadraticUp"   x="60" />
//      <use xlink:href="#SmoothQuadraticDown" x="120" />
//    </g>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
	"image/color"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().
		ViewBox([]float64{0, 0, 200, 100}).XmlnsXLink("http://www.w3.org/1999/xlink").Append(

		// Quadratic Bézier curve with implicit repetition
		factoryBrowser.NewTagSvgPath().Fill(nil).Stroke(factoryColor.NewRed()).D(factoryBrowser.NewPath().M(10, 50).Q(25, 25, 40, 50).Td(30, 0).Td(30, 0).Td(30, 0).Td(30, 0).Td(30, 0)),

		// Highlight the curve vertex and control points
		factoryBrowser.NewTagSvgG().Append(
			factoryBrowser.NewTagSvgPolyline().Points([][]float64{{10, 50}, {25, 25}, {40, 50}}).Stroke(color.RGBA{R: 0, G: 0, B: 0, A: uint8(float64(255) * 0.2)}).Fill(nil),
			factoryBrowser.NewTagSvgCircle().Cx(25).Cy(25).R(1.5),

			// Curve vertex points
			factoryBrowser.NewTagSvgCircle().Cx(10).Cy(50).R(1.5),
			factoryBrowser.NewTagSvgCircle().Cx(40).Cy(50).R(1.5),

			factoryBrowser.NewTagSvgG().Id("SmoothQuadraticDown").Append(
				factoryBrowser.NewTagSvgPolyline().Points([][]float64{{40, 50}, {55, 75}, {70, 50}}).Stroke(color.RGBA{R: 0, G: 0, B: 0, A: uint8(float64(255) * 0.2)}).StrokeDasharray(2).Fill(nil),
				factoryBrowser.NewTagSvgCircle().Cx(55).Cy(75).R(1.5).Fill(factoryColor.NewLightgray()),
				factoryBrowser.NewTagSvgCircle().Cx(70).Cy(50).R(1.5),
			),

			factoryBrowser.NewTagSvgG().Id("SmoothQuadraticUp").Append(
				factoryBrowser.NewTagSvgPolyline().Points([][]float64{{70, 50}, {85, 25}, {100, 50}}).Stroke(color.RGBA{R: 0, G: 0, B: 0, A: uint8(float64(255) * 0.2)}).StrokeDasharray(2).Fill(nil),
				factoryBrowser.NewTagSvgCircle().Cx(85).Cy(25).R(1.5).Fill(factoryColor.NewLightgray()),
				factoryBrowser.NewTagSvgCircle().Cx(100).Cy(50).R(1.5),
			),

			factoryBrowser.NewTagSvgUse().HRef("#SmoothQuadraticDown").X(60),
			factoryBrowser.NewTagSvgUse().HRef("#SmoothQuadraticUp").X(60),
			factoryBrowser.NewTagSvgUse().HRef("#SmoothQuadraticDown").X(120),
		),
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
