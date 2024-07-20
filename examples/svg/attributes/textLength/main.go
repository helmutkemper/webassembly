// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/textLength
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/textLength
//
//  <svg viewBox="0 0 200 60" xmlns="http://www.w3.org/2000/svg">
//    <text y="20" textLength="6em">Small text length</text>
//    <text y="40" textLength="120%">Big text length</text>
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
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 200, 60}).Append(
		factoryBrowser.NewTagSvgText().Y(20).TextLength("6em").Text("Small text length"),
		factoryBrowser.NewTagSvgText().Y(40).TextLength(float32(1.2)).Text("Big text length"),
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
