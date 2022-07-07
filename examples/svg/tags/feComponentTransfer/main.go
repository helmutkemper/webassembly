// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feComponentTransfer
//
//   Notes:
//     * The CSS is inside the example HTML file.
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feComponentTransfer
//
//   Notas:
//     * O CSS está dentro do arquivo HTML de exemplo.
//
// CSS:
//
//  rect {
//    fill: url(#rainbow);
//  }
//
// HTML:
//
//  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 600 300">
//    <defs>
//      <linearGradient id="rainbow" gradientUnits="userSpaceOnUse" x1="0" y1="0" x2="100%" y2="0">
//        <stop offset="0" stop-color="#ff0000"></stop>
//        <stop offset="0.2" stop-color="#ffff00"></stop>
//        <stop offset="0.4" stop-color="#00ff00"></stop>
//        <stop offset="0.6" stop-color="#00ffff"></stop>
//        <stop offset="0.8" stop-color="#0000ff"></stop>
//        <stop offset="1" stop-color="#800080"></stop>
//      </linearGradient>
//      <filter id="identity" x="0" y="0" width="100%" height="100%">
//        <feComponentTransfer>
//          <feFuncR type="identity"></feFuncR>
//          <feFuncG type="identity"></feFuncG>
//          <feFuncB type="identity"></feFuncB>
//          <feFuncA type="identity"></feFuncA>
//        </feComponentTransfer>
//      </filter>
//      <filter id="table" x="0" y="0" width="100%" height="100%">
//        <feComponentTransfer>
//          <feFuncR type="table" tableValues="0 0 1 1"></feFuncR>
//          <feFuncG type="table" tableValues="1 1 0 0"></feFuncG>
//          <feFuncB type="table" tableValues="0 1 1 0"></feFuncB>
//        </feComponentTransfer>
//      </filter>
//      <filter id="discrete" x="0" y="0" width="100%" height="100%">
//        <feComponentTransfer>
//          <feFuncR type="discrete" tableValues="0 0 1 1"></feFuncR>
//          <feFuncG type="discrete" tableValues="1 1 0 0"></feFuncG>
//          <feFuncB type="discrete" tableValues="0 1 1 0"></feFuncB>
//        </feComponentTransfer>
//      </filter>
//      <filter id="linear" x="0" y="0" width="100%" height="100%">
//        <feComponentTransfer>
//          <feFuncR type="linear" slope="0.5" intercept="0"></feFuncR>
//          <feFuncG type="linear" slope="0.5" intercept="0.25"></feFuncG>
//          <feFuncB type="linear" slope="0.5" intercept="0.5"></feFuncB>
//        </feComponentTransfer>
//      </filter>
//      <filter id="gamma" x="0" y="0" width="100%" height="100%">
//        <feComponentTransfer>
//          <feFuncR type="gamma" amplitude="4" exponent="7" offset="0"></feFuncR>
//          <feFuncG type="gamma" amplitude="4" exponent="4" offset="0"></feFuncG>
//          <feFuncB type="gamma" amplitude="4" exponent="1" offset="0"></feFuncB>
//        </feComponentTransfer>
//      </filter>
//    </defs>
//    <g font-weight="bold">
//      <text x="0" y="20">Default</text>
//      <rect x="0" y="30" width="100%" height="20"></rect>
//      <text x="0" y="70">Identity</text>
//      <rect x="0" y="80" width="100%" height="20" style="filter:url(#identity)"></rect>
//      <text x="0" y="120">Table lookup</text>
//      <rect x="0" y="130" width="100%" height="20" style="filter:url(#table)"></rect>
//      <text x="0" y="170">Discrete table lookup</text>
//      <rect x="0" y="180" width="100%" height="20" style="filter:url(#discrete)"></rect>
//      <text x="0" y="220">Linear function</text>
//      <rect x="0" y="230" width="100%" height="20" style="filter:url(#linear)"></rect>
//      <text x="0" y="270">Gamma function</text>
//      <rect x="0" y="280" width="100%" height="20" style="filter:url(#gamma)"></rect>
//    </g>
//  </svg>

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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 600, 300}).Append(
		factoryBrowser.NewTagSvgDefs().Append(
			factoryBrowser.NewTagSvgLinearGradient().Id("rainbow").GradientUnits(html.KSvgGradientUnitsUserSpaceOnUse).X1(0).Y1(0).X2(float32(1.0)).Y2(0).Append(
				factoryBrowser.NewTagSvgStop().Offset(0.0).StopColor("#ff0000"),
				factoryBrowser.NewTagSvgStop().Offset(0.2).StopColor("#ffff00"),
				factoryBrowser.NewTagSvgStop().Offset(0.4).StopColor("#00ff00"),
				factoryBrowser.NewTagSvgStop().Offset(0.6).StopColor("#00ffff"),
				factoryBrowser.NewTagSvgStop().Offset(0.8).StopColor("#0000ff"),
				factoryBrowser.NewTagSvgStop().Offset(1.0).StopColor("#800080"),
			),

			factoryBrowser.NewTagSvgFilter().Id("identity").X(0).Y(0).Width(float32(1.0)).Height(float32(1.0)).Append(
				factoryBrowser.NewTagSvgFeComponentTransfer().Append(
					factoryBrowser.NewTagSvgFeFuncR().Type(html.KSvgTypeFeFuncIdentity),
					factoryBrowser.NewTagSvgFeFuncG().Type(html.KSvgTypeFeFuncIdentity),
					factoryBrowser.NewTagSvgFeFuncB().Type(html.KSvgTypeFeFuncIdentity),
					factoryBrowser.NewTagSvgFeFuncA().Type(html.KSvgTypeFeFuncIdentity),
				),
			),

			factoryBrowser.NewTagSvgFilter().Id("table").X(0).Y(0).Width(float32(1.0)).Height(float32(1.0)).Append(
				factoryBrowser.NewTagSvgFeComponentTransfer().Append(
					factoryBrowser.NewTagSvgFeFuncR().Type(html.KSvgTypeFeFuncTable).TableValues([]float64{0, 0, 1, 1}),
					factoryBrowser.NewTagSvgFeFuncG().Type(html.KSvgTypeFeFuncTable).TableValues([]float64{1, 1, 0, 0}),
					factoryBrowser.NewTagSvgFeFuncB().Type(html.KSvgTypeFeFuncTable).TableValues([]float64{0, 1, 1, 0}),
				),
			),

			factoryBrowser.NewTagSvgFilter().Id("discrete").X(0).Y(0).Width(float32(1.0)).Height(float32(1.0)).Append(
				factoryBrowser.NewTagSvgFeComponentTransfer().Append(
					factoryBrowser.NewTagSvgFeFuncR().Type(html.KSvgTypeFeFuncDiscrete).TableValues([]float64{0, 0, 1, 1}),
					factoryBrowser.NewTagSvgFeFuncG().Type(html.KSvgTypeFeFuncDiscrete).TableValues([]float64{1, 1, 0, 0}),
					factoryBrowser.NewTagSvgFeFuncB().Type(html.KSvgTypeFeFuncDiscrete).TableValues([]float64{0, 1, 1, 0}),
				),
			),

			factoryBrowser.NewTagSvgFilter().Id("linear").X(0).Y(0).Width(float32(1.0)).Height(float32(1.0)).Append(
				factoryBrowser.NewTagSvgFeComponentTransfer().Append(
					factoryBrowser.NewTagSvgFeFuncR().Type(html.KSvgTypeFeFuncLinear).Slope(0.5).Intercept(0.0),
					factoryBrowser.NewTagSvgFeFuncR().Type(html.KSvgTypeFeFuncLinear).Slope(0.5).Intercept(0.25),
					factoryBrowser.NewTagSvgFeFuncR().Type(html.KSvgTypeFeFuncLinear).Slope(0.5).Intercept(0.5),
				),
			),

			factoryBrowser.NewTagSvgFilter().Id("gamma").X(0).Y(0).Width(float32(1.0)).Height(float32(1.0)).Append(
				factoryBrowser.NewTagSvgFeComponentTransfer().Append(
					factoryBrowser.NewTagSvgFeFuncR().Type(html.KSvgTypeFeFuncGamma).Amplitude(4).Exponent(7).Offset(0),
					factoryBrowser.NewTagSvgFeFuncR().Type(html.KSvgTypeFeFuncGamma).Amplitude(4).Exponent(4).Offset(0),
					factoryBrowser.NewTagSvgFeFuncR().Type(html.KSvgTypeFeFuncGamma).Amplitude(4).Exponent(1).Offset(0),
				),
			),
		),

		factoryBrowser.NewTagSvgG().FontWeight(html.KFontWeightRuleBold).Append(

			factoryBrowser.NewTagSvgText().X(0).Y(20).Text("Default"),
			factoryBrowser.NewTagSvgRect().X(0).Y(30).Width(float32(1.0)).Height(20),

			factoryBrowser.NewTagSvgText().X(0).Y(70).Text("Identity"),
			factoryBrowser.NewTagSvgRect().X(0).Y(80).Width(float32(1.0)).Height(20).Style("filter:url(#identity)"),

			factoryBrowser.NewTagSvgText().X(0).Y(120).Text("Table lookup"),
			factoryBrowser.NewTagSvgRect().X(0).Y(130).Width(float32(1.0)).Height(20).Style("filter:url(#table)"),

			factoryBrowser.NewTagSvgText().X(0).Y(170).Text("Discrete table lookup"),
			factoryBrowser.NewTagSvgRect().X(0).Y(180).Width(float32(1.0)).Height(20).Style("filter:url(#discrete)"),

			factoryBrowser.NewTagSvgText().X(0).Y(220).Text("Linear function"),
			factoryBrowser.NewTagSvgRect().X(0).Y(230).Width(float32(1.0)).Height(20).Style("filter:url(#linear)"),

			factoryBrowser.NewTagSvgText().X(0).Y(270).Text("Gamma function"),
			factoryBrowser.NewTagSvgRect().X(0).Y(280).Width(float32(1.0)).Height(20).Style("filter:url(#gamma)"),
		),
	)

	stage.Append(s1)

	<-done
}
