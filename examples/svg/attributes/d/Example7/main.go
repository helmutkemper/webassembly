// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/d
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/d
//
//  <svg viewBox="0 -1 30 11" xmlns="http://www.w3.org/2000/svg">
//
//    <!--
//    An open shape with the last point of
//    the path different to the first one
//    -->
//    <path stroke="red"
//          d="M 5,1
//             l -4,8 8,0" />
//
//    <!--
//    An open shape with the last point of
//    the path matching the first one
//    -->
//    <path stroke="red"
//          d="M 15,1
//             l -4,8 8,0 -4,-8" />
//
//    <!--
//    A closed shape with the last point of
//    the path different to the first one
//    -->
//    <path stroke="red"
//          d="M 25,1
//             l -4,8 8,0
//             z" />
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, -1, 30, 11}).Append(

		// An open shape with the last point of
		// the path different to the first one
		factoryBrowser.NewTagSvgPath().Stroke(factoryColor.NewRed()).D(factoryBrowser.NewPath().M(5, 1).Ld(-4, 8).Ld(8, 0)),

		// An open shape with the last point of
		// the path matching the first one
		factoryBrowser.NewTagSvgPath().Stroke(factoryColor.NewRed()).D(factoryBrowser.NewPath().M(15, 1).Ld(-4, 8).Ld(8, 0).Ld(-4, -8)),

		// A closed shape with the last point of
		// the path different to the first one
		factoryBrowser.NewTagSvgPath().Stroke(factoryColor.NewRed()).D(factoryBrowser.NewPath().M(25, 1).Ld(-4, 8).Ld(8, 0).Z()),
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
