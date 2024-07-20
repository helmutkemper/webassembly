// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/width
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/width
//
//  <svg viewBox="0 0 100 300" xmlns="http://www.w3.org/2000/svg">
//    <!-- With a width of 0 or less, nothing will be rendered -->
//    <rect x="0" y="0" width="0" height="90"/>
//    <rect x="0" y="100" width="60" height="90"/>
//    <rect x="0" y="200" width="100%" height="90"/>
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 300}).Append(
		// With a width of 0 or less, nothing will be rendered
		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(0).Height(90),
		factoryBrowser.NewTagSvgRect().X(0).Y(100).Width(60).Height(90),
		factoryBrowser.NewTagSvgRect().X(0).Y(200).Width(float32(1)).Height(90),
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
