// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/class
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/class
//
//  <svg width="120" height="220"
//      viewPort="0 0 120 120" version="1.1"
//      xmlns="http://www.w3.org/2000/svg">
//
//      <style type="text/css" >
//          <![CDATA[
//              rect.rectClass {
//                  stroke: #000066;
//                  fill:   #00cc00;
//              }
//              circle.circleClass {
//                  stroke: #006600;
//                  fill:   #cc0000;
//              }
//          ]]>
//      </style>
//
//      <rect class="rectClass" x="10" y="10" width="100" height="100"/>
//      <circle class="circleClass" cx="40" cy="50" r="26"/>
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

	s1 := factoryBrowser.NewTagSvg().Width(120).Height(220).ViewBox([]float64{0, 0, 120, 120}).Append(
		factoryBrowser.NewTagSvgStyle().Style(
			//"<![CDATA[\n"+
			"rect.rectClass {\n"+
				"stroke: #000066;\n"+
				"fill:   #00cc00;\n"+
				"}\n"+
				"circle.circleClass {\n"+
				"stroke: #006600;\n"+
				"fill:   #cc0000;\n"+
				"}\n",
			//"]]>",
		),

		factoryBrowser.NewTagSvgRect().Class("rectClass").X(10).Y(10).Width(100).Height(100),
		factoryBrowser.NewTagSvgCircle().Class("circleClass").Cx(40).Cy(50).R(26),
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
