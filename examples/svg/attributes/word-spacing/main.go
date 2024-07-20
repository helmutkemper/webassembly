// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/word-spacing
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/word-spacing
//
//  <svg viewBox="0 0 250 50" xmlns="http://www.w3.org/2000/svg">
//    <text y="20" word-spacing="2">Bigger spacing between words</text>
//    <text x="0" y="40" word-spacing="-0.5">Smaller spacing between words</text>
//  </svg>

//go:build js

// bug: o original não apararece nada

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 250, 50}).Append(
		factoryBrowser.NewTagSvgText().Y(20).WordSpacing(2).Text("Bigger spacing between words"),
		factoryBrowser.NewTagSvgText().X(0).Y(40).WordSpacing(-0.5).Text("Smaller spacing between words"),
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
