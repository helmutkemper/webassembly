// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/paint-order
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/paint-order
//
//  <svg xmlns="http://www.w3.org/2000/svg" width="400" height="200">
//    <linearGradient id="g" x1="0" y1="0" x2="0" y2="1">
//      <stop stop-color="#888"/>
//      <stop stop-color="#ccc" offset="1"/>
//    </linearGradient>
//    <rect width="400" height="200" fill="url(#g)"/>
//    <g fill="crimson" stroke="white" stroke-width="6" stroke-linejoin="round"
//       text-anchor="middle" font-family="sans-serif" font-size="50px" font-weight="bold">
//      <text x="200" y="75">stroke over</text>
//      <text x="200" y="150" paint-order="stroke" id="stroke-under">stroke under</text>
//    </g>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryFontFamily"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().Width(400).Height(200).Append(
		factoryBrowser.NewTagSvgLinearGradient().Id("g").X1(0).Y1(0).X2(0).Y2(1).Append(
			factoryBrowser.NewTagSvgStop().StopColor("#888"),
			factoryBrowser.NewTagSvgStop().StopColor("#CCC").Offset(1),
		),
		factoryBrowser.NewTagSvgRect().Width(400).Height(200).Fill("url(#g)"),
		factoryBrowser.NewTagSvgG().Fill(factoryColor.NewCrimson()).Stroke(factoryColor.NewWhite()).StrokeWidth(6).StrokeLineJoin(html.KSvgStrokeLinejoinRound).TextAnchor(html.KSvgTextAnchorMiddle).FontFamily(factoryFontFamily.NewSansSerif()).FontSize("50px").FontWeight(html.KFontWeightRuleBold).Append(
			factoryBrowser.NewTagSvgText().X(200).Y(75).Text("stroke over"),
			factoryBrowser.NewTagSvgText().X(200).Y(150).Text("stroke over").PaintOrder(html.KSvgPaintOrderStroke).Id("stroke-under").Text("stroke under"),
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
