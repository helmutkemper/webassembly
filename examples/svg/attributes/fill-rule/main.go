// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/fill-rule
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/fill-rule
//
//  <svg viewBox="-10 -10 220 120" xmlns="http://www.w3.org/2000/svg">
//    <!-- Default value for fill-rule -->
//    <polygon fill-rule="nonzero" stroke="red"
//     points="50,0 21,90 98,35 2,35 79,90"/>
//
//    <!--
//    The center of the shape has two
//    path segments (shown by the red stroke)
//    between it and infinity. It is therefore
//    considered outside the shape, and not filled.
//    -->
//    <polygon fill-rule="evenodd" stroke="red"
//     points="150,0 121,90 198,35 102,35 179,90"/>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{-10, -10, 220, 120}).Append(
		// Default value for fill-rule
		factoryBrowser.NewTagSvgPolygon().FillRule(html.KSvgFillRuleNonzero).Stroke(factoryColor.NewRed()).Points([][]float64{{50, 0}, {21, 90}, {98, 35}, {2, 35}, {79, 90}}),

		// The center of the shape has two
		// path segments (shown by the red stroke)
		// between it and infinity. It is therefore
		// considered outside the shape, and not filled.

		factoryBrowser.NewTagSvgPolygon().FillRule(html.KSvgFillRuleEvenOdd).Stroke(factoryColor.NewRed()).Points([][]float64{{150, 0}, {121, 90}, {198, 35}, {102, 35}, {179, 90}}),
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
