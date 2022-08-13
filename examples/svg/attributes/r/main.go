// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/r
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/r
//
//  <svg viewBox="0 0 300 200" xmlns="http://www.w3.org/2000/svg">
//    <radialGradient r="0" id="myGradient000">
//      <stop offset="0"    stop-color="white" />
//      <stop offset="100%" stop-color="black" />
//    </radialGradient>
//    <radialGradient r="50%" id="myGradient050">
//      <stop offset="0"    stop-color="white" />
//      <stop offset="100%" stop-color="black" />
//    </radialGradient>
//    <radialGradient r="100%" id="myGradient100">
//      <stop offset="0"    stop-color="white" />
//      <stop offset="100%" stop-color="black" />
//    </radialGradient>
//
//    <circle cx="50"  cy="50" r="0"/>
//    <circle cx="150" cy="50" r="25"/>
//    <circle cx="250" cy="50" r="50"/>
//
//    <rect x="20"  y="120" width="60" height="60" fill="url(#myGradient000)" />
//    <rect x="120" y="120" width="60" height="60" fill="url(#myGradient050)" />
//    <rect x="220" y="120" width="60" height="60" fill="url(#myGradient100)" />
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
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 300, 200}).Append(
		factoryBrowser.NewTagSvgRadialGradient().R(float32(0)).Id("myGradient000").Append(
			factoryBrowser.NewTagSvgStop().Offset(float32(0)).StopColor(factoryColor.NewWhite()),
			factoryBrowser.NewTagSvgStop().Offset(float32(1)).StopColor(factoryColor.NewBlack()),
		),
		factoryBrowser.NewTagSvgRadialGradient().R(float32(0.5)).Id("myGradient050").Append(
			factoryBrowser.NewTagSvgStop().Offset(float32(0)).StopColor(factoryColor.NewWhite()),
			factoryBrowser.NewTagSvgStop().Offset(float32(1)).StopColor(factoryColor.NewBlack()),
		),
		factoryBrowser.NewTagSvgRadialGradient().R(float32(1.0)).Id("myGradient100").Append(
			factoryBrowser.NewTagSvgStop().Offset(float32(0)).StopColor(factoryColor.NewWhite()),
			factoryBrowser.NewTagSvgStop().Offset(float32(1)).StopColor(factoryColor.NewBlack()),
		),

		factoryBrowser.NewTagSvgCircle().Cx(50).Cy(50).R(0),
		factoryBrowser.NewTagSvgCircle().Cx(150).Cy(50).R(25),
		factoryBrowser.NewTagSvgCircle().Cx(250).Cy(50).R(50),

		factoryBrowser.NewTagSvgRect().X(20).Y(120).Width(60).Height(60).Fill("url(#myGradient000)"),
		factoryBrowser.NewTagSvgRect().X(120).Y(120).Width(60).Height(60).Fill("url(#myGradient050)"),
		factoryBrowser.NewTagSvgRect().X(220).Y(120).Width(60).Height(60).Fill("url(#myGradient100)"),
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
