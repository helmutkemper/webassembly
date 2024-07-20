// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/pathLength
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/pathLength
//
//  <svg viewBox="0 0 100 60" xmlns="http://www.w3.org/2000/svg">
//    <style>
//    path {
//      fill: none;
//      stroke: black;
//      stroke-width: 2;
//      stroke-dasharray: 10;
//    }
//    </style>
//
//    <!-- No pathLength, the real length of the path is used. In that case, 100 user units -->
//    <path d="M 0,10 h100"/>
//
//    <!-- compute everything like if the path length was 90 user units long -->
//    <path d="M 0,20 h100" pathLength="90"/>
//
//    <!-- compute everything like if the path length was 50 user units long -->
//    <path d="M 0,30 h100" pathLength="50"/>
//
//    <!-- compute everything like if the path length was 30 user units long -->
//    <path d="M 0,40 h100" pathLength="30"/>
//
//    <!-- compute everything like if the path length was 10 user units long -->
//    <path d="M 0,50 h100" pathLength="10"/>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 60}).Append(
		factoryBrowser.NewTagSvgStyle().Style(
			"path {\n"+
				"fill: none;\n"+
				"stroke: black;\n"+
				"stroke-width: 2;\n"+
				"stroke-dasharray: 10;\n"+
				"}",
		),

		// No pathLength, the real length of the path is used. In that case, 100 user units
		factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(0, 10).Hd(100)),

		// compute everything like if the path length was 90 user units long
		factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(0, 20).Hd(100)).PathLength(90),

		// compute everything like if the path length was 50 user units long
		factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(0, 30).Hd(100)).PathLength(50),

		// compute everything like if the path length was 30 user units long
		factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(0, 40).Hd(100)).PathLength(30),

		// compute everything like if the path length was 10 user units long
		factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(0, 50).Hd(100)).PathLength(10),
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
