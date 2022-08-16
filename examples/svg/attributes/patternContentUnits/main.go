// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/patternContentUnits
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/patternContentUnits
//
//  <svg viewBox="0 0 200 100" xmlns="http://www.w3.org/2000/svg">
//    <!--
//    A pattern tile that content coordinates and values are
//    computed against the current coordinate user space.
//    Note that the size of the tile is computed against
//    the bounding box of the target element
//    -->
//    <pattern id="p1" width="20%" height="20%"
//             patternContentUnits="userSpaceOnUse">
//      <circle cx="10" cy="10" r="10" />
//    </pattern>
//
//    <!--
//    A pattern tile that content coordinates and values are
//    computed against the bounding box of the target element.
//    Note that the size of the tile is also computed against
//    the bounding box of the target element
//    -->
//    <pattern id="p2" width="20%" height="20%"
//             patternContentUnits="objectBoundingBox">
//      <circle cx=".1" cy=".1" r=".1" />
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
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 200, 100}).Append(
		//  A pattern tile that content coordinates and values are
		//  computed against the current coordinate user space.
		//  Note that the size of the tile is computed against
		//  the bounding box of the target element

		factoryBrowser.NewTagSvgPattern().Id("p1").Width(float32(0.2)).Height(float32(0.2)).PatternContentUnits(html.KSvgUnitsUserSpaceOnUse).Append(
			factoryBrowser.NewTagSvgCircle().Cx(10).Cy(10).R(10),
		),

		//  A pattern tile that content coordinates and values are
		//  computed against the bounding box of the target element.
		//  Note that the size of the tile is also computed against
		//  the bounding box of the target element
		factoryBrowser.NewTagSvgPattern().Id("p2").Width(float32(0.2)).Height(float32(0.2)).PatternContentUnits(html.KSvgUnitsObjectBoundingBox).Append(
			factoryBrowser.NewTagSvgCircle().Cx(0.1).Cy(0.1).R(0.1),
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
