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
//                   id="myLoop" begin="0s;myLoop.end" dur="4s"
//                   repeatCount="3" />
//
//          <set attributeType="CSS" attributeName="fill" to="green"
//               begin="myLoop.begin" />
//
//          <set attributeType="CSS" attributeName="fill" to="gold"
//               begin="myLoop.repeat(1)" />
//
//          <set attributeType="CSS" attributeName="fill" to="red"
//               begin="myLoop.repeat(2)" />
//      </rect>
//
//      <!-- grid -->
//      <text x="10" y="20" text-anchor="middle">0s</text>
//      <line x1="10" y1="25" x2="10" y2="55" stroke="grey" stroke-width=".5" />
//      <text x="35" y="20" text-anchor="middle">1s</text>
//      <line x1="35" y1="25" x2="35" y2="55" stroke="grey" stroke-width=".5" />
//      <text x="60" y="20" text-anchor="middle">2s</text>
//      <line x1="60" y1="25" x2="60" y2="55" stroke="grey" stroke-width=".5" />
//      <text x="85" y="20" text-anchor="middle">3s</text>
//      <line x1="85" y1="25" x2="85" y2="55" stroke="grey" stroke-width=".5" />
//      <text x="110" y="20" text-anchor="middle">4s</text>
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

	s1 := factoryBrowser.NewTagSvg().Width(120).Height(120).ViewBox([]float64{0, 0, 120, 120}).XmlnsXLink("http://www.w3.org/1999/xlink").Append(

		// animated rectangle
		factoryBrowser.NewTagSvgRect().X(10).Y(35).Height(15).Width(0).Append(
			factoryBrowser.NewTagSvgAnimate().AttributeName("width").From(0).To(100).Id("myLoop").Begin("0s;myLoop.end").Dur(4*time.Second).RepeatCount(3),
			factoryBrowser.NewTagSvgSet().AttributeName("fill").To(factoryColor.NewGreen()).Begin("myLoop.begin"),
			factoryBrowser.NewTagSvgSet().AttributeName("fill").To(factoryColor.NewGold()).Begin("myLoop.repeat(1)"),
			factoryBrowser.NewTagSvgSet().AttributeName("fill").To(factoryColor.NewRed()).Begin("myLoop.repeat(2)"),
		),

		// grid
		factoryBrowser.NewTagSvgText().X(10).Y(20).TextAnchor(html.KSvgTextAnchorMiddle).Text("0s"),
		factoryBrowser.NewTagSvgLine().X1(10).Y1(25).X2(10).Y2(55).Stroke(factoryColor.NewGray()).StrokeWidth(0.5),
		factoryBrowser.NewTagSvgText().X(35).Y(20).TextAnchor(html.KSvgTextAnchorMiddle).Text("1s"),
		factoryBrowser.NewTagSvgLine().X1(35).Y1(25).X2(35).Y2(55).Stroke(factoryColor.NewGray()).StrokeWidth(0.5),
		factoryBrowser.NewTagSvgText().X(60).Y(20).TextAnchor(html.KSvgTextAnchorMiddle).Text("2s"),
		factoryBrowser.NewTagSvgLine().X1(60).Y1(25).X2(60).Y2(55).Stroke(factoryColor.NewGray()).StrokeWidth(0.5),
		factoryBrowser.NewTagSvgText().X(85).Y(20).TextAnchor(html.KSvgTextAnchorMiddle).Text("3s"),
		factoryBrowser.NewTagSvgLine().X1(85).Y1(25).X2(85).Y2(55).Stroke(factoryColor.NewGray()).StrokeWidth(0.5),
		factoryBrowser.NewTagSvgText().X(110).Y(20).TextAnchor(html.KSvgTextAnchorMiddle).Text("4s"),
		factoryBrowser.NewTagSvgLine().X1(110).Y1(25).X2(110).Y2(55).Stroke(factoryColor.NewGray()).StrokeWidth(0.5),

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
