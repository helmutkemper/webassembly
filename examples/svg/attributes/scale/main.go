// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/scale
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/scale
//
//  <svg viewBox="0 0 480 220" xmlns="http://www.w3.org/2000/svg">
//    <filter id="displacementFilter" x="-20%" y="-20%" width="140%" height="140%">
//      <feTurbulence type="turbulence" baseFrequency="0.05" numOctaves="2" result="turbulence"/>
//      <feDisplacementMap in2="turbulence" in="SourceGraphic" scale="5"/>
//    </filter>
//    <filter id="displacementFilter2" x="-20%" y="-20%" width="140%" height="140%">
//      <feTurbulence type="turbulence" baseFrequency="0.05" numOctaves="2" result="turbulence"/>
//      <feDisplacementMap in2="turbulence" in="SourceGraphic" scale="50"/>
//    </filter>
//
//    <circle cx="100" cy="100" r="80" style="filter: url(#displacementFilter);""/>
//    <circle cx="100" cy="100" r="80" style="filter: url(#displacementFilter2); transform: translateX(240px);""/>
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
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 480, 220}).Append(
		factoryBrowser.NewTagSvgFilter().Id("displacementFilter").X(float32(-0.2)).Y(float32(-0.2)).Width(float32(1.4)).Height(float32(1.4)).Append(
			factoryBrowser.NewTagSvgFeTurbulence().Type(html.KSvgTypeTurbulenceTurbulence).BaseFrequency(0.05).NumOctaves(2).Result("turbulence"),
			factoryBrowser.NewTagSvgFeDisplacementMap().In2("turbulence").In(html.KSvgInSourceGraphic).Scale(5),
		),
		factoryBrowser.NewTagSvgFilter().Id("displacementFilter2").X(float32(-0.2)).Y(float32(-0.2)).Width(float32(1.4)).Height(float32(1.4)).Append(
			factoryBrowser.NewTagSvgFeTurbulence().Type(html.KSvgTypeTurbulenceTurbulence).BaseFrequency(0.05).NumOctaves(2).Result("turbulence"),
			factoryBrowser.NewTagSvgFeDisplacementMap().In2("turbulence").In(html.KSvgInSourceGraphic).Scale(50),
		),

		factoryBrowser.NewTagSvgCircle().Cx(100).Cy(100).R(80).Style("filter: url(#displacementFilter);"),
		factoryBrowser.NewTagSvgCircle().Cx(100).Cy(100).R(80).Style("filter: url(#displacementFilter2); transform: translateX(240px);"),
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
