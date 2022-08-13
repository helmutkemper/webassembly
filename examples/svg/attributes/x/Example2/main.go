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
//    <line x1="0" y1="90%" x2="100%" y2="90%" />
//
//    <!-- vertical line to materialized the x positioning -->
//    <line x1="25%" y1="0" x2="25%" y2="100%" />
//    <line x1="50%" y1="0" x2="50%" y2="100%" />
//    <line x1="75%" y1="0" x2="75%" y2="100%" />
//
//    <!-- x with a single value -->
//    <text y="40%" x="50%">SVG</text>
//
//    <!-- x with multiple values -->
//    <text y="90%" x="25%, 50%, 75%">SVG</text>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 200, 100}).Append(
		// horizontal line to materialized the text base line
		factoryBrowser.NewTagSvgLine().X1(0).Y1(float32(0.4)).X2(float32(1.0)).Y2(float32(0.4)),
		factoryBrowser.NewTagSvgLine().X1(0).Y1(float32(0.9)).X2(float32(1.0)).Y2(float32(0.9)),

		// vertical line to materialized the x positioning
		factoryBrowser.NewTagSvgLine().X1(float32(0.25)).Y1(0).X2(float32(0.25)).Y2(float32(1)),
		factoryBrowser.NewTagSvgLine().X1(float32(0.50)).Y1(0).X2(float32(0.50)).Y2(float32(1)),
		factoryBrowser.NewTagSvgLine().X1(float32(0.75)).Y1(0).X2(float32(0.75)).Y2(float32(1)),

		// x with a single value
		factoryBrowser.NewTagSvgText().Y(float32(0.4)).X(float32(0.5)).Text("SVG"),

		// x with multiple values
		factoryBrowser.NewTagSvgText().Y(float32(0.9)).X([]float32{0.25, 0.50, 0.75}).Text("SVG"),
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
