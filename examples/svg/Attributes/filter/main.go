// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/filter
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/filter
//
//  <svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
//    <filter id="blur">
//      <feGaussianBlur stdDeviation="2" />
//    </filter>
//
//    <rect x="10" y="10" width="80" height="80" filter="url(#blur)" />
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Append(
		factoryBrowser.NewTagSvgFilter().Id("blur").Append(
			factoryBrowser.NewTagSvgFeGaussianBlur().StdDeviation(2),
		),

		factoryBrowser.NewTagSvgRect().X(10).Y(10).Width(80).Height(80).Filter("url(#blur)"),
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
