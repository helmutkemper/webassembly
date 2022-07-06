// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/dominant-baseline
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/dominant-baseline
//
//  <svg width="400" height="300" viewBox="0 0 300 300"
//      xmlns="http://www.w3.org/2000/svg">
//
//    <!-- Materialization of anchors -->
//    <path d="M60,20 L60,270
//             M30,20 L400,20
//             M30,70 L400,70
//             M30,120 L400,120
//             M30,170 L400,170
//             M30,220 L400,220
//             M30,270 L400,270" stroke="grey" />
//
//      <!-- Anchors in action -->
//      <text dominant-baseline="auto" x="70" y="20">auto</text>
//      <text dominant-baseline="middle" x="70" y="70">middle</text>
//      <text dominant-baseline="hanging" x="70" y="170">hanging</text>
//      <text dominant-baseline="mathematical" x="70" y="220">mathematical</text>
//      <text dominant-baseline="text-top" x="70" y="270">text-top</text>
//
//    <!-- Materialization of anchors -->
//    <circle cx="60" cy="20" r="3" fill="red" />
//    <circle cx="60" cy="70" r="3" fill="red" />
//    <circle cx="60" cy="120" r="3" fill="red" />
//    <circle cx="60" cy="170" r="3" fill="red" />
//    <circle cx="60" cy="220" r="3" fill="red" />
//    <circle cx="60" cy="270" r="3" fill="red" />
//
//    <style><![CDATA[
//    text {
//      font: bold 30px Verdana, Helvetica, Arial, sans-serif;
//    }
//    ]]></style>
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().Width(400).Height(300).ViewBox([]float64{0, 0, 300, 300}).Append(
		// Materialization of anchors
		factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(60, 20).L(60, 270).M(30, 20).L(400, 20).M(30, 70).L(400, 70).M(30, 120).L(400, 120).M(30, 170).L(400, 170).M(30, 220).L(400, 220).M(30, 270).L(400, 270)).Stroke(factoryColor.NewGray()),

		// Anchors in action
		factoryBrowser.NewTagSvgText().DominantBaseline(html.KSvgDominantBaselineAuto).X(70).Y(20).Text("auto"),
		factoryBrowser.NewTagSvgText().DominantBaseline(html.KSvgDominantBaselineMiddle).X(70).Y(70).Text("middle"),
		factoryBrowser.NewTagSvgText().DominantBaseline(html.KSvgDominantBaselineHanging).X(70).Y(170).Text("hanging"),
		factoryBrowser.NewTagSvgText().DominantBaseline(html.KSvgDominantBaselineMathematical).X(70).Y(220).Text("mathematical"),
		factoryBrowser.NewTagSvgText().DominantBaseline(html.KSvgDominantBaselineTextTop).X(70).Y(270).Text("text-top"),

		// Materialization of anchors
		factoryBrowser.NewTagSvgCircle().Cx(60).Cy(20).R(3).Fill(factoryColor.NewRed()),
		factoryBrowser.NewTagSvgCircle().Cx(60).Cy(70).R(3).Fill(factoryColor.NewRed()),
		factoryBrowser.NewTagSvgCircle().Cx(60).Cy(120).R(3).Fill(factoryColor.NewRed()),
		factoryBrowser.NewTagSvgCircle().Cx(60).Cy(170).R(3).Fill(factoryColor.NewRed()),
		factoryBrowser.NewTagSvgCircle().Cx(60).Cy(220).R(3).Fill(factoryColor.NewRed()),
		factoryBrowser.NewTagSvgCircle().Cx(60).Cy(270).R(3).Fill(factoryColor.NewRed()),

		factoryBrowser.NewTagSvgStyle().Style(
			//"<![CDATA[\n"+
			"text {\n"+
				"font: bold 30px Verdana, Helvetica, Arial, sans-serif;\n"+
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
