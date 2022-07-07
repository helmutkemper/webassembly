// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/keySplines
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/keySplines
//
//  <svg viewBox="0 0 120 120" xmlns="http://www.w3.org/2000/svg">
//    <circle cx="60" cy="10" r="10">
//      <animate attributeName="cx" dur="4s" calcMode="spline" repeatCount="indefinite"
//          values="60; 110; 60; 10; 60" keyTimes="0; 0.25; 0.5; 0.75; 1"
//          keySplines="0.5 0 0.5 1; 0.5 0 0.5 1; 0.5 0 0.5 1; 0.5 0 0.5 1"/>
//      <animate attributeName="cy" dur="4s" calcMode="spline" repeatCount="indefinite"
//          values="10; 60; 110; 60; 10" keyTimes="0; 0.25; 0.5; 0.75; 1"
//          keySplines="0.5 0 0.5 1; 0.5 0 0.5 1; 0.5 0 0.5 1; 0.5 0 0.5 1"/>
//    </circle>
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"time"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 120, 120}).XmlnsXLink("http://www.w3.org/1999/xlink").Append(
		factoryBrowser.NewTagSvgCircle().Cx(60).Cy(10).R(10).Append(
			factoryBrowser.NewTagSvgAnimate().AttributeName("cx").Dur(4*time.Second).CalcMode(html.KSvgCalcModeSpline).RepeatCount(html.KSvgDurIndefinite).
				Values([]float64{60, 110, 60, 10, 60}).KeyTimes([]float64{0, 0.25, 0.5, 0.75, 1}).KeySplines([][]float64{{0.5, 0, 0.5, 1}, {0.5, 0, 0.5, 1}, {0.5, 0, 0.5, 1}, {0.5, 0, 0.5, 1}}),
			factoryBrowser.NewTagSvgAnimate().AttributeName("cy").Dur(4*time.Second).CalcMode(html.KSvgCalcModeSpline).RepeatCount(html.KSvgDurIndefinite).
				Values([]float64{10, 60, 110, 60, 10}).KeyTimes([]float64{0, 0.25, 0.5, 0.75, 1}).KeySplines([][]float64{{0.5, 0, 0.5, 1}, {0.5, 0, 0.5, 1}, {0.5, 0, 0.5, 1}, {0.5, 0, 0.5, 1}}),
		),
	)

	stage.Append(s1)

	<-done
}
