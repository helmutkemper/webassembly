// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/elevation
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/elevation
//
//  <svg viewBox="0 0 440 200" xmlns="http://www.w3.org/2000/svg">
//    <filter id="distantLight1">
//      <feDiffuseLighting>
//        <feDistantLight elevation="0" />
//      </feDiffuseLighting>
//    </filter>
//    <filter id="distantLight2">
//      <feDiffuseLighting>
//        <feDistantLight elevation="45" />
//      </feDiffuseLighting>
//    </filter>
//
//    <circle cx="100" cy="100" r="80" style="filter: url(#distantLight1);" />
//    <circle cx="100" cy="100" r="80" style="filter: url(#distantLight2); transform: translateX(240px);" />
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 440, 200}).Append(
		factoryBrowser.NewTagSvgFilter().Id("distantLight1").Append(
			factoryBrowser.NewTagSvgFeDiffuseLighting().Append(
				factoryBrowser.NewTagSvgFeDistantLight().Elevation(0),
			),
		),
		factoryBrowser.NewTagSvgFilter().Id("distantLight2").Append(
			factoryBrowser.NewTagSvgFeDiffuseLighting().Append(
				factoryBrowser.NewTagSvgFeDistantLight().Elevation(45),
			),
		),

		factoryBrowser.NewTagSvgCircle().Cx(100).Cy(100).R(80).Style("filter: url(#distantLight1);"),
		factoryBrowser.NewTagSvgCircle().Cx(100).Cy(100).R(80).Style("filter: url(#distantLight2); transform: translateX(240px);"),
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
