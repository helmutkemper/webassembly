// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/lengthAdjust
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/lengthAdjust
//
//  <svg viewBox="0 0 400 30" xmlns="http://www.w3.org/2000/svg">
//    <text y="20" letter-spacing="2">Bigger letter-spacing</text>
//    <text x="200" y="20" letter-spacing="-0.5">Smaller letter-spacing</text>
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 400, 30}).Append(
		factoryBrowser.NewTagSvgText().Y(20).LetterSpacing(2).Text("Bigger letter-spacing"),
		factoryBrowser.NewTagSvgText().X(200).Y(20).LetterSpacing(-0.5).Text("Smaller letter-spacing"),
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
