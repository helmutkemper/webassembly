// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/style
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/style
//
//  <svg viewBox="0 0 10 10" xmlns="http://www.w3.org/2000/svg">
//    <style>
//      circle {
//        fill: gold;
//        stroke: maroon;
//        stroke-width: 2px;
//      }
//    </style>
//
//    <circle cx="5" cy="5" r="4" />
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 10, 10}).XmlnsXLink("http://www.w3.org/1999/xlink").Append(
		factoryBrowser.NewTagSvgStyle().Style(
			"circle {\n"+
				"fill: gold;\n"+
				"stroke: maroon;\n"+
				"stroke-width: 2px;\n"+
				"}",
		),
		factoryBrowser.NewTagSvgCircle().Cx(5).Cy(5).R(4),
	)

	stage.Append(s1)

	<-done
}
