// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/height
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/height
//
//  <svg viewBox="0 0 300 100" xmlns="http://www.w3.org/2000/svg">
//    <!-- With a height of 0 or less, nothing will be rendered -->
//    <rect y="0" x="0" width="90" height="0"/>
//    <rect y="0" x="100" width="90" height="60"/>
//    <rect y="0" x="200" width="90" height="100%"/>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().
		ViewBox([]float64{0, 0, 420, 200}).
		Append(
			// With a height of 0 or less, nothing will be rendered
			factoryBrowser.NewTagSvgRect().Y(0).X(0).Width(90).Height(0),
			factoryBrowser.NewTagSvgRect().Y(0).X(100).Width(90).Height(60),
			factoryBrowser.NewTagSvgRect().Y(0).X(200).Width(90).Height(float32(1.0)),
		)

	stage.Append(s1)

	<-done
}
