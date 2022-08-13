// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/g
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/g
//
//  <svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
//    <!-- Using g to inherit presentation attributes -->
//    <g fill="white" stroke="green" stroke-width="5">
//      <circle cx="40" cy="40" r="25" />
//      <circle cx="60" cy="60" r="25" />
//    </g>
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Append(
		// Using g to inherit presentation attributes
		factoryBrowser.NewTagSvgG().Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewGreen()).StrokeWidth(5).Append(
			factoryBrowser.NewTagSvgCircle().Cx(40).Cy(40).R(25),
			factoryBrowser.NewTagSvgCircle().Cx(60).Cy(60).R(25),
		),
	)

	stage.Append(s1)

	<-done
}
