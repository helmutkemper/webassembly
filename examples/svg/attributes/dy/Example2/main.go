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
//  <svg viewBox="0 0 150 100" xmlns="http://www.w3.org/2000/svg">
//    <!-- Horizontal lines -->
//    <line x1="0" x2="100%" y1="30" y2="30" />
//    <line x1="0" x2="100%" y1="40" y2="40" />
//    <line x1="0" x2="100%" y1="50" y2="50" />
//    <line x1="0" x2="100%" y1="60" y2="60" />
//
//    <!-- Vertical lines -->
//    <line x1="10" x2="10" y1="0" y2="100%" />
//    <line x1="50" x2="50" y1="0" y2="100%" />
//    <line x1="90" x2="90" y1="0" y2="100%" />
//
//    <!-- Behaviors change based on the number of values in the attributes -->
//    <text dy="20"      x="10" y="30">SVG</text>
//    <text dy="0 10"    x="50" y="30">SVG</text>
//    <text dy="0 10 20" x="90" y="30">SVG</text>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 150, 100}).Append(
		// Horizontal lines
		factoryBrowser.NewTagSvgLine().X1(float32(0.0)).X2(float32(1.0)).Y1(30).Y2(30),
		factoryBrowser.NewTagSvgLine().X1(float32(0.0)).X2(float32(1.0)).Y1(40).Y2(40),
		factoryBrowser.NewTagSvgLine().X1(float32(0.0)).X2(float32(1.0)).Y1(50).Y2(50),
		factoryBrowser.NewTagSvgLine().X1(float32(0.0)).X2(float32(1.0)).Y1(60).Y2(60),

		// Vertical lines
		factoryBrowser.NewTagSvgLine().X1(10).X2(10).Y1(0).Y2(float32(1.0)),
		factoryBrowser.NewTagSvgLine().X1(50).X2(50).Y1(0).Y2(float32(1.0)),
		factoryBrowser.NewTagSvgLine().X1(90).X2(90).Y1(0).Y2(float32(1.0)),

		// Behaviors change based on the number of values in the attributes
		factoryBrowser.NewTagSvgText().Dy(20).X(10).Y(30).Text("SVG"),
		factoryBrowser.NewTagSvgText().Dy([]float64{0, 10}).X(50).Y(30).Text("SVG"),
		factoryBrowser.NewTagSvgText().Dy([]float64{0, 10, 20}).X(90).Y(30).Text("SVG"),
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
