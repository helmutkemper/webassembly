// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/spreadMethod
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/spreadMethod
//
//  <svg width="220" height="150" xmlns="http://www.w3.org/2000/svg">
//    <defs>
//      <linearGradient id="PadGradient"
//                      x1="33%" x2="67%">
//        <stop offset="0%"  stop-color="fuchsia"/>
//        <stop offset="100%" stop-color="orange"/>
//      </linearGradient>
//      <linearGradient id="ReflectGradient" spreadMethod="reflect"
//                      x1="33%" x2="67%">
//        <stop offset="0%"  stop-color="fuchsia"/>
//        <stop offset="100%" stop-color="orange"/>
//      </linearGradient>
//      <linearGradient id="RepeatGradient" spreadMethod="repeat"
//                      x1="33%" x2="67%">
//        <stop offset="0%"  stop-color="fuchsia"/>
//        <stop offset="100%" stop-color="orange"/>
//      </linearGradient>
//    </defs>
//
//    <rect fill="url(#PadGradient)"
//            x="10" y="0" width="200" height="40"/>
//    <rect fill="url(#ReflectGradient)"
//            x="10" y="50" width="200" height="40"/>
//    <rect fill="url(#RepeatGradient)"
//            x="10" y="100" width="200" height="40"/>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().Width(220).Height(150).Append(
		factoryBrowser.NewTagSvgDefs().Append(
			factoryBrowser.NewTagSvgLinearGradient().Id("PadGradient").X1(float32(0.33)).X2(float32(0.67)).Append(
				factoryBrowser.NewTagSvgStop().Offset(float32(0)).StopColor(factoryColor.NewFuchsia()),
				factoryBrowser.NewTagSvgStop().Offset(float32(1)).StopColor(factoryColor.NewOrange()),
			),
			factoryBrowser.NewTagSvgLinearGradient().Id("ReflectGradient").SpreadMethod(html.KSvgSpreadMethodReflect).X1(float32(0.33)).X2(float32(0.67)).Append(
				factoryBrowser.NewTagSvgStop().Offset(float32(0)).StopColor(factoryColor.NewFuchsia()),
				factoryBrowser.NewTagSvgStop().Offset(float32(1)).StopColor(factoryColor.NewOrange()),
			),
			factoryBrowser.NewTagSvgLinearGradient().Id("RepeatGradient").SpreadMethod(html.KSvgSpreadMethodRepeat).X1(float32(0.33)).X2(float32(0.67)).Append(
				factoryBrowser.NewTagSvgStop().Offset(float32(0)).StopColor(factoryColor.NewFuchsia()),
				factoryBrowser.NewTagSvgStop().Offset(float32(1)).StopColor(factoryColor.NewOrange()),
			),
		),

		factoryBrowser.NewTagSvgRect().Fill("url(#PadGradient)").X(10).Y(0).Width(200).Height(40),
		factoryBrowser.NewTagSvgRect().Fill("url(#ReflectGradient)").X(10).Y(50).Width(200).Height(40),
		factoryBrowser.NewTagSvgRect().Fill("url(#RepeatGradient)").X(10).Y(100).Width(200).Height(40),
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
