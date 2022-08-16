// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/media
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/media
//
//  <svg viewBox="0 0 240 220" xmlns="http://www.w3.org/2000/svg">
//    <style>
//      rect { fill: black; }
//    </style>
//    <style media="all and (min-width: 600px)">
//      rect { fill: seagreen; }
//    </style>
//
//    <text y="15">Resize the window to see the effect</text>
//    <rect y="20" width="200" height="200" />
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 240, 220}).Append(
		factoryBrowser.NewTagSvgStyle().Style("rect { fill: black; }"),
		factoryBrowser.NewTagSvgStyle().Media("all and (min-width: 600px)").Style("rect { fill: seagreen; }"),
		factoryBrowser.NewTagSvgText().Y(15).Text("Resize the window to see the effect"),
		factoryBrowser.NewTagSvgRect().Y(20).Width(200).Height(200),
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
