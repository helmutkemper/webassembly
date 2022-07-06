// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/fy
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/fy
//
//  <svg viewBox="0 0 120 120" width="200" height="200" xmlns="http://www.w3.org/2000/svg">
//    <defs>
//      <radialGradient id="Gradient" cx="0.5" cy="0.5" r="0.5"
//          fx="0.35" fy="0.35" fr="5%">
//        <stop offset="0%" stop-color="red"/>
//        <stop offset="100%" stop-color="blue"/>
//      </radialGradient>
//    </defs>
//
//    <rect x="10" y="10" rx="15" ry="15" width="100" height="100"
//        fill="url(#Gradient)" stroke="black" stroke-width="2"/>
//
//    <circle cx="60" cy="60" r="50" fill="transparent" stroke="white" stroke-width="2"/>
//    <circle cx="35" cy="35" r="2" fill="white" stroke="white"/>
//    <circle cx="60" cy="60" r="2" fill="white" stroke="white"/>
//    <text x="38" y="40" fill="white" font-family="sans-serif" font-size="10pt">(fx,fy)</text>
//    <text x="63" y="63" fill="white" font-family="sans-serif" font-size="10pt">(cx,cy)</text>
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 120, 120}).Width(200).Height(200).Append(
		factoryBrowser.NewTagSvgDefs().Append(
			factoryBrowser.NewTagSvgRadialGradient().Id("Gradient").Cx(0.5).Cy(0.5).R(0.5).Fx(0.35).Fy(0.35).Fr(float32(0.05)).Append(
				factoryBrowser.NewTagSvgStop().Offset(float32(0.0)).StopColor(factoryColor.NewRed()),
				factoryBrowser.NewTagSvgStop().Offset(float32(1.0)).StopColor(factoryColor.NewBlue()),
			),
		),

		factoryBrowser.NewTagSvgRect().X(10).Y(10).Rx(15).Ry(15).Width(100).Height(100).Fill("url(#Gradient)").Stroke(factoryColor.NewBlack()).StrokeWidth(2),

		factoryBrowser.NewTagSvgCircle().Cx(60).Cy(60).R(50).Fill("transparent").Stroke(factoryColor.NewWhite()).StrokeWidth(2),
		factoryBrowser.NewTagSvgCircle().Cx(35).Cy(35).R(2).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewWhite()),
		factoryBrowser.NewTagSvgCircle().Cx(60).Cy(60).R(2).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewWhite()),
		factoryBrowser.NewTagSvgText().X(38).Y(40).Fill(factoryColor.NewWhite()).FontFamily("sans-serif").FontSize("10pt").Text("(fx,fy)"),
		factoryBrowser.NewTagSvgText().X(63).Y(63).Fill(factoryColor.NewWhite()).FontFamily("sans-serif").FontSize("10pt").Text("(cx,cy)"),
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
