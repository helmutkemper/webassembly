// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/result
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/result
//
//  <svg viewBox="0 0 220 220" xmlns="http://www.w3.org/2000/svg">
//    <filter id="displacementFilter">
//      <feTurbulence type="turbulence" baseFrequency="0.05"
//          numOctaves="2" result="turbulence"/>
//      <feDisplacementMap in2="turbulence" in="SourceGraphic"
//          scale="50" xChannelSelector="R" yChannelSelector="G"/>
//    </filter>
//
//    <circle cx="100" cy="100" r="100"
//        style="filter: url(#displacementFilter)"/>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 220, 220}).Append(
		factoryBrowser.NewTagSvgFilter().Id("displacementFilter").Append(
			factoryBrowser.NewTagSvgFeTurbulence().Type(html.KSvgTypeTurbulenceTurbulence).BaseFrequency(0.05).NumOctaves(2).Result("turbulence"),
			factoryBrowser.NewTagSvgFeDisplacementMap().In2("turbulence").In(html.KSvgInSourceGraphic).Scale(50).XChannelSelector(html.KSvgChannelSelectorR).YChannelSelector(html.KSvgChannelSelectorG),
		),

		factoryBrowser.NewTagSvgCircle().Cx(100).Cy(100).R(100).Style("filter: url(#displacementFilter)"),
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
