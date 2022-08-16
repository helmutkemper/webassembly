// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/y
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/y
//
//  <svg viewBox="0 0 100 300" xmlns="http://www.w3.org/2000/svg">
//    <rect y="20"  x="20" width="60" height="60" />
//    <rect y="120" x="20" width="60" height="60" />
//    <rect y="220" x="20" width="60" height="60" />
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 300}).Append(
		factoryBrowser.NewTagSvgRect().Y(20).X(20).Width(60).Height(60),
		factoryBrowser.NewTagSvgRect().Y(120).X(20).Width(60).Height(60),
		factoryBrowser.NewTagSvgRect().Y(220).X(20).Width(60).Height(60),
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
