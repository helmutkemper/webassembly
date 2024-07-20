// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stroke-dashoffset
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stroke-dashoffset
//
//  <svg viewBox="-3 0 33 10" xmlns="http://www.w3.org/2000/svg">
//    <!-- No dash array -->
//    <line x1="0" y1="1" x2="30" y2="1" stroke="black" />
//
//    <!-- No dash offset -->
//    <line x1="0" y1="3" x2="30" y2="3" stroke="black"
//          stroke-dasharray="3 1" />
//
//    <!--
//    The start of the dash array computation
//    is pulled by 3 user units
//    -->
//    <line x1="0" y1="5" x2="30" y2="5" stroke="black"
//          stroke-dasharray="3 1"
//          stroke-dashoffset="3" />
//
//    <!--
//    The start of the dash array computation
//    is pushed by 3 user units
//    -->
//    <line x1="0" y1="7" x2="30" y2="7" stroke="black"
//          stroke-dasharray="3 1"
//          stroke-dashoffset="-3" />
//
//    <!--
//    The start of the dash array computation
//    is pulled by 1 user units which ends up
//    in the same rendering as the previous example
//    -->
//    <line x1="0" y1="9" x2="30" y2="9" stroke="black"
//          stroke-dasharray="3 1"
//          stroke-dashoffset="1" />
//
//    <!--
//    the following red lines highlight the
//    offset of the dash array for each line
//    -->
//    <path d="M0,5 h-3 M0,7 h3 M0,9 h-1" stroke="rgba(255,0,0,.5)" />
//  </svg>

//go:build js

// bug: o original não apararece nada

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
	"image/color"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{-3, 0, 33, 10}).Append(
		// No dash array
		factoryBrowser.NewTagSvgLine().X1(0).Y1(1).X2(30).Y2(1).Stroke(factoryColor.NewBlack()),

		// No dash offset
		factoryBrowser.NewTagSvgLine().X1(0).Y1(3).X2(30).Y2(3).Stroke(factoryColor.NewBlack()).StrokeDasharray([]float64{3, 1}),

		// The start of the dash array computation
		// is pulled by 3 user units
		factoryBrowser.NewTagSvgLine().X1(0).Y1(5).X2(30).Y2(5).Stroke(factoryColor.NewBlack()).StrokeDasharray([]float64{3, 1}).StrokeDashOffset(3),

		// The start of the dash array computation
		// is pushed by 3 user units
		factoryBrowser.NewTagSvgLine().X1(0).Y1(7).X2(30).Y2(7).Stroke(factoryColor.NewBlack()).StrokeDasharray([]float64{3, 1}).StrokeDashOffset(-3),

		// The start of the dash array computation
		// is pulled by 1 user units which ends up
		// in the same rendering as the previous example
		factoryBrowser.NewTagSvgLine().X1(0).Y1(9).X2(30).Y2(9).Stroke(factoryColor.NewBlack()).StrokeDasharray([]float64{3, 1}).StrokeDashOffset(1),

		// the following red lines highlight the
		// offset of the dash array for each line
		factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(0, 5).Hd(-3).M(0, 7).Hd(3).M(0, 9).Hd(-1)).Stroke(color.RGBA{R: 255, G: 0, B: 0, A: 128}),
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
