// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/color
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/color
//
//  <svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
//    <g color="green">
//      <rect width="50" height="50" fill="currentcolor" />
//      <circle r="25" cx="70" cy="70" stroke="currentcolor" fill="none" stroke-width="5" />
//    </g>
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Append(
		factoryBrowser.NewTagSvgG().Color(factoryColor.NewGreen()).Append(
			factoryBrowser.NewTagSvgRect().Width(50).Height(50).Fill("currentcolor"),
			factoryBrowser.NewTagSvgCircle().R(25).Cx(70).Cy(70).Stroke("currentcolor").Fill(nil).StrokeWidth(5),
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
