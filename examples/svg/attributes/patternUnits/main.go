// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/patternUnits
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/patternUnits
//
//  <svg viewBox="0 0 200 100" xmlns="http://www.w3.org/2000/svg">
//    <!-- All geometry properties are relative to the current user space -->
//    <pattern id="p1" x="12.5" y="12.5" width="25" height="25"
//             patternUnits="userSpaceOnUse">
//      <circle cx="10" cy="10" r="10" />
//    </pattern>
//
//    <!-- All geometry properties are relative to the target bounding box -->
//    <pattern id="p2" x=".125" y=".125" width=".25" height=".25"
//             patternUnits="objectBoundingBox">
//      <circle cx="10" cy="10" r="10" />
//    </pattern>
//
//    <!-- Left square with user space tiles -->
//    <rect x="10"  y="10" width="80" height="80"
//          fill="url(#p1)" />
//
//    <!-- Right square with bounding box tiles -->
//    <rect x="110" y="10" width="80" height="80"
//          fill="url(#p2)" />
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
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 200, 100}).Append(
		// All geometry properties are relative to the current user space
		factoryBrowser.NewTagSvgPattern().Id("p1").X(12.5).Y(12.5).Width(25).Height(25).PatternUnits(html.KSvgUnitsUserSpaceOnUse).Append(
			factoryBrowser.NewTagSvgCircle().Cx(10).Cy(10).R(10),
		),

		// All geometry properties are relative to the target bounding box
		factoryBrowser.NewTagSvgPattern().Id("p2").X(0.125).Y(0.125).Width(0.25).Height(0.25).PatternUnits(html.KSvgUnitsObjectBoundingBox).Append(
			factoryBrowser.NewTagSvgCircle().Cx(10).Cy(10).R(10),
		),

		// Left square with user space tiles
		factoryBrowser.NewTagSvgRect().X(10).Y(10).Width(80).Height(80).Fill("url(#p1)"),

		// Right square with bounding box tiles
		factoryBrowser.NewTagSvgRect().X(110).Y(10).Width(80).Height(80).Fill("url(#p2)"),
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
