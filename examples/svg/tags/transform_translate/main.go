// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/transform
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/transform
//
//  <svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
//    <!-- No translation -->
//    <rect x="5" y="5" width="40" height="40" fill="green" />
//
//    <!-- Horizontal translation -->
//    <rect x="5" y="5" width="40" height="40" fill="blue"
//          transform="translate(50)" />
//
//    <!-- Vertical translation -->
//    <rect x="5" y="5" width="40" height="40" fill="red"
//          transform="translate(0 50)" />
//
//    <!-- Both horizontal and vertical translation -->
//    <rect x="5" y="5" width="40" height="40" fill="yellow"
//           transform="translate(50,50)" />
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Append(
		// No translation
		factoryBrowser.NewTagSvgRect().X(5).Y(5).Width(40).Height(40).Fill(factoryColor.NewGreen()),
		factoryBrowser.NewTagSvgRect().X(5).Y(5).Width(40).Height(40).Fill(factoryColor.NewBlue()).Transform(
			factoryBrowser.NewTransform().
				// When y is not defined, use zero
				// Quando y não é definido, use zero
				Translate(50, 0),
		),

		factoryBrowser.NewTagSvgRect().X(5).Y(5).Width(40).Height(40).Fill(factoryColor.NewRed()).Transform(factoryBrowser.NewTransform().Translate(0, 50)),
		factoryBrowser.NewTagSvgRect().X(5).Y(5).Width(40).Height(40).Fill(factoryColor.NewYellow()).Transform(factoryBrowser.NewTransform().Translate(50, 50)),
	)

	stage.Append(s1)

	<-done
}
