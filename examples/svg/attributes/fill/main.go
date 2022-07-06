// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/fill
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/fill
//
//  <svg viewBox="0 0 300 100" xmlns="http://www.w3.org/2000/svg">
//    <!-- Simple color fill -->
//    <circle cx="50" cy="50" r="40" fill="pink" />
//
//    <!-- Fill circle with a gradient -->
//    <defs>
//      <radialGradient id="myGradient">
//        <stop offset="0%"   stop-color="pink" />
//        <stop offset="100%" stop-color="black" />
//      </radialGradient>
//    </defs>
//
//    <circle cx="150" cy="50" r="40" fill="url(#myGradient)" />
//
//    <!--
//    Keeping the final state of an animated circle
//    which is a circle with a radius of 40.
//    -->
//    <circle cx="250" cy="50" r="20">
//      <animate attributeType="XML"
//               attributeName="r"
//               from="0" to="40" dur="5s"
//               fill="freeze" />
//    </circle>
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
	"time"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 300, 100}).Append(
		// Simple color fill
		factoryBrowser.NewTagSvgCircle().Cx(50).Cy(50).R(40).Fill(factoryColor.NewPink()),

		// Fill circle with a gradient
		factoryBrowser.NewTagSvgDefs().Append(
			factoryBrowser.NewTagSvgRadialGradient().Id("myGradient").Append(
				factoryBrowser.NewTagSvgStop().Offset(float32(0.0)).StopColor(factoryColor.NewPink()),
				factoryBrowser.NewTagSvgStop().Offset(float32(1.0)).StopColor(factoryColor.NewBlack()),
			),
		),

		factoryBrowser.NewTagSvgCircle().Cx(150).Cy(50).R(40).Fill("url(#myGradient)"),

		// Keeping the final state of an animated circle
		// which is a circle with a radius of 40.

		factoryBrowser.NewTagSvgCircle().Cx(250).Cy(50).R(20).Append(
			factoryBrowser.NewTagSvgAnimate().AttributeName("r").From(0).To(40).Dur(5*time.Second).Fill("freeze"),
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
