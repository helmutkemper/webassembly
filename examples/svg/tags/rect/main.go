// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/rect
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/rect
//
//  <svg viewBox="0 0 220 100" xmlns="http://www.w3.org/2000/svg">
//    <!-- Simple rectangle -->
//    <rect width="100" height="100" />
//
//    <!-- Rounded corner rectangle -->
//    <rect x="120" width="100" height="100" rx="15" />
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 220, 100}).Append(
		// Simple rectangle
		factoryBrowser.NewTagSvgRect().Width(100).Height(100),

		// Rounded corner rectangle
		factoryBrowser.NewTagSvgRect().X(120).Width(100).Height(100).Rx(15),
	)

	stage.Append(s1)

	<-done
}
