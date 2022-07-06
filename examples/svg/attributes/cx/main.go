// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/cx
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/cx
//
//  <svg viewBox="0 0 300 100" xmlns="http://www.w3.org/2000/svg">
//    <radialGradient cx="25%" id="myGradient">
//      <stop offset="0"    stop-color="white" />
//      <stop offset="100%" stop-color="black" />
//    </radialGradient>
//
//    <circle cx="50" cy="50" r="45"/>
//    <ellipse cx="150" cy="50" rx="45" ry="25" />
//    <rect x="205" y="5" width="90" height="90" fill="url(#myGradient)" />
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 300, 100}).Append(
		factoryBrowser.NewTagSvgRadialGradient().Cx(float32(0.25)).Id("myGradient").Append(
			factoryBrowser.NewTagSvgStop().Offset(0).StopColor(factoryColor.NewWhite()),
			factoryBrowser.NewTagSvgStop().Offset(float32(1.0)).StopColor(factoryColor.NewBlack()),
		),

		factoryBrowser.NewTagSvgCircle().Cx(50).Cy(50).R(45),
		factoryBrowser.NewTagSvgEllipse().Cx(150).Cy(50).Rx(45).Ry(25),
		factoryBrowser.NewTagSvgRect().X(205).Y(5).Width(90).Height(90).Fill("url(#myGradient)"),
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
