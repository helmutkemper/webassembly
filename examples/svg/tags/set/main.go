// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/set
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/set
//
//  <svg viewBox="0 0 10 10" xmlns="http://www.w3.org/2000/svg">
//    <style>
//      rect { cursor: pointer }
//      .round { rx: 5px; fill: green; }
//    </style>
//
//    <rect id="me" width="10" height="10">
//      <set attributeName="class" to="round" begin="me.click" dur="2s" />
//    </rect>
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"time"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().
		ViewBox([]float64{0, 0, 10, 10}).
		Append(

			factoryBrowser.NewTagSvgStyle().Style(
				"rect { cursor: pointer }\n"+
					".round { rx: 5px; fill: green; }",
			),

			factoryBrowser.NewTagSvgRect().Id("me").Width(10).Height(10).Append(

				factoryBrowser.NewTagSvgSet().AttributeName("class").To("round").Begin("me.click").Dur(2*time.Second),
			),
		)

	stage.Append(s1)

	<-done
}
