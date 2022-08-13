// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/d
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/d
//
//  <svg viewBox="0 0 200 100" xmlns="http://www.w3.org/2000/svg">
//    <!-- LineTo commands with absolute coordinates -->
//    <path fill="none" stroke="red"
//          d="M 10,10
//             L 90,90
//             V 10
//             H 50" />
//
//    <!-- LineTo commands with relative coordinates -->
//    <path fill="none" stroke="red"
//          d="M 110,10
//             l 80,80
//             v -80
//             h -40" />
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 200, 100}).Append(

		// LineTo commands with absolute coordinates
		factoryBrowser.NewTagSvgPath().Fill(nil).Stroke(factoryColor.NewRed()).D(factoryBrowser.NewPath().M(10, 10).L(90, 90).V(10).H(50)),

		// LineTo commands with relative coordinates
		factoryBrowser.NewTagSvgPath().Fill(nil).Stroke(factoryColor.NewRed()).D(factoryBrowser.NewPath().M(110, 10).Ld(80, 80).Vd(-80).Hd(-40)),
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
