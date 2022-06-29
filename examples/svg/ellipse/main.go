// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/transform
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/transform
//
//  <svg viewBox="0 0 200 100" xmlns="http://www.w3.org/2000/svg">
//    <ellipse cx="100" cy="50" rx="100" ry="50" />
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
		ViewBox([]float64{0, 0, 200, 100}).
		Append(

			factoryBrowser.NewTagSvgEllipse().
				Cx(100).
				Cy(50).
				Rx(100).
				Ry(50),
		)

	stage.Append(s1)

	<-done
}
