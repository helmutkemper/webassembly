// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/transform
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/transform
//
//  <svg viewBox="-5 -5 10 10" xmlns="http://www.w3.org/2000/svg">
//    <rect x="-3" y="-3" width="6" height="6" />
//
//    <rect x="-3" y="-3" width="6" height="6" fill="red"
//          transform="skewY(30)" />
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

	s1 := factoryBrowser.NewTagSvg().
		ViewBox([]float64{-5, -5, 10, 10}).
		Append(

			factoryBrowser.NewTagSvgRect().
				X(-3).
				Y(-3).
				Width(6).
				Height(6),

			factoryBrowser.NewTagSvgRect().
				X(-3).
				Y(-3).
				Width(6).
				Height(6).
				Fill(factoryColor.NewRed()).
				Transform(

					factoryBrowser.NewTransform().
						SkewY(30),
				),
		)

	stage.Append(s1)

	<-done
}
