// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/y1
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/y1
//
//  <svg viewBox="0 0 10 10" xmlns="http://www.w3.org/2000/svg">
//    <line x1="1" x2="9" y1="1" y2="5" stroke="red" />
//    <line x1="1" x2="9" y1="5" y2="5" stroke="green" />
//    <line x1="1" x2="9" y1="9" y2="5" stroke="blue" />
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 10, 10}).Append(
		factoryBrowser.NewTagSvgLine().X1(1).X2(9).Y1(1).Y2(5).Stroke(factoryColor.NewRed()),
		factoryBrowser.NewTagSvgLine().X1(1).X2(9).Y1(5).Y2(5).Stroke(factoryColor.NewGreen()),
		factoryBrowser.NewTagSvgLine().X1(1).X2(9).Y1(9).Y2(5).Stroke(factoryColor.NewBlue()),
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
