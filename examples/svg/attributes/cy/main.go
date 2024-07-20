// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/cy
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/cy
//
//  <svg viewBox="0 0 100 300" xmlns="http://www.w3.org/2000/svg">
//    <radialGradient cy="25%" id="myGradient">
//      <stop offset="0"    stop-color="white" />
//      <stop offset="100%" stop-color="black" />
//    </radialGradient>
//
//    <circle cy="50"  cx="50" r="45"/>
//    <ellipse cy="150" cx="50" rx="45" ry="25" />
//    <rect x="5" y="205" width="90" height="90" fill="url(#myGradient)" />
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 300}).Append(
		factoryBrowser.NewTagSvgRadialGradient().Cx(float32(0.25)).Id("myGradient").Append(
			factoryBrowser.NewTagSvgStop().Offset(0).StopColor(factoryColor.NewWhite()),
			factoryBrowser.NewTagSvgStop().Offset(float32(1.0)).StopColor(factoryColor.NewBlack()),
		),

		factoryBrowser.NewTagSvgCircle().Cy(50).Cx(50).R(45),
		factoryBrowser.NewTagSvgEllipse().Cy(150).Cx(50).Rx(45).Ry(25),
		factoryBrowser.NewTagSvgRect().X(5).Y(205).Width(90).Height(90).Fill("url(#myGradient)"),
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
