// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/desc
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/desc
//
//  <svg viewBox="0 0 10 10" xmlns="http://www.w3.org/2000/svg">
//    <circle cx="5" cy="5" r="4">
//      <desc>
//        I'm a circle and that description is here to
//        demonstrate how I can be described, but is it
//        really necessary to describe a simple circle
//        like me?
//      </desc>
//    </circle>
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 10, 10}).Append(
		factoryBrowser.NewTagSvgCircle().Cx(5).Cy(5).R(4).Append(
			factoryBrowser.NewTagSvgDesc().Text(
				"I'm a circle and that description is here to" +
					"demonstrate how I can be described, but is it" +
					"really necessary to describe a simple circle" +
					"like me?",
			),
		),
	)

	stage.Append(s1)

	<-done
}
