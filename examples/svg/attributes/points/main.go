// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/points
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/points
//
//  <svg viewBox="-10 -10 220 120" xmlns="http://www.w3.org/2000/svg">
//    <!-- polyline is an open shape -->
//    <polyline stroke="black" fill="none"
//     points="50,0 21,90 98,35 2,35 79,90"/>
//
//    <!-- polygon is a closed shape -->
//    <polygon stroke="black" fill="none" transform="translate(100,0)"
//     points="50,0 21,90 98,35 2,35 79,90"/>
//
//    <!--
//    It is usually considered best practices to separate a X and Y
//    coordinate with a comma and a group of coordinates by a space.
//    It makes things more readable for human beings.
//    -->
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
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{-10, -10, 220, 120}).Append(
		// polyline is an open shape
		factoryBrowser.NewTagSvgPolyline().Stroke(factoryColor.NewBlack()).Fill(nil).Points([][]float64{{50, 0}, {21, 90}, {98, 35}, {2, 35}, {79, 90}}),

		// polygon is a closed shape
		factoryBrowser.NewTagSvgPolygon().Stroke(factoryColor.NewBlack()).Fill(nil).Transform(factoryBrowser.NewTransform().Translate(100, 0)).Points([][]float64{{50, 0}, {21, 90}, {98, 35}, {2, 35}, {79, 90}}),

		// It is usually considered best practices to separate a X and Y
		// coordinate with a comma and a group of coordinates by a space.
		// It makes things more readable for human beings.
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
