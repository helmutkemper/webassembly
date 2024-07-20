// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/seed
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/seed
//
//  <svg viewBox="0 0 420 200" xmlns="http://www.w3.org/2000/svg">
//    <filter id="noise1" x="0" y="0" width="100%" height="100%">
//      <feTurbulence baseFrequency="0.025" seed="0" />
//    </filter>
//    <filter id="noise2" x="0" y="0" width="100%" height="100%">
//      <feTurbulence baseFrequency="0.025" seed="100" />
//    </filter>
//
//    <rect x="0" y="0" width="200" height="200" style="filter:url(#noise1);" />
//    <rect x="0" y="0" width="200" height="200" style="filter:url(#noise2); transform: translateX(220px);" />
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 420, 200}).Append(
		factoryBrowser.NewTagSvgFilter().Id("noise1").X(0).Y(0).Width(float32(1)).Height(float32(1)).Append(
			factoryBrowser.NewTagSvgFeTurbulence().BaseFrequency(0.025).Seed(0),
		),
		factoryBrowser.NewTagSvgFilter().Id("noise2").X(0).Y(0).Width(float32(1)).Height(float32(1)).Append(
			factoryBrowser.NewTagSvgFeTurbulence().BaseFrequency(0.025).Seed(100),
		),

		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(200).Height(200).Style("filter:url(#noise1);"),
		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(200).Height(200).Style("filter:url(#noise2); transform: translateX(220px);"),
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
