// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/tspan
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/tspan
//
//  <svg viewBox="0 0 240 40" xmlns="http://www.w3.org/2000/svg">
//    <style>
//      text  { font: italic 12px serif; }
//      tspan { font: bold 10px sans-serif; fill: red; }
//    </style>
//
//    <text x="10" y="30" class="small">
//      You are
//      <tspan>not</tspan>
//      a banana!
//    </text>
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
		ViewBox([]float64{0, 0, 240, 40}).
		Append(

			factoryBrowser.NewTagSvgStyle().Style(
				"text  { font: italic 12px serif; }\n"+
					"tspan { font: bold 10px sans-serif; fill: red; }",
			),

			factoryBrowser.NewTagSvgText().X(10).Y(30).Class("small").Html("You are\n<tspan>not</tspan>\na banana!"),
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
//
//
//
//
//
//
//
//
