// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/x1
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/x1
//
//  <svg viewBox="0 0 20 10" xmlns="http://www.w3.org/2000/svg">
//    <!--
//    By default the gradient vector start at the left
//    bounding limit of the shape it is applied to
//    -->
//    <linearGradient x1="0%" id="g0">
//      <stop offset="0"    stop-color="black" />
//      <stop offset="100%" stop-color="red" />
//    </linearGradient>
//
//    <rect x="1"  y="1" width="8" height="8" fill="url(#g0)" />
//
//    <!--
//    Here the gradient vector start at 80% of the left
//    bounding limit of the shape it is applied to
//    -->
//    <linearGradient x1="80%" id="g1">
//      <stop offset="0"    stop-color="black" />
//      <stop offset="100%" stop-color="red" />
//    </linearGradient>
//
//    <rect x="11" y="1" width="8" height="8" fill="url(#g1)" />
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 20, 10}).Append(
		// By default the gradient vector start at the left
		// bounding limit of the shape it is applied to
		factoryBrowser.NewTagSvgLinearGradient().X1(float32(0)).Id("g0").Append(
			factoryBrowser.NewTagSvgStop().Offset(float32(0)).StopColor(factoryColor.NewBlack()),
			factoryBrowser.NewTagSvgStop().Offset(float32(1)).StopColor(factoryColor.NewRed()),
		),

		factoryBrowser.NewTagSvgRect().X(1).Y(1).Width(8).Height(8).Fill("url(#g0)"),

		// Here the gradient vector start at 80% of the left
		// bounding limit of the shape it is applied to
		factoryBrowser.NewTagSvgLinearGradient().X1(float32(0.8)).Id("g1").Append(
			factoryBrowser.NewTagSvgStop().Offset(float32(0)).StopColor(factoryColor.NewBlack()),
			factoryBrowser.NewTagSvgStop().Offset(float32(1)).StopColor(factoryColor.NewRed()),
		),

		factoryBrowser.NewTagSvgRect().X(11).Y(1).Width(8).Height(8).Fill("url(#g1)"),
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
