// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/transform
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/transform
//
//    <svg viewBox="-50 -50 100 100" xmlns="http://www.w3.org/2000/svg">
//    <!-- uniform scale -->
//    <circle cx="0" cy="0" r="10" fill="red"
//            transform="scale(4)" />
//
//    <!-- vertical scale -->
//    <circle cx="0" cy="0" r="10" fill="yellow"
//            transform="scale(1,4)" />
//
//    <!-- horizontal scale -->
//    <circle cx="0" cy="0" r="10" fill="pink"
//            transform="scale(4,1)" />
//
//    <!-- No scale -->
//    <circle cx="0" cy="0" r="10" fill="black" />
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{-50, -50, 100, 100}).Append(
		// uniform scale
		factoryBrowser.NewTagSvgCircle().Cx(0).Cy(0).R(10).Fill(factoryColor.NewRed()).Transform(factoryBrowser.NewTransform().Scale(4, 4)),

		// vertical scale
		factoryBrowser.NewTagSvgCircle().Cx(0).Cy(0).R(10).Fill(factoryColor.NewYellow()).Transform(factoryBrowser.NewTransform().Scale(1, 4)),

		// horizontal scale
		factoryBrowser.NewTagSvgCircle().Cx(0).Cy(0).R(10).Fill(factoryColor.NewPink()).Transform(factoryBrowser.NewTransform().Scale(4, 1)),

		// No scale
		factoryBrowser.NewTagSvgCircle().Cx(0).Cy(0).R(10).Fill(factoryColor.NewBlack()),
	)

	stage.Append(s1)

	<-done
}
