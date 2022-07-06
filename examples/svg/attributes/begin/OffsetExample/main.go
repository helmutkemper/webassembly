// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/begin
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/begin
//
//  <svg width="120" height="120"  viewBox="0 0 120 120"
//       xmlns="http://www.w3.org/2000/svg" version="1.1">
//
//      <!-- animated rectangles -->
//      <rect x="10" y="35" height="15" width="0">
//          <animate attributeType="XML" attributeName="width" to="100"
//                   begin="0s" dur="8s"
//                   fill="freeze" />
//      </rect>
//
//      <rect x="35" y="60" height="15" width="0">
//          <animate attributeType="XML" attributeName="width" to="75"
//                   begin="2s" dur="6s"
//                   fill="freeze" />
//      </rect>
//
//      <rect x="60" y="85" height="15" width="0">
//          <animate attributeType="XML" attributeName="width" to="50"
//                   begin="4s" dur="4s"
//                   fill="freeze" />
//      </rect>
//
//      <!-- grid -->
//      <text x="10" y="20" text-anchor="middle">0s</text>
//      <line x1="10" y1="25" x2="10" y2="105" stroke="grey" stroke-width=".5" />
//      <text x="35" y="20" text-anchor="middle">2s</text>
//      <line x1="35" y1="25" x2="35" y2="105" stroke="grey" stroke-width=".5" />
//      <text x="60" y="20" text-anchor="middle">4s</text>
//      <line x1="60" y1="25" x2="60" y2="105" stroke="grey" stroke-width=".5" />
//      <text x="85" y="20" text-anchor="middle">6s</text>
//      <line x1="85" y1="25" x2="85" y2="105" stroke="grey" stroke-width=".5" />
//      <text x="110" y="20" text-anchor="middle">8s</text>
//      <line x1="110" y1="25" x2="110" y2="105" stroke="grey" stroke-width=".5" />
//
//      <line x1="10" y1="30" x2="110" y2="30" stroke="grey" stroke-width=".5" />
//      <line x1="10" y1="105" x2="110" y2="105" stroke="grey" stroke-width=".5" />
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

	s1 := factoryBrowser.NewTagSvg().Width(120).Height(120).ViewBox([]float64{0, 0, 120, 120}).Append(

		// animated rectangles
		factoryBrowser.NewTagSvgRect().X(10).Y(35).Height(15).Width(0).Append(
			factoryBrowser.NewTagSvgAnimate().AttributeName("width").To(100).Begin(0*time.Second).Dur(8*time.Second).Fill("freeze"),
		),
		factoryBrowser.NewTagSvgRect().X(35).Y(60).Height(15).Width(0).Append(
			factoryBrowser.NewTagSvgAnimate().AttributeName("width").To(75).Begin(2*time.Second).Dur(6*time.Second).Fill("freeze"),
		),
		factoryBrowser.NewTagSvgRect().X(60).Y(85).Height(15).Width(0).Append(
			factoryBrowser.NewTagSvgAnimate().AttributeName("width").To(50).Begin(4*time.Second).Dur(4*time.Second).Fill("freeze"),
		),

		// Grid
		factoryBrowser.NewTagSvgText().X(10).Y(20).TextAnchor(html.KSvgTextAnchorMiddle).Text("0s"),
		factoryBrowser.NewTagSvgLine().X1(10).Y1(25).X2(10).Y2(105).Stroke(factoryColor.NewGray()).StrokeWidth(0.5),
		factoryBrowser.NewTagSvgText().X(35).Y(20).TextAnchor(html.KSvgTextAnchorMiddle).Text("2s"),
		factoryBrowser.NewTagSvgLine().X1(35).Y1(25).X2(35).Y2(105).Stroke(factoryColor.NewGray()).StrokeWidth(0.5),
		factoryBrowser.NewTagSvgText().X(60).Y(20).TextAnchor(html.KSvgTextAnchorMiddle).Text("4s"),
		factoryBrowser.NewTagSvgLine().X1(60).Y1(25).X2(60).Y2(105).Stroke(factoryColor.NewGray()).StrokeWidth(0.5),
		factoryBrowser.NewTagSvgText().X(85).Y(20).TextAnchor(html.KSvgTextAnchorMiddle).Text("6s"),
		factoryBrowser.NewTagSvgLine().X1(85).Y1(25).X2(85).Y2(105).Stroke(factoryColor.NewGray()).StrokeWidth(0.5),
		factoryBrowser.NewTagSvgText().X(110).Y(20).TextAnchor(html.KSvgTextAnchorMiddle).Text("8s"),
		factoryBrowser.NewTagSvgLine().X1(110).Y1(25).X2(110).Y2(105).Stroke(factoryColor.NewGray()).StrokeWidth(0.5),

		factoryBrowser.NewTagSvgLine().X1(10).Y1(30).X2(110).Y2(30).Stroke(factoryColor.NewGray()).StrokeWidth(0.5),
		factoryBrowser.NewTagSvgLine().X1(10).Y1(105).X2(110).Y2(105).Stroke(factoryColor.NewGray()).StrokeWidth(0.5),
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
