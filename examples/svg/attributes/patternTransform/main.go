// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/patternTransform
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/patternTransform
//
//  <svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
//    <!-- Apply a transform on the tile -->
//    <pattern id="p1" width=".25" height=".25"
//             patternTransform="rotate(20)
//                               skewX(30)
//                               scale(1 0.5)">
//      <circle cx="10" cy="10" r="10" />
//    </pattern>
//
//    <!-- Apply the transformed pattern tile -->
//    <rect x="10" y="10" width="80" height="80"
//          fill="url(#p1)" />
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Append(
		// Apply a transform on the tile
		factoryBrowser.NewTagSvgPattern().Id("p1").Width(0.25).Height(0.25).PatternTransform(factoryBrowser.NewTransform().RotateAngle(20).SkewX(30).Scale(1, 0.5)).Append(
			factoryBrowser.NewTagSvgCircle().Cx(10).Cy(10).R(10),
		),

		// Apply the transformed pattern tile
		factoryBrowser.NewTagSvgRect().X(10).Y(10).Width(80).Height(80).Fill("url(#p1)"),
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
