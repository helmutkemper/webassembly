// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/id
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/id
//
//  <svg width="120" height="120" viewBox="0 0 120 120" xmlns="http://www.w3.org/2000/svg">
//    <style type="text/css">
//      <![CDATA[
//        #smallRect {
//          stroke: #000066;
//          fill: #00cc00;
//        }
//      ]]>
//    </style>
//
//    <rect id="smallRect" x="10" y="10" width="100" height="100" />
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

	s1 := factoryBrowser.NewTagSvg().Width(120).Height(120).ViewBox([]float64{0, 0, 120, 120}).
		Append(
			factoryBrowser.NewTagSvgStyle().Style(
				//"<![CDATA[\n" +
				"#smallRect {\n"+
					"stroke: #000066;\n"+
					"fill: #00cc00;\n"+
					"}\n",
				//"]]>",
			),

			factoryBrowser.NewTagSvgRect().Id("smallRect").X(10).Y(10).Width(100).Height(100),
		)

	stage.Append(s1)

	<-done
}
