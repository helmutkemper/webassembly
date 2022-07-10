// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/style
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/style
//
//  <svg viewbox="0 0 100 60" xmlns="http://www.w3.org/2000/svg">
//    <rect width="80"  height="40" x="10" y="10"
//        style="fill: skyblue; stroke: cadetblue; stroke-width: 2;"/>
//  </svg>

//go:build js
// +build js

// bug: o original não apararece nada

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 60}).Append(
		factoryBrowser.NewTagSvgRect().Width(80).Height(40).X(10).Y(10).Style("fill: skyblue; stroke: cadetblue; stroke-width: 2;"),
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
