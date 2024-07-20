// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/dy
//
//   Notes:
//     * The CSS is inside the example HTML file.
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/dy
//
//   Notas:
//     * O CSS está dentro do arquivo HTML de exemplo.
//
// CSS:
//
//  line {
//    stroke: red;
//    stroke-width: .5px;
//    stroke-dasharray: 3px;
//  }
//
// HTML:
//
//  <svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
//    <!-- Lines materialized the position of the glyphs -->
//    <line x1="10%" x2="10%"  y1="0"   y2="100%" />
//    <line x1="0"   x2="100%" y1="30%" y2="30%"  />
//    <line x1="0"   x2="100%" y1="80%" y2="80%"  />
//
//    <!-- Some reference text -->
//    <text x="10%" y="30%" fill="grey">SVG</text>
//
//    <!-- The same text with a shift along the y-axis -->
//    <text dy="50%" x="10%" y="30%">SVG</text>
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
		// Lines materialized the position of the glyphs
		factoryBrowser.NewTagSvgLine().X1(float32(0.1)).X2(float32(0.1)).Y1(float32(0.0)).Y2(float32(1.0)),
		factoryBrowser.NewTagSvgLine().X1(float32(0.0)).X2(float32(1.0)).Y1(float32(0.3)).Y2(float32(0.3)),
		factoryBrowser.NewTagSvgLine().X1(float32(0.0)).X2(float32(1.0)).Y1(float32(0.8)).Y2(float32(0.8)),

		// Some reference text
		factoryBrowser.NewTagSvgText().X(float32(0.1)).Y(float32(0.3)).Fill(factoryColor.NewGray()).Text("SVG"),

		// The same text with a shift along the x-axis
		factoryBrowser.NewTagSvgText().Dy(float32(0.5)).X(float32(0.1)).Y(float32(0.3)).Text("SVG"),
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
