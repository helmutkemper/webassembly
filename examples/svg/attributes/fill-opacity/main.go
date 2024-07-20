// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/fill-opacity
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/fill-opacity
//
//  <svg viewBox="0 0 400 100" xmlns="http://www.w3.org/2000/svg">
//    <!-- Default fill opacity: 1 -->
//    <circle cx="50" cy="50" r="40" />
//
//    <!-- Fill opacity as a number -->
//    <circle cx="150" cy="50" r="40"
//            fill-opacity="0.7" />
//
//    <!-- Fill opacity as a percentage -->
//    <circle cx="250" cy="50" r="40"
//            fill-opacity="50%" />
//
//    <!-- Fill opacity as a CSS property -->
//    <circle cx="350" cy="50" r="40"
//            style="fill-opacity: .25;" />
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 400, 100}).Append(
		// Default fill opacity: 1
		factoryBrowser.NewTagSvgCircle().Cx(50).Cy(50).R(40),

		// Fill opacity as a number
		factoryBrowser.NewTagSvgCircle().Cx(150).Cy(50).R(40).FillOpacity(0.7),

		// Fill opacity as a percentage
		factoryBrowser.NewTagSvgCircle().Cx(250).Cy(50).R(40).FillOpacity(float32(0.5)),

		// Fill opacity as a CSS property
		factoryBrowser.NewTagSvgCircle().Cx(350).Cy(50).R(40).Style("fill-opacity: .25;"),
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
