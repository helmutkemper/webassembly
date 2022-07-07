// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/title
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/title
//
//  <svg viewBox="0 0 20 10" xmlns="http://www.w3.org/2000/svg">
//    <circle cx="5" cy="5" r="4">
//      <title>I'm a circle</title>
//    </circle>
//
//    <rect x="11" y="1" width="8" height="8">
//      <title>I'm a square</title>
//    </rect>
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 20, 10}).Append(
		factoryBrowser.NewTagSvgCircle().Cx(5).Cy(5).R(4).Append(
			factoryBrowser.NewTagSvgTitle().Title("I'm a circle"),
		),

		factoryBrowser.NewTagSvgRect().X(11).Y(1).Width(8).Height(8).Append(
			factoryBrowser.NewTagSvgTitle().Title("I'm a square"),
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
//
//
//
//
//
//
//
//
