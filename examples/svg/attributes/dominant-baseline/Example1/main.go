// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/dominant-baseline
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/dominant-baseline
//
//  <svg viewBox="0 0 200 120" xmlns="http://www.w3.org/2000/svg">
//      <path d="M20,20 L180,20 M20,50 L180,50 M20,80 L180,80" stroke="grey" />
//
//      <text dominant-baseline="auto" x="30" y="20">Auto</text>
//      <text dominant-baseline="middle" x="30" y="50">Middle</text>
//      <text dominant-baseline="hanging" x="30" y="80">Hanging</text>
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 200, 120}).Append(
		factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(20, 20).L(180, 20).M(20, 50).L(180, 50).M(20, 80).L(180, 80)).Stroke(factoryColor.NewGray()),

		factoryBrowser.NewTagSvgText().DominantBaseline(html.KSvgDominantBaselineAuto).X(30).Y(20).Text("Auto"),
		factoryBrowser.NewTagSvgText().DominantBaseline(html.KSvgDominantBaselineMiddle).X(30).Y(50).Text("Middle"),
		factoryBrowser.NewTagSvgText().DominantBaseline(html.KSvgDominantBaselineAuto).X(30).Y(80).Text("Hanging"),
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
