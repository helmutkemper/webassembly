// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feMerge
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feMerge
//
//    <svg width="200" height="200"
//    xmlns="http://www.w3.org/2000/svg">
//
//    <filter id="feOffset" x="-40" y="-20" width="100" height="200">
//      <feOffset in="SourceGraphic" dx="60" dy="60" />
//      <feGaussianBlur stdDeviation="5" result="blur2" />
//      <feMerge>
//        <feMergeNode in="blur2" />
//        <feMergeNode in="SourceGraphic" />
//      </feMerge>
//    </filter>
//
//    <rect x="40" y="40" width="100" height="100"
//      style="stroke: #000000; fill: green; filter: url(#feOffset);" />
//  </svg>

//fixme: bug. Filter doesn't seem to work as expected

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().
		Width(200).
		Height(200).
		Append(

			factoryBrowser.NewTagSvgFilter().
				Id("feOffset").
				X(-40).
				Y(-20).
				Width(100).
				Height(200).
				Append(

					factoryBrowser.NewTagSvgFeOffset().
						In(html.KSvgInSourceGraphic).
						Dx(60).
						Dy(60),

					factoryBrowser.NewTagSvgFeGaussianBlur().
						StdDeviation(5).
						Result("blur2").
						Append(

							factoryBrowser.NewTagSvgFeMerge().
								Append(

									factoryBrowser.NewTagSvgFeMergeNode().In("blur2"),
									factoryBrowser.NewTagSvgFeMergeNode().In(html.KSvgInSourceGraphic),
								),
						),
				),

			factoryBrowser.NewTagSvgRect().
				X(40).
				Y(40).
				Width(100).
				Height(100).
				Style("stroke: #000000; fill: green; filter: url(#feOffset);"),
		)

	stage.Append(s1)

	<-done
}
