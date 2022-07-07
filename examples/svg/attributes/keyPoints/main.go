// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/keyPoints
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/keyPoints
//
//  <svg viewBox="0 0 120 120" xmlns="http://www.w3.org/2000/svg"
//      xmlns:xlink="http://www.w3.org/1999/xlink">
//    <path d="M10,110 A120,120 -45 0,1 110 10 A120,120 -45 0,1 10,110"
//        stroke="lightgrey" stroke-width="2" fill="none" id="motionPath"/>
//    <circle cx="10" cy="110" r="3" fill="lightgrey"/>
//    <circle cx="110" cy="10" r="3" fill="lightgrey"/>
//
//    <circle r="5" fill="red">
//      <animateMotion dur="3s" repeatCount="indefinite"
//          keyPoints="0;0.5;1" keyTimes="0;0.15;1" calcMode="linear">
//        <mpath xlink:href="#motionPath"/>
//      </animateMotion>
//    </circle>
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 120, 120}).XmlnsXLink("http://www.w3.org/1999/xlink").Append(
		factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(10, 110).A(120, 120, -45, 0, 1, 110, 10).A(120, 120, -45, 0, 1, 10, 110)).Stroke(factoryColor.NewLightgray()).StrokeWidth(2).Fill(nil).Id("motionPath"),
		factoryBrowser.NewTagSvgCircle().Cx(10).Cy(110).R(3).Fill(factoryColor.NewLightgray()),
		factoryBrowser.NewTagSvgCircle().Cx(110).Cy(10).R(3).Fill(factoryColor.NewLightgray()),

		factoryBrowser.NewTagSvgCircle().R(5).Fill(factoryColor.NewRed()).Append(
			factoryBrowser.NewTagSvgAnimateMotion().Dur(3*time.Second).RepeatCount(html.KSvgDurIndefinite).KeyPoints([]float64{0.0, 0.5, 1.0}).KeyTimes([]float64{0.0, 0.15, 1.0}).CalcMode(html.KSvgCalcModeLinear).Append(
				factoryBrowser.NewTagSvgMPath().HRef("#motionPath"),
			),
		),
	)

	stage.Append(s1)

	<-done
}
