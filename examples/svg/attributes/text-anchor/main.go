// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/text-anchor
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/text-anchor
//
//  <svg viewBox="0 0 120 120" xmlns="http://www.w3.org/2000/svg">
//    <!-- Materialization of anchors -->
//    <path d="M60,15 L60,110 M30,40 L90,40 M30,75 L90,75 M30,110 L90,110" stroke="grey" />
//
//    <!-- Anchors in action -->
//    <text text-anchor="start" x="60" y="40">A</text>
//    <text text-anchor="middle" x="60" y="75">A</text>
//    <text text-anchor="end" x="60" y="110">A</text>
//
//    <!-- Materialization of anchors -->
//    <circle cx="60" cy="40" r="3" fill="red" />
//    <circle cx="60" cy="75" r="3" fill="red" />
//    <circle cx="60" cy="110" r="3" fill="red" />
//
//    <style><![CDATA[
//    text {
//      font: bold 36px Verdana, Helvetica, Arial, sans-serif;
//    }
//    ]]></style>
//  </svg>

//go:build js

// bug: o original não apararece nada

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 120, 120}).Append(
		// Materialization of anchors
		factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(60, 15).L(60, 110).M(30, 40).L(90, 40).M(30, 75).L(90, 75).M(30, 110).L(90, 110)).Stroke(factoryColor.NewGray()),

		// Anchors in action
		factoryBrowser.NewTagSvgText().TextAnchor(html.KSvgTextAnchorStart).X(60).Y(40).Text("A"),
		factoryBrowser.NewTagSvgText().TextAnchor(html.KSvgTextAnchorMiddle).X(60).Y(75).Text("A"),
		factoryBrowser.NewTagSvgText().TextAnchor(html.KSvgTextAnchorEnd).X(60).Y(110).Text("A"),

		// Materialization of anchors
		factoryBrowser.NewTagSvgCircle().Cx(60).Cy(40).R(3).Fill(factoryColor.NewRed()),
		factoryBrowser.NewTagSvgCircle().Cx(60).Cy(75).R(3).Fill(factoryColor.NewRed()),
		factoryBrowser.NewTagSvgCircle().Cx(60).Cy(110).R(3).Fill(factoryColor.NewRed()),

		factoryBrowser.NewTagSvgStyle().Style(
			//"<![CDATA[\n" +
			"text {\n"+
				"font: bold 36px Verdana, Helvetica, Arial, sans-serif;\n"+
				"}\n",
			//"]]>",
		),
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
