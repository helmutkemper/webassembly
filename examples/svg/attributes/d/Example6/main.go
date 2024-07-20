// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/d
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/d
//
//  <svg viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
//
//    <!-- The influence of the arc flags with which the arc is drawn -->
//    <path fill="none" stroke="red"
//          d="M 6,10
//             A 6 4 10 1 0 14,10" />
//
//    <path fill="none" stroke="lime"
//          d="M 6,10
//             A 6 4 10 1 1 14,10" />
//
//    <path fill="none" stroke="purple"
//          d="M 6,10
//             A 6 4 10 0 1 14,10" />
//
//    <path fill="none" stroke="pink"
//          d="M 6,10
//             A 6 4 10 0 0 14,10" />
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 20, 20}).Append(

		// The influence of the arc flags with which the arc is drawn
		factoryBrowser.NewTagSvgPath().Fill(nil).Stroke(factoryColor.NewRed()).D(factoryBrowser.NewPath().M(6, 10).A(6, 4, 10, 1, 0, 14, 10)),
		factoryBrowser.NewTagSvgPath().Fill(nil).Stroke(factoryColor.NewLime()).D(factoryBrowser.NewPath().M(6, 10).A(6, 4, 10, 1, 1, 14, 10)),
		factoryBrowser.NewTagSvgPath().Fill(nil).Stroke(factoryColor.NewPurple()).D(factoryBrowser.NewPath().M(6, 10).A(6, 4, 10, 0, 1, 14, 10)),
		factoryBrowser.NewTagSvgPath().Fill(nil).Stroke(factoryColor.NewPink()).D(factoryBrowser.NewPath().M(6, 10).A(6, 4, 10, 0, 0, 14, 10)),
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
