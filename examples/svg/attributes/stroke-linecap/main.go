// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stroke-linecap
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stroke-linecap
//
//  <svg viewBox="0 0 6 6" xmlns="http://www.w3.org/2000/svg">
//
//    <!-- Effect of the (default) "butt" value -->
//    <line x1="1" y1="1" x2="5" y2="1" stroke="black"
//          stroke-linecap="butt" />
//
//    <!-- Effect of the "round" value -->
//    <line x1="1" y1="3" x2="5" y2="3" stroke="black"
//          stroke-linecap="round" />
//
//    <!-- Effect of the "square" value -->
//    <line x1="1" y1="5" x2="5" y2="5" stroke="black"
//          stroke-linecap="square" />
//
//    <!--
//    the following pink lines highlight the
//    position of the path for each stroke
//    -->
//    <path d="M1,1 h4 M1,3 h4 M1,5 h4" stroke="pink" stroke-width="0.025" />
//  </svg>

//go:build js

// bug: o original não apararece nada

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 6, 6}).Append(
		// Effect of the (default) "butt" value
		factoryBrowser.NewTagSvgLine().X1(1).Y1(1).X2(5).Y2(1).Stroke(factoryColor.NewBlack()).StrokeLineCap(html.KSvgStrokeLinecapButt),

		// Effect of the "round" value
		factoryBrowser.NewTagSvgLine().X1(1).Y1(3).X2(5).Y2(3).Stroke(factoryColor.NewBlack()).StrokeLineCap(html.KSvgStrokeLinecapRound),

		// Effect of the "square" value
		factoryBrowser.NewTagSvgLine().X1(1).Y1(5).X2(5).Y2(5).Stroke(factoryColor.NewBlack()).StrokeLineCap(html.KSvgStrokeLinecapSquare),

		// the following pink lines highlight the
		// position of the path for each stroke
		factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(1, 1).Hd(4).M(1, 3).Hd(4).M(1, 5).Hd(4)).Stroke(factoryColor.NewPink()).StrokeWidth(0.025),
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
