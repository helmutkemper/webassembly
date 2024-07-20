// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/x
//
//   Notes:
//     * The CSS is inside the example HTML file.
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/x
//
//   Notas:
//     * O CSS está dentro do arquivo HTML de exemplo.
//
// CSS:
//
//  text {
//    font: 40px sans-serif;
//  }
//
//  line {
//    fill: none;
//    stroke: red;
//    stroke-width: .5px;
//    stroke-dasharray: 2px;
//  }
//
// HTML:
//
//  <svg viewBox="0 0 200 100" xmlns="http://www.w3.org/2000/svg">
//    <!-- horizontal line to materialized the text base line -->
//    <line x1="0" y1="40%" x2="100%" y2="40%" />
//    <line x1="0" y1="60%" x2="100%" y2="60%" />
//    <line x1="0" y1="80%" x2="100%" y2="80%" />
//
//    <!-- vertical line to materialized the x positioning -->
//    <line x1="5%"  y1="0" x2="5%"  y2="100%" />
//    <line x1="55%" y1="0" x2="55%" y2="100%" />
//
//    <!-- y with a single value -->
//    <text y="40%" x="5%">SVG</text>
//
//    <!-- y with multiple values -->
//    <text y="40%,60%,80%" x="55%">SVG</text>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 200, 100}).Append(
		// horizontal line to materialized the text base line
		factoryBrowser.NewTagSvgLine().X1(0).Y1(float32(0.4)).X2(float32(1.0)).Y2(float32(0.4)),
		factoryBrowser.NewTagSvgLine().X1(0).Y1(float32(0.6)).X2(float32(1.0)).Y2(float32(0.6)),
		factoryBrowser.NewTagSvgLine().X1(0).Y1(float32(0.8)).X2(float32(1.0)).Y2(float32(0.8)),

		// vertical line to materialized the x positioning
		factoryBrowser.NewTagSvgLine().X1(float32(0.05)).Y1(0).X2(float32(0.05)).Y2(float32(1)),
		factoryBrowser.NewTagSvgLine().X1(float32(0.55)).Y1(0).X2(float32(0.55)).Y2(float32(1)),

		// y with a single value
		factoryBrowser.NewTagSvgText().Y(float32(0.4)).X(float32(0.05)).Text("SVG"),

		// y with multiple values
		factoryBrowser.NewTagSvgText().Y([]float32{0.4, 0.6, 0.8}).X(float32(0.55)).Text("SVG"),
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
