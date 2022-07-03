// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/begin
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/begin
//
//  <svg width="120" height="120"  viewBox="0 0 120 120"
//       xmlns="http://www.w3.org/2000/svg" version="1.1"
//       xmlns:xlink="http://www.w3.org/1999/xlink">
//
//      <!-- animated rectangle -->
//      <rect x="10" y="35" height="15" width="0">
//          <animate attributeType="XML" attributeName="width" from="0" to="100"
//                   begin="startButton.click" dur="8s"
//                   fill="freeze" />
//      </rect>
//
//      <!-- trigger -->
//      <rect id="startButton" style="cursor:pointer;"
//            x="19.5" y="62.5" rx="5" height="25" width="80"
//            fill="#EFEFEF" stroke="black" stroke-width="1" />
//
//      <text x="60" y="80" text-anchor="middle"
//            style="pointer-events:none;">Click me.</text>
//
//      <!-- grid -->
//      <text x="10" y="20" text-anchor="middle">0s</text>
//      <line x1="10" y1="25" x2="10" y2="55" stroke="grey" stroke-width=".5" />
//      <text x="35" y="20" text-anchor="middle">2s</text>
//      <line x1="35" y1="25" x2="35" y2="55" stroke="grey" stroke-width=".5" />
//      <text x="60" y="20" text-anchor="middle">4s</text>
//      <line x1="60" y1="25" x2="60" y2="55" stroke="grey" stroke-width=".5" />
//      <text x="85" y="20" text-anchor="middle">6s</text>
//      <line x1="85" y1="25" x2="85" y2="55" stroke="grey" stroke-width=".5" />
//      <text x="110" y="20" text-anchor="middle">8s</text>
//      <line x1="110" y1="25" x2="110" y2="55" stroke="grey" stroke-width=".5" />
//
//      <line x1="10" y1="30" x2="110" y2="30" stroke="grey" stroke-width=".5" />
//      <line x1="10" y1="55" x2="110" y2="55" stroke="grey" stroke-width=".5" />
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
	"time"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().
		Width(120).
		Height(120).
		ViewBox([]float64{0, 0, 120, 120}).
		XmlnsXLink("http://www.w3.org/1999/xlink").
		Append(

			// animated rectangle
			factoryBrowser.NewTagSvgRect().X(10).Y(35).Height(15).Width(0).Append(
				factoryBrowser.NewTagSvgAnimate().AttributeName("width").From(0).To(100).Begin("startButton.click").Dur(8*time.Second).Fill("freeze"),
			),

			// trigger
			factoryBrowser.NewTagSvgRect().Id("startButton").Style("cursor:pointer;").X(19.5).Y(62.5).Rx(5).Height(25).Width(80).Fill("#EFEFEF").Stroke(factoryColor.NewBlack()).StrokeWidth(1),

			factoryBrowser.NewTagSvgText().X(60).Y(80).TextAnchor(html.KSvgTextAnchorMiddle).Style("pointer-events:none;").Text("Click me."),

			// grid
			factoryBrowser.NewTagSvgText().X(10).Y(20).TextAnchor(html.KSvgTextAnchorMiddle).Text("0s"),
			factoryBrowser.NewTagSvgLine().X1(10).Y1(25).X2(10).Y2(55).Stroke(factoryColor.NewGrey()).StrokeWidth(0.5),
			factoryBrowser.NewTagSvgText().X(35).Y(20).TextAnchor(html.KSvgTextAnchorMiddle).Text("2s"),
			factoryBrowser.NewTagSvgLine().X1(35).Y1(25).X2(35).Y2(55).Stroke(factoryColor.NewGrey()).StrokeWidth(0.5),
			factoryBrowser.NewTagSvgText().X(60).Y(20).TextAnchor(html.KSvgTextAnchorMiddle).Text("4s"),
			factoryBrowser.NewTagSvgLine().X1(60).Y1(25).X2(60).Y2(55).Stroke(factoryColor.NewGrey()).StrokeWidth(0.5),
			factoryBrowser.NewTagSvgText().X(85).Y(20).TextAnchor(html.KSvgTextAnchorMiddle).Text("6s"),
			factoryBrowser.NewTagSvgLine().X1(85).Y1(25).X2(85).Y2(55).Stroke(factoryColor.NewGrey()).StrokeWidth(0.5),
			factoryBrowser.NewTagSvgText().X(110).Y(20).TextAnchor(html.KSvgTextAnchorMiddle).Text("8s"),
			factoryBrowser.NewTagSvgLine().X1(110).Y1(25).X2(110).Y2(55).Stroke(factoryColor.NewGrey()).StrokeWidth(0.5),

			factoryBrowser.NewTagSvgLine().X1(10).Y1(30).X2(110).Y2(30).Stroke(factoryColor.NewGrey()).StrokeWidth(0.5),
			factoryBrowser.NewTagSvgLine().X1(10).Y1(55).X2(110).Y2(55).Stroke(factoryColor.NewGrey()).StrokeWidth(0.5),
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
