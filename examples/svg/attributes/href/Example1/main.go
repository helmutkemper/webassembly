// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/href
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/href
//
//  <svg viewBox="0 0 160 40" xmlns="http://www.w3.org/2000/svg">
//    <a href="https://developer.mozilla.org/"><text x="10" y="25">MDN Web Docs</text></a>
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

	s1 := factoryBrowser.NewTagSvg().
		ViewBox([]float64{0, 0, 160, 40}).
		Append(
			factoryBrowser.NewTagSvgA().HRef("https://developer.mozilla.org/").Append(
				factoryBrowser.NewTagSvgText().X(10).Y(10).Text("MDN Web Docs"),
			),
		)

	stage.Append(s1)

	<-done
}
