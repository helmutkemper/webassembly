// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/d
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/d
//
//  <svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
//    <path fill="none" stroke="red"
//      d="M 10,30
//         A 20,20 0,0,1 50,30
//         A 20,20 0,0,1 90,30
//         Q 90,60 50,90
//         Q 10,60 10,30 z" />
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Append(

		factoryBrowser.NewTagSvgPath().Fill(nil).Stroke(factoryColor.NewRed()).D(factoryBrowser.NewPath().M(10, 30).A(20, 20, 0, 0, 1, 50, 30).A(20, 20, 0, 0, 1, 90, 30).Q(90, 60, 50, 90).Q(10, 60, 10, 30).Z()),
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
