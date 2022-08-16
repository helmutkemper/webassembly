// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/x1
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/x1
//
//  <svg viewBox="0 0 10 10" xmlns="http://www.w3.org/2000/svg">
//    <line x1="1" x2="5" y1="1" y2="9" stroke="red" />
//    <line x1="5" x2="5" y1="1" y2="9" stroke="green" />
//    <line x1="9" x2="5" y1="1" y2="9" stroke="blue" />
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 10, 10}).Append(
		factoryBrowser.NewTagSvgLine().X1(1).X2(5).Y1(1).Y2(9).Stroke(factoryColor.NewRed()),
		factoryBrowser.NewTagSvgLine().X1(5).X2(5).Y1(1).Y2(9).Stroke(factoryColor.NewGreen()),
		factoryBrowser.NewTagSvgLine().X1(9).X2(5).Y1(1).Y2(9).Stroke(factoryColor.NewBlue()),
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
