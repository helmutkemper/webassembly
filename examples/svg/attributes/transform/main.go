// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/transform
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/transform
//
//  <svg viewBox="-40 0 150 100" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
//    <g fill="grey"
//       transform="rotate(-10 50 100)
//                  translate(-36 45.5)
//                  skewX(40)
//                  scale(1 0.5)">
//      <path id="heart" d="M 10,30 A 20,20 0,0,1 50,30 A 20,20 0,0,1 90,30 Q 90,60 50,90 Q 10,60 10,30 z" />
//    </g>
//
//    <use xlink:href="#heart" fill="none" stroke="red"/>
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{-40, 0, 150, 100}).XmlnsXLink("http://www.w3.org/1999/xlink").Append(
		factoryBrowser.NewTagSvgG().Fill(factoryColor.NewGray()).Transform(
			factoryBrowser.NewTransform().Rotate(-10, 50, 100).Translate(-36, 45.5).SkewX(40).Scale(1, 0.5)).Append(
			factoryBrowser.NewTagSvgPath().Id("heart").D(factoryBrowser.NewPath().M(10, 30).A(20, 20, 0, 0, 1, 50, 30).A(20, 20, 0, 0, 1, 90, 30).Q(90, 60, 50, 90).Q(10, 60, 10, 30).Z()),
		),

		factoryBrowser.NewTagSvgUse().HRef("#heart").Fill("none").Stroke(factoryColor.NewRed()),
	)

	stage.Append(s1)

	<-done
}
