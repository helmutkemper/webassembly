// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/alignment-baseline
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/alignment-baseline
//
//  <svg width="300" height="120" viewBox="0 0 300 120"
//       xmlns="http://www.w3.org/2000/svg">
//
//     <!-- Materialization of anchors -->
//     <path d="M60,10 L60,110
//               M30,10 L300,10
//               M30,65 L300,65
//               M30,110 L300,110
//               " stroke="grey" />
//
//     <!-- Anchors in action -->
//     <text alignment-baseline="hanging"
//           x="60" y="10">A hanging</text>
//
//     <text alignment-baseline="middle"
//           x="60" y="65">A middle</text>
//
//     <text alignment-baseline="baseline"
//           x="60" y="110">A baseline</text>
//
//     <!-- Materialization of anchors -->
//     <circle cx="60" cy="10" r="3" fill="red" />
//     <circle cx="60" cy="65" r="3" fill="red" />
//     <circle cx="60" cy="110" r="3" fill="red" />
//
// <style><![CDATA[
// text{
//     font: bold 36px Verdana, Helvetica, Arial, sans-serif;
// }
// ]]></style>
// </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().Width(300).Height(120).ViewBox([]float64{0, 0, 300, 120}).Append(

		// Materialization of anchors
		factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(60, 10).L(60, 110).M(30, 10).L(300, 10).M(30, 65).L(300, 65).M(30, 110).L(300, 110)).Stroke(factoryColor.NewGray()),

		// Anchors in action
		factoryBrowser.NewTagSvgText().AlignmentBaseline(html.KSvgAlignmentBaselineHanging).X(60).Y(10).Text("A hanging"),
		factoryBrowser.NewTagSvgText().AlignmentBaseline(html.KSvgAlignmentBaselineMiddle).X(60).Y(65).Text("A middle"),
		factoryBrowser.NewTagSvgText().AlignmentBaseline(html.KSvgAlignmentBaselineBaseline).X(60).Y(115).Text("A baseline"),

		// Materialization of anchors
		factoryBrowser.NewTagSvgCircle().Cx(60).Cy(10).R(3).Fill(factoryColor.NewRed()),
		factoryBrowser.NewTagSvgCircle().Cx(60).Cy(65).R(3).Fill(factoryColor.NewRed()),
		factoryBrowser.NewTagSvgCircle().Cx(60).Cy(110).R(3).Fill(factoryColor.NewRed()),

		factoryBrowser.NewTagSvgStyle().Style(
			//"<![CDATA[\n"+
			"text{\n"+
				"font: bold 36px Verdana, Helvetica, Arial, sans-serif;\n"+
				"}\n"),
		//"]]>"),
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
